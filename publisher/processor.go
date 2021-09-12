package publisher

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/Luismorlan/newsmux/model"
	. "github.com/Luismorlan/newsmux/protocol"
	"github.com/Luismorlan/newsmux/server/resolver"
	. "github.com/Luismorlan/newsmux/utils"
	. "github.com/Luismorlan/newsmux/utils/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

type CrawlerpublisherMessageProcessor struct {
	Reader MessageQueueReader
	DB     *gorm.DB
}

// Create new processor with reader dependency injection
func NewPublisherMessageProcessor(reader MessageQueueReader, db *gorm.DB) *CrawlerpublisherMessageProcessor {
	return &CrawlerpublisherMessageProcessor{
		Reader: reader,
		DB:     db,
	}
}

// Use Reader to read N messages and process them in parallel
// Time out or queue name etc are defined in reader
// Reader focus on how to get message from queue
// Processor focus on how to process the message
// This function doesn't return anything, only log errors
func (processor *CrawlerpublisherMessageProcessor) ReadAndProcessMessages(maxNumberOfMessages int64) {
	// Pull queued messages from queue
	msgs, err := processor.Reader.ReceiveMessages(maxNumberOfMessages)

	if err != nil {
		Log.Error("fail read crawler messages from queue : ", err)
		return
	}

	// Process
	// TODO: process in parallel
	for _, msg := range msgs {
		if err := processor.ProcessOneCralwerMessage(msg); err != nil {
			Log.Error("fail process one crawler message : ", err)
			continue
		}
	}
}

// Dedup posts using only content, switch to use deduplication_id generated by crawler
func (processor *CrawlerpublisherMessageProcessor) findDuplicatedPost(decodedMsg *CrawlerMessage) (bool, *model.Post) {
	var post model.Post
	queryResult := processor.DB.Where(
		"deduplicate_id = ? ",
		decodedMsg.Post.DeduplicateId,
	).First(&post)

	return queryResult.RowsAffected != 0, &post
}

func (processor *CrawlerpublisherMessageProcessor) prepareFeedCandidates(
	subSource *model.SubSource,
) map[string]*model.Feed {
	feedCandidates := make(map[string]*model.Feed)

	if subSource != nil {
		for _, feed := range subSource.Feeds {
			feedCandidates[feed.Id] = feed
		}
	}
	return feedCandidates
}

func (processor *CrawlerpublisherMessageProcessor) prepareSource(id string) (*model.Source, error) {
	var res model.Source
	if len(id) > 0 {
		result := processor.DB.Preload("Feeds").Where("id = ?", id).First(&res)
		if result.RowsAffected != 1 {
			return nil, errors.New(fmt.Sprintf("source not found: %s", id))
		}
		return &res, nil
	} else {
		return nil, errors.New("source id can not be empty")
	}
}

func (processor *CrawlerpublisherMessageProcessor) prepareSubSourceRecursive(post *CrawlerMessage_CrawledPost, isRoot bool) (*model.SubSource, error) {
	subSource, err := resolver.UpsertSubsourceImpl(processor.DB, model.UpsertSubSourceInput{
		Name:               post.SubSource.SubSourceName,
		ExternalIdentifier: post.SubSource.SubSourceExternalId,
		SourceID:           post.SubSource.SubSourceSourceId,
		AvatarURL:          post.SubSource.SubSourceProfileUrl,
		OriginURL:          post.SubSource.SubSourceOriginUrl,
		IsFromSharedPost:   !isRoot,
	})
	if err != nil {
		return nil, err
	}
	post.SubSource.SubSourceId = subSource.Id

	if post.SharedFromCrawledPost != nil {
		if _, err = processor.prepareSubSourceRecursive(post.SharedFromCrawledPost, false); err != nil {
			return nil, err
		}
	}
	return subSource, nil
}

