package publisher

import (
	"context"
	"encoding/base64"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/Luismorlan/newsmux/bot"
	"github.com/Luismorlan/newsmux/collector"
	"github.com/Luismorlan/newsmux/model"
	"github.com/Luismorlan/newsmux/protocol"
	. "github.com/Luismorlan/newsmux/protocol"
	"github.com/Luismorlan/newsmux/server/resolver"
	. "github.com/Luismorlan/newsmux/utils"
	. "github.com/Luismorlan/newsmux/utils/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

const SemanticHashingLength = 128

type CrawlerpublisherMessageProcessor struct {
	Reader MessageQueueReader
	DB     *gorm.DB

	// gRPC Client and connection
	Client protocol.DeduplicatorClient
}

// Create new processor with reader dependency injection
func NewPublisherMessageProcessor(
	reader MessageQueueReader,
	db *gorm.DB,
	client protocol.DeduplicatorClient,
) *CrawlerpublisherMessageProcessor {
	return &CrawlerpublisherMessageProcessor{
		Reader: reader,
		DB:     db,
		Client: client,
	}
}

// Use Reader to read N messages and process them in parallel
// Time out or queue name etc are defined in reader
// Reader focus on how to get message from queue
// Processor focus on how to process the message
// This function doesn't return anything, only log errors
func (processor *CrawlerpublisherMessageProcessor) ReadAndProcessMessages(sqsReadBatchSize int64) int {
	// Pull queued messages from queue
	msgs, err := processor.Reader.ReceiveMessages(sqsReadBatchSize)

	successCount := 0
	if err != nil {
		Log.Error("fail read crawler messages from queue : ", err)
		return successCount
	}

	// TODO: process in parallel, but can involve time ordering issue
	// Process all messages
	for _, msg := range msgs {
		if _, err := processor.ProcessOneCralwerMessage(msg); err != nil {
			Log.Errorf("fail process one crawler message. err: %s , message: %s", err, *msg.Message)
			continue
		}
		successCount++
	}
	return successCount
}

func (processor *CrawlerpublisherMessageProcessor) calculateSemanticHashing(decodedMsg *CrawlerMessage) (string, error) {
	// We don't calculate semantic hashing for Wechat message because their
	// contents are all same.
	if decodedMsg.Post.SubSource.SourceId == collector.WeixinSourceId {
		return "", nil
	}

	// Calculate semanticHashing by calling Deduplicator.
	ctx := context.Background()
	res, err := processor.Client.GetSimHash(ctx, &protocol.GetSimHashRequest{
		Text:   decodedMsg.Post.Content,
		Length: SemanticHashingLength,
	})
	if err != nil || len(res.Binary) != SemanticHashingLength {
		Log.Errorln("fail to calculate the semantic hashing for post: ", decodedMsg.String(), "error: ", err, "hashing: ", res.Binary)
		return "", err
	}
	return res.Binary, nil
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

func (processor *CrawlerpublisherMessageProcessor) prepareSubSourceRecursive(post *CrawlerMessage_CrawledPost, isRoot bool) (*model.SubSource, error) {
	subSource, err := resolver.UpsertSubsourceImpl(processor.DB, model.UpsertSubSourceInput{
		Name:               post.SubSource.Name,
		ExternalIdentifier: post.SubSource.ExternalId,
		SourceID:           post.SubSource.SourceId,
		AvatarURL:          post.SubSource.AvatarUrl,
		OriginURL:          post.SubSource.OriginUrl,
		IsFromSharedPost:   !isRoot,
	})
	if err != nil {
		return nil, err
	}
	post.SubSource.Id = subSource.Id

	if post.SharedFromCrawledPost != nil {
		if _, err = processor.prepareSubSourceRecursive(post.SharedFromCrawledPost, false); err != nil {
			return nil, err
		}
	}
	return subSource, nil
}

func (processor *CrawlerpublisherMessageProcessor) preparePostChainFromMessage(msg *CrawlerMessage, currentPost *CrawlerMessage_CrawledPost, isRoot bool) (post *model.Post, e error) {
	var subSource model.SubSource
	res := processor.DB.Where("id = ?", currentPost.SubSource.Id).First(&subSource)
	if res.RowsAffected == 0 {
		return nil, errors.New("invalid subsource id " + currentPost.SubSource.Id)
	}

	post = &model.Post{
		Id:                 uuid.New().String(),
		Title:              currentPost.Title,
		Content:            currentPost.Content,
		CreatedAt:          time.Now(),
		SubSource:          subSource,
		SubSourceID:        currentPost.SubSource.Id,
		SavedByUser:        []*model.User{},
		PublishedFeeds:     []*model.Feed{},
		InSharingChain:     !isRoot,
		DeduplicateId:      currentPost.DeduplicateId,
		CrawledAt:          msg.CrawledAt.AsTime(),
		ContentGeneratedAt: currentPost.ContentGeneratedAt.AsTime(),
		ImageUrls:          currentPost.ImageUrls,
		FileUrls:           currentPost.FilesUrls,
		OriginUrl:          currentPost.OriginUrl,
		// transform tags into serialized string separated by ","
		Tag: strings.Join(currentPost.Tags, ","),
	}
	if currentPost.SharedFromCrawledPost != nil {
		sharedFromPost, e := processor.preparePostChainFromMessage(msg, currentPost.SharedFromCrawledPost, false)
		if e != nil {
			return nil, e
		}
		post.SharedFromPost = sharedFromPost
		post.SharedFromPostID = &sharedFromPost.Id
	}
	return post, nil
}

func (processor *CrawlerpublisherMessageProcessor) MatchMessageWithFeeds(feedCandidates map[string]*model.Feed, post *model.Post) ([]*model.Feed, error) {
	var wg sync.WaitGroup
	ch := make(chan *model.Feed, len(feedCandidates))
	errCh := make(chan error, len(feedCandidates))
	for _, feed := range feedCandidates {
		// since we will append pointer, we need to have a var each iteration
		// otherwise feeds appended will be reused and all feeds in the slice are same
		// feed := feedCandidates[ind]
		// Once a message is matched to a feed, write the PostFeedPublish relation to DB
		wg.Add(1)
		go func(feed *model.Feed) {
			defer wg.Done()
			matched, err := DataExpressionMatchPostChain(feed.FilterDataExpression.String(), post)
			if err != nil {
				errCh <- err
			} else if matched {
				ch <- feed
			}
		}(feed)
	}

	// wait for all goroutines to finish
	wg.Wait()
	close(ch)
	close(errCh)

	feedsToPublish := []*model.Feed{}
	for feed := range ch {
		feedsToPublish = append(feedsToPublish, feed)
	}
	if err, ok := <-errCh; ok {
		return nil, err
	}
	return feedsToPublish, nil
}

// Process one cralwer-publisher message in following major steps:
// Step1. decode into protobuf generated struct
// Step2. update subsource
// Step2. deduplication
// Step3. do publishing with new post, also handle recursive shared_from posts
// Step4. if publishing succeeds, delete message in queue
func (processor *CrawlerpublisherMessageProcessor) ProcessOneCralwerMessage(msg *MessageQueueMessage) (*CrawlerMessage, error) {
	// TODO: bump counter in ddog for number of message processed
	decodedMsg, err := processor.decodeCrawlerMessage(msg)
	if err != nil {
		processor.Reader.DeleteMessage(msg)
		return nil, err
	}

	// Once get a message, check if there is exact same Post (same sources, same content), if not store into DB as Post
	if duplicated, _ := processor.findDuplicatedPost(decodedMsg); duplicated {
		// Log.Infof("[duplicated message] message has already been processed, existing deduplicate_id: %s, existing post_id: %s ", decodedMsg.Post.DeduplicateId, existingPost.Id)
		// TODO: bump counter for deduplicated messages
		err := processor.Reader.DeleteMessage(msg)
		return decodedMsg, err
	}

	// Prepare Post relations to Subsources (Sources can be inferred)
	subSource, err := processor.prepareSubSourceRecursive(decodedMsg.Post /*isRoot*/, true)
	if err != nil {
		return decodedMsg, err
	}

	// Load feeds into memory based on source and subsource of the post
	feedCandidates := processor.prepareFeedCandidates(subSource)

	// Create new post based on message
	post, err := processor.preparePostChainFromMessage(
		decodedMsg,
		decodedMsg.Post,
		/*isRoot*/ true,
	)
	if err != nil {
		return decodedMsg, err
	}

	h, err := processor.calculateSemanticHashing(decodedMsg)
	// Only do soft failure for semantic hashing uncalculated. This is because
	// semantic hashing is a "Good to have" feature, App can still work properly
	// without it.
	if err == nil && len(h) == SemanticHashingLength {
		post.SemanticHashing = h
	} else {
		Log.Logger.Errorln("fail to calculate semantic hashing for message:", decodedMsg.String(), "err:", err, "hashing:", h)
	}

	// Match post with candidate feeds in parallel
	feedsToPublish, err := processor.MatchMessageWithFeeds(feedCandidates, post)
	if err != nil {
		return decodedMsg, err
	}

	for _, f := range feedsToPublish {
		for _, c := range f.SubscribedChannels {
			go bot.PushPostViaWebhook(*post, c.WebhookUrl)
		}
	}

	// Write to DB, post creation and publish is in a transaction
	err = processor.DB.Transaction(func(tx *gorm.DB) error {
		processor.DB.Create(&post)
		err := processor.DB.Model(&post).Association("PublishedFeeds").Append(feedsToPublish)
		return err
	})
	if err != nil {
		return decodedMsg, err
	}

	// Delete message from queue
	return decodedMsg, processor.Reader.DeleteMessage(msg)
}

// Parse message into meaningful structure CrawlerMessage
// This function assumes message passed in can be parsed, otherwise it will throw error
func (processor *CrawlerpublisherMessageProcessor) decodeCrawlerMessage(msg *MessageQueueMessage) (*CrawlerMessage, error) {
	str, err := msg.Read()
	if err != nil {
		return nil, err
	}

	sDec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	decodedMsg := &CrawlerMessage{}
	if err := proto.Unmarshal(sDec, decodedMsg); err != nil {
		return nil, err
	}

	return decodedMsg, nil
}
