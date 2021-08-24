package resolver

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Luismorlan/newsmux/model"
	"github.com/Luismorlan/newsmux/server/graph/generated"
	"github.com/Luismorlan/newsmux/utils"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func PrepareTestForGraphQLAPIs(db *gorm.DB) *client.Client {
	client := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
		DB:             db,
		SeedStateChans: nil,
	}})))
	return client
}

func TestCreateUser(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test User Creation", func(t *testing.T) {
		utils.TestCreateUserAndValidate(t, "test_user_name", db, client)
	})
}

func TestCreateFeed(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test Feed Creation", func(t *testing.T) {
		uid := utils.TestCreateUserAndValidate(t, "test_user_name", db, client)
		feedId := utils.TestCreateFeedAndValidate(t, uid, "test_feed_for_feeds_api", `{\"a\":1}`, []string{}, []string{}, db, client)
		require.NotEmpty(t, feedId)
	})
}

func TestCreateSource(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test Source Creation", func(t *testing.T) {
		uid := utils.TestCreateUserAndValidate(t, "test_user_name", db, client)
		sourceId := utils.TestCreateSourceAndValidate(t, uid, "test_source_for_feeds_api", "test_domain", db, client)
		require.NotEmpty(t, sourceId)
	})
}

func TestCreateSubSource(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test Source Creation", func(t *testing.T) {
		uid := utils.TestCreateUserAndValidate(t, "test_user_name", db, client)
		sourceId := utils.TestCreateSourceAndValidate(t, uid, "test_source_for_feeds_api", "test_domain", db, client)
		subSourceId := utils.TestCreateSubSourceAndValidate(t, uid, "test_subsource_for_feeds_api", "test_externalid", sourceId, db, client)
		require.NotEmpty(t, subSourceId)
	})
}

func TestUserSubscribeFeed(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test User subscribe Feed", func(t *testing.T) {
		uid := utils.TestCreateUserAndValidate(t, "test_user_name", db, client)
		feedId := utils.TestCreateFeedAndValidate(t, uid, "test_feed_for_feeds_api", `{\"a\":1}`, []string{}, []string{}, db, client)
		utils.TestUserSubscribeFeedAndValidate(t, uid, feedId, db, client)
	})
}