func (processor *CrawlerpublisherMessageProcessor) preparePostChainFromMessage(crawledPost *CrawlerMessage_CrawledPost, isRoot bool) (post *model.Post, e error) {
	var subSource model.SubSource
	res := processor.DB.Where("id = ?", crawledPost.SubSource.SubSourceId).First(&subSource)
	if res.RowsAffected == 0 {
		return nil, errors.New("invalid subsource id " + crawledPost.SubSource.SubSourceId)
	}

	post = &model.Post{
		Id:             uuid.New().String(),
		Title:          crawledPost.Title,
		Content:        crawledPost.Content,
		CreatedAt:      time.Now(),
		SubSource:      subSource,
		SubSourceID:    crawledPost.SubSource.SubSourceId,
		SavedByUser:    []*model.User{},
		PublishedFeeds: []*model.Feed{},
		InSharingChain: !isRoot,
		DeduplicateId:  crawledPost.DeduplicateId,
	}
	if crawledPost.SharedFromCrawledPost != nil {
		sharedFromPost, e := processor.preparePostChainFromMessage(crawledPost.SharedFromCrawledPost, false)
		if e != nil {
			return nil, e
		}
		post.SharedFromPost = sharedFromPost
		post.SharedFromPostID = &sharedFromPost.Id
	}
	return post, nil
}

// Process one cralwer-publisher message in following major steps:
// Step1. decode into protobuf generated struct
// Step2. update subsource
// Step2. deduplication
// Step3. do publishing with new post, also handle recursive shared_from posts
// Step4. if publishing succeeds, delete message in queue
func (processor *CrawlerpublisherMessageProcessor) ProcessOneCralwerMessage(msg *MessageQueueMessage) error {
	Log.Info("process queued message")

	decodedMsg, err := processor.decodeCrawlerMessage(msg)
	if err != nil {
		processor.Reader.DeleteMessage(msg)
		return err
	}

	// Once get a message, check if there is exact same Post (same sources, same content), if not store into DB as Post
	// TODO: use cralwer generated dedup_id for dedup (dedup_id is something like external identifier)
	if duplicated, existingPost := processor.findDuplicatedPost(decodedMsg); duplicated == true {
		return errors.New(fmt.Sprintf("message has already been processed, existing post_id: %s", existingPost.Id))
	}

	// Prepare Post relations to Subsources (Sources can be inferred)
	subSource, err := processor.prepareSubSourceRecursive(decodedMsg.Post /*isRoot*/, true)
	if err != nil {
		return err
	}

	// Load feeds into memory based on source and subsource of the post
	feedCandidates := processor.prepareFeedCandidates(subSource)

	// Create new post based on message
	post, err := processor.preparePostChainFromMessage(decodedMsg.Post /*isRoot*/, true)
	if err != nil {
		return err
	}

	// Check each feed's source/subsource and data expression
	feedsToPublish := []*model.Feed{}
	for _, feed := range feedCandidates {
		// since we will append pointer, we need to have a var each iteration
		// otherwise feeds appended will be reused and all feeds in the slice are same
		// feed := feedCandidates[ind]
		// Once a message is matched to a feed, write the PostFeedPublish relation to DB
		matched, err := DataExpressionMatchPostChain(feed.FilterDataExpression.String(), post)
		if err != nil {
			return err
		}
		if matched {
			feedsToPublish = append(feedsToPublish, feed)
		}
	}

	// Write to DB, post creation and publish is in a transaction
	err = processor.DB.Transaction(func(tx *gorm.DB) error {
		processor.DB.Create(&post)
		processor.DB.Model(&post).Association("PublishedFeeds").Append(feedsToPublish)
		return nil
	})
	if err != nil {
		return err
	}

	// Delete message from queue
	return processor.Reader.DeleteMessage(msg)
}

// Parse message into meaningful structure CrawlerMessage
// This function assumes message passed in can be parsed, otherwise it will throw error
func (processor *CrawlerpublisherMessageProcessor) decodeCrawlerMessage(msg *MessageQueueMessage) (*CrawlerMessage, error) {
	str, err := msg.Read()
	if err != nil {
		return nil, err
	}

	sDec, err := b64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	decodedMsg := &CrawlerMessage{}
	if err := proto.Unmarshal(sDec, decodedMsg); err != nil {
		return nil, err
	}

	return decodedMsg, nil
}
