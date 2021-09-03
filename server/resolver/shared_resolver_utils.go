package resolver

import (
	"errors"
	"fmt"

	"github.com/Luismorlan/newsmux/model"
	"github.com/Luismorlan/newsmux/utils"
	. "github.com/Luismorlan/newsmux/utils/log"
	"gorm.io/gorm"
)

const (
	feedRefreshLimit      = 30
	defaultCursor         = -1
	maxRepublishDBBatches = 5
)

// Given a list of FeedRefreshInput, get posts for the requested feeds
// Do it by iterating through feeds
func getRefreshPosts(r *queryResolver, queries []*model.FeedRefreshInput) ([]*model.Feed, error) {
	results := []*model.Feed{}

	//TODO: can be run in parallel
	for ind, _ := range queries {
		query := queries[ind]
		if query == nil {
			// This is not expected since gqlgen guarantees it is not nil
			continue
		}
		// Prepare feed basic info
		var feed model.Feed
		queryResult := r.DB.Preload("SubSources").Where("id = ?", query.FeedID).First(&feed)
		if queryResult.RowsAffected != 1 {
			return []*model.Feed{}, fmt.Errorf("invalid feed id %s", query.FeedID)
		}
		sanitizeFeedsQueryInput(query, &feed)
		if err := getFeedPostsOrRePublish(r.DB, &feed, query); err != nil {
			return []*model.Feed{}, fmt.Errorf("failure when get posts for feed id %s", feed.Id)
		}
		results = append(results, &feed)
	}

	return results, nil
}

func getFeedPostsOrRePublish(db *gorm.DB, feed *model.Feed, query *model.FeedRefreshInput) error {
	var posts []*model.Post
	// try to read published posts
	Log.Info("read published post for feed: ", feed.Id, " query: ", query)
	if query.Direction == model.FeedRefreshDirectionNew {
		db.Model(&model.Post{}).
			Joins("LEFT JOIN post_feed_publishes ON post_feed_publishes.post_id = posts.id").
			Joins("LEFT JOIN feeds ON post_feed_publishes.feed_id = feeds.id").
			Where("feed_id = ? AND posts.cursor > ?", feed.Id, query.Cursor).
			Order("cursor desc").
			Limit(query.Limit).
			Find(&posts)
	} else {
		db.Model(&model.Post{}).
			Joins("LEFT JOIN post_feed_publishes ON post_feed_publishes.post_id = posts.id").
			Joins("LEFT JOIN feeds ON post_feed_publishes.feed_id = feeds.id").
			Where("feed_id = ? AND posts.cursor < ?", feed.Id, query.Cursor).
			Order("cursor desc").
			Limit(query.Limit).
			Find(&posts)
	}

	feed.Posts = posts

	// cases where we need republish first
	if query.Direction == model.FeedRefreshDirectionNew {
		// query NEW but publish table is empty
		var count int64
		db.Model(&model.PostFeedPublish{}).
			Joins("LEFT JOIN feeds ON post_feed_publishes.feed_id = feeds.id").
			Where("feed_id = ?", feed.Id).
			Count(&count)
		if count == 0 {
			Log.Info("run ondemand publish posts to feed: ", feed.Id, " triggered by NEW in {feeds} API")
			rePublishPostsFromCursor(db, feed, query.Limit, -1)
		} else {
			fmt.Println("NOT REPUBLISHING", count)
		}
	} else {
		// query OLD but can't satisfy the limit
		lastCursor := query.Cursor
		if len(posts) < query.Limit {
			if len(posts) > 0 {
				lastCursor = int(posts[len(posts)-1].Cursor)
			}
			Log.Info("run ondemand publish posts to feed: ", feed.Id, " triggered by NEW in {feeds} API from curosr ", lastCursor)
			// republish to fulfill the query limit
			rePublishPostsFromCursor(db, feed, query.Limit-len(posts), lastCursor)
		}
	}
	return nil
}

// Redo posts publish to feeds
// From a particular cursor down
// If cursor is -1, republish from NEWest
func rePublishPostsFromCursor(db *gorm.DB, feed *model.Feed, limit int, fromCursor int) {
	var (
		postsToPublish []*model.Post
		batches        = 0
	)

	limit = utils.Min(feedRefreshLimit, limit)

	if fromCursor == -1 {
		fromCursor = 2147483647
	}

	var subsourceIds []string
	for _, subsource := range feed.SubSources {
		subsourceIds = append(subsourceIds, subsource.Id)
	}

	for {
		if len(postsToPublish) >= limit || batches > maxRepublishDBBatches {
			break
		}
		var postsCandidates []model.Post
		// 1. Read subsources' most recent posts
		db.Model(&model.Post{}).
			Joins("LEFT JOIN sub_sources ON posts.sub_source_id = sub_sources.id").
			Where("sub_sources.id IN ? AND posts.cursor < ?", subsourceIds, fromCursor).
			Order("cursor desc").
			Limit(limit).
			Find(&postsCandidates)

		// 2. Try match postsCandidate with Feed
		for ind := range postsCandidates {
			post := postsCandidates[ind]
			fromCursor = int(post.Cursor)
			matched, error := utils.DataExpressionMatchPost(string(feed.FilterDataExpression), post)
			if error != nil {
				continue
			}
			if matched {
				postsToPublish = append(postsToPublish, &post)
				// to publish exact same number of posts queried
				if len(postsToPublish) >= limit {
					break
				}
			}
		}
		batches = batches + 1
	}

	// This call will also update feed object with posts, no need to append
	db.Model(feed).UpdateColumns(model.Feed{UpdatedAt: feed.UpdatedAt}).Association("Posts").Append(postsToPublish)
}

// get all feeds a user subscribed
func getUserSubscriptions(r *queryResolver, userID string) ([]*model.Feed, error) {
	var user model.User
	queryResult := r.DB.Where("id = ?", userID).Preload("SubscribedFeeds").First(&user)
	if queryResult.RowsAffected != 1 {
		return nil, errors.New("User not found")
	}
	return user.SubscribedFeeds, nil
}

func sanitizeFeedsQueryInput(query *model.FeedRefreshInput, feed *model.Feed) {
	// Check if requested cursors are out of sync from last feed update
	// If out of sync, default to query latest posts
	// Use unix() to avoid accuracy loss due to gqlgen serialization impacting matching
	if query.FeedUpdatedTime == nil || query.FeedUpdatedTime.Unix() != feed.UpdatedAt.Unix() {
		Log.Info(
			"requested with outdated feed updated time, feed_id=", feed.Id,
			" query updated time=", query.FeedUpdatedTime,
			" feed updated at=", feed.UpdatedAt)
		query.Cursor = -1
		query.Direction = model.FeedRefreshDirectionNew
	}

	// Cap query limit
	if query.Limit < 0 || query.Limit > feedRefreshLimit {
		query.Limit = feedRefreshLimit
	}
}