func TestQueryFeeds(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	userId := utils.TestCreateUserAndValidate(t, "test_user_for_feeds_api", db, client)
	feedIdOne := utils.TestCreateFeedAndValidate(t, userId, "test_feed_for_feeds_api", `{\"a\":1}`, []string{}, []string{}, db, client)
	feedIdTwo := utils.TestCreateFeedAndValidate(t, userId, "test_feed_for_feeds_api", `{\"a\":1}`, []string{}, []string{}, db, client)
	sourceId := utils.TestCreateSourceAndValidate(t, userId, "test_source_for_feeds_api", "test_domain", db, client)
	utils.TestCreateSubSourceAndValidate(t, userId, "test_subsource_for_feeds_api", "test_externalid", sourceId, db, client)
	utils.TestUserSubscribeFeedAndValidate(t, userId, feedIdOne, db, client)
	utils.TestUserSubscribeFeedAndValidate(t, userId, feedIdTwo, db, client)

	// 0 is oldest post, 6 is newest post
	utils.TestCreatePostAndValidate(t, "test_title_0", "test_content_0", sourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_1", "test_content_1", sourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_2", "test_content_2", sourceId, feedIdOne, db, client)
	_, midCursorFirst := utils.TestCreatePostAndValidate(t, "test_title_3", "test_content_3", sourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_4", "test_content_4", sourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_5", "test_content_5", sourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_6", "test_content_6", sourceId, feedIdOne, db, client)

	// 0 is oldest post, 6 is newest post
	utils.TestCreatePostAndValidate(t, "test_title_0", "test_content_0", sourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_1", "test_content_1", sourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_2", "test_content_2", sourceId, feedIdTwo, db, client)
	_, midCursorSecond := utils.TestCreatePostAndValidate(t, "test_title_3", "test_content_3", sourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_4", "test_content_4", sourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_5", "test_content_5", sourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_6", "test_content_6", sourceId, feedIdTwo, db, client)

	checkFeedTopPosts(t, userId, feedIdOne, midCursorFirst, db, client)
	checkFeedBottomPosts(t, userId, feedIdOne, midCursorFirst, db, client)

	checkFeedTopPostsMultipleFeeds(t, userId, feedIdOne, feedIdTwo, midCursorFirst, midCursorSecond, db, client)
	checkFeedBottomPostsMultipleFeeds(t, userId, feedIdOne, feedIdTwo, midCursorFirst, midCursorSecond, db, client)

	checkFeedTopPostsWithoutSpecifyFeed(t, userId, feedIdOne, feedIdTwo, db, client)
}

func checkFeedTopPosts(t *testing.T, userId string, feedId string, cursor int, db *gorm.DB, client *client.Client) {
	var resp struct {
		Feeds []struct {
			Id        string `json:"id"`
			UpdatedAt string `json:"updatedAt"`
			Posts     []struct {
				Id      string `json:"id"`
				Title   string `json:"title"`
				Content string `json:"content"`
				Cursor  int    `json:"cursor"`
			} `json:"posts"`
		} `json:"feeds"`
	}

	client.MustPost(fmt.Sprintf(`
	query{
		feeds (input : {
		  userId : "%s"
		  feedRefreshInputs : [
			{feedId: "%s", limit: %d, cursor: %d, direction: %s}
		  ]
		}) {
		  id
		  updatedAt
		  posts {
			id
			title 
			content
			cursor
		  }
		}
	  }
	`, userId, feedId, 2, cursor, model.FeedRefreshDirectionNew), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	require.Equal(t, 1, len(resp.Feeds))
	require.Equal(t, feedId, resp.Feeds[0].Id)
	require.Equal(t, 2, len(resp.Feeds[0].Posts))
	require.Equal(t, "test_title_6", resp.Feeds[0].Posts[0].Title)
	require.Equal(t, "test_title_5", resp.Feeds[0].Posts[1].Title)
}

func checkFeedBottomPosts(t *testing.T, userId string, feedId string, cursor int, db *gorm.DB, client *client.Client) {
	var resp struct {
		Feeds []struct {
			Id        string `json:"id"`
			UpdatedAt string `json:"updatedAt"`
			Posts     []struct {
				Id      string `json:"id"`
				Title   string `json:"title"`
				Content string `json:"content"`
				Cursor  int    `json:"cursor"`
			} `json:"posts"`
		} `json:"feeds"`
	}

	client.MustPost(fmt.Sprintf(`
	query{
		feeds (input : {
		  userId : "%s"
		  feedRefreshInputs : [
			{feedId: "%s", limit: %d, cursor: %d, direction: %s}
		  ]
		}) {
		  id
		  updatedAt
		  posts {
			id
			title 
			content
			cursor
		  }
		}
	  }
	`, userId, feedId, 2, cursor, model.FeedRefreshDirectionOld), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	require.Equal(t, 1, len(resp.Feeds))
	require.Equal(t, feedId, resp.Feeds[0].Id)
	require.Equal(t, 2, len(resp.Feeds[0].Posts))
	require.Equal(t, "test_title_2", resp.Feeds[0].Posts[0].Title)
	require.Equal(t, "test_title_1", resp.Feeds[0].Posts[1].Title)
}

func checkFeedTopPostsMultipleFeeds(t *testing.T, userId string, feedIdOne string, feedIdTwo string, cursorOne int, cursorTwo int, db *gorm.DB, client *client.Client) {
	var resp struct {
		Feeds []struct {
			Id        string `json:"id"`
			UpdatedAt string `json:"updatedAt"`
			Posts     []struct {
				Id      string `json:"id"`
				Title   string `json:"title"`
				Content string `json:"content"`
				Cursor  int    `json:"cursor"`
			} `json:"posts"`
		} `json:"feeds"`
	}

	client.MustPost(fmt.Sprintf(`
	query{
		feeds (input : {
		  userId : "%s"
		  feedRefreshInputs : [
			{feedId: "%s", limit: %d, cursor: %d, direction: %s}
			{feedId: "%s", limit: %d, cursor: %d, direction: %s}
		  ]
		}) {
		  id
		  updatedAt
		  posts {
			id
			title 
			content
			cursor
		  }
		}
	  }
	`, userId, feedIdOne, 2, cursorOne, model.FeedRefreshDirectionNew, feedIdTwo, 2, cursorTwo, model.FeedRefreshDirectionNew), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	require.Equal(t, 2, len(resp.Feeds))
	require.Equal(t, feedIdOne, resp.Feeds[0].Id)
	require.Equal(t, 2, len(resp.Feeds[0].Posts))
	require.Equal(t, "test_title_6", resp.Feeds[0].Posts[0].Title)
	require.Equal(t, "test_title_5", resp.Feeds[0].Posts[1].Title)

	require.Equal(t, feedIdTwo, resp.Feeds[1].Id)
	require.Equal(t, 2, len(resp.Feeds[1].Posts))
	require.Equal(t, "test_title_6", resp.Feeds[1].Posts[0].Title)
	require.Equal(t, "test_title_5", resp.Feeds[1].Posts[1].Title)
}

func checkFeedBottomPostsMultipleFeeds(t *testing.T, userId string, feedIdOne string, feedIdTwo string, cursorOne int, cursorTwo int, db *gorm.DB, client *client.Client) {
	var resp struct {
		Feeds []struct {
			Id        string `json:"id"`
			UpdatedAt string `json:"updatedAt"`
			Posts     []struct {
				Id      string `json:"id"`
				Title   string `json:"title"`
				Content string `json:"content"`
				Cursor  int    `json:"cursor"`
			} `json:"posts"`
		} `json:"feeds"`
	}

	client.MustPost(fmt.Sprintf(`
	query{
		feeds (input : {
		  userId : "%s"
		  feedRefreshInputs : [
			{feedId: "%s", limit: %d, cursor: %d, direction: %s}
			{feedId: "%s", limit: %d, cursor: %d, direction: %s}
		  ]
		}) {
		  id
		  updatedAt
		  posts {
			id
			title 
			content
			cursor
		  }
		}
	  }
	`, userId, feedIdOne, 2, cursorOne, model.FeedRefreshDirectionOld, feedIdTwo, 2, cursorTwo, model.FeedRefreshDirectionOld), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	require.Equal(t, 2, len(resp.Feeds))
	require.Equal(t, feedIdOne, resp.Feeds[0].Id)
	require.Equal(t, 2, len(resp.Feeds[0].Posts))
	require.Equal(t, "test_title_2", resp.Feeds[0].Posts[0].Title)
	require.Equal(t, "test_title_1", resp.Feeds[0].Posts[1].Title)

	require.Equal(t, feedIdTwo, resp.Feeds[1].Id)
	require.Equal(t, 2, len(resp.Feeds[1].Posts))
	require.Equal(t, "test_title_2", resp.Feeds[1].Posts[0].Title)
	require.Equal(t, "test_title_1", resp.Feeds[1].Posts[1].Title)
}

func checkFeedTopPostsWithoutSpecifyFeed(t *testing.T, userId string, feedIdOne string, feedIdTwo string, db *gorm.DB, client *client.Client) {
	var resp struct {
		Feeds []struct {
			Id        string `json:"id"`
			UpdatedAt string `json:"updatedAt"`
			Posts     []struct {
				Id      string `json:"id"`
				Title   string `json:"title"`
				Content string `json:"content"`
				Cursor  int    `json:"cursor"`
			} `json:"posts"`
		} `json:"feeds"`
	}

	client.MustPost(fmt.Sprintf(`
	query{
		feeds (input : {
		  userId : "%s"
		  feedRefreshInputs : []
		}) {
		  id
		  updatedAt
		  posts {
			id
			title 
			content
			cursor
		  }
		}
	  }
	`, userId), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	require.Equal(t, 2, len(resp.Feeds))
	require.Equal(t, feedIdOne, resp.Feeds[0].Id)
	require.Equal(t, 7, len(resp.Feeds[0].Posts))
	require.Equal(t, "test_title_6", resp.Feeds[0].Posts[0].Title)
	require.Equal(t, "test_title_5", resp.Feeds[0].Posts[1].Title)
	require.Equal(t, "test_title_4", resp.Feeds[0].Posts[2].Title)
	require.Equal(t, "test_title_3", resp.Feeds[0].Posts[3].Title)
	require.Equal(t, "test_title_2", resp.Feeds[0].Posts[4].Title)
	require.Equal(t, "test_title_1", resp.Feeds[0].Posts[5].Title)
	require.Equal(t, "test_title_0", resp.Feeds[0].Posts[6].Title)

	require.Equal(t, feedIdTwo, resp.Feeds[1].Id)
	require.Equal(t, 7, len(resp.Feeds[1].Posts))
	require.Equal(t, "test_title_6", resp.Feeds[1].Posts[0].Title)
	require.Equal(t, "test_title_5", resp.Feeds[1].Posts[1].Title)
	require.Equal(t, "test_title_4", resp.Feeds[1].Posts[2].Title)
	require.Equal(t, "test_title_3", resp.Feeds[1].Posts[3].Title)
	require.Equal(t, "test_title_2", resp.Feeds[1].Posts[4].Title)
	require.Equal(t, "test_title_1", resp.Feeds[1].Posts[5].Title)
	require.Equal(t, "test_title_0", resp.Feeds[1].Posts[6].Title)
}
