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
	"gorm.io/datatypes"
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
		utils.TestCreateUserAndValidate(t, "test_user_name", "test_user_id", db, client)
	})

	// Test no double creation for the same id
	var resp struct {
		CreateUser struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"createUser"`
	}
	client.MustPost(fmt.Sprintf(`mutation {
		createUser(input:{name:"%s" id: "%s"}) {
			id
			name
		}
	}`, "test_user_name_new", "test_user_id"), &resp)

	require.NotEmpty(t, resp.CreateUser.Id)
	require.Equal(t, resp.CreateUser.Id, "test_user_id")
	require.Equal(t, resp.CreateUser.Name, "test_user_name")
}

func TestCreateFeed(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test Feed Creation", func(t *testing.T) {
		uid := utils.TestCreateUserAndValidate(t, "test_user_name", "test_user_id", db, client)
		feedId, _ := utils.TestCreateFeedAndValidate(t, uid, "test_feed_for_feeds_api", `{\"a\":1}`, []string{}, db, client)
		require.NotEmpty(t, feedId)
	})
}

func TestCreateSource(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test Source Creation", func(t *testing.T) {
		uid := utils.TestCreateUserAndValidate(t, "test_user_name", "test_user_id", db, client)
		sourceId := utils.TestCreateSourceAndValidate(t, uid, "test_source_for_feeds_api", "test_domain", db, client)
		require.NotEmpty(t, sourceId)

		var source model.Source
		queryResult := db.Where("id = ?", sourceId).Preload("SubSources").First(&source)
		require.Equal(t, int64(1), queryResult.RowsAffected)
		require.Equal(t, 1, len(source.SubSources))
		require.Equal(t, DefaultSubSourceName, source.SubSources[0].Name)
	})
}

func TestCreateSubSource(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test Source Creation", func(t *testing.T) {
		uid := utils.TestCreateUserAndValidate(t, "test_user_name", "test_user_id", db, client)
		sourceId := utils.TestCreateSourceAndValidate(t, uid, "test_source_for_feeds_api", "test_domain", db, client)
		subSourceId := utils.TestCreateSubSourceAndValidate(t, uid, "test_subsource_for_feeds_api", "test_externalid", sourceId, db, client)
		require.NotEmpty(t, subSourceId)
	})
}

func TestUserSubscribeFeed(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	t.Run("Test User subscribe Feed", func(t *testing.T) {
		uid := utils.TestCreateUserAndValidate(t, "test_user_name", "test_user_id", db, client)
		feedId, _ := utils.TestCreateFeedAndValidate(t, uid, "test_feed_for_feeds_api", `{\"a\":1}`, []string{}, db, client)
		utils.TestUserSubscribeFeedAndValidate(t, uid, feedId, db, client)
	})
}

func TestQueryFeeds(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	userId := utils.TestCreateUserAndValidate(t, "test_user_for_feeds_api", "test_user_id", db, client)
	feedIdOne, updatedTimeOne := utils.TestCreateFeedAndValidate(t, userId, "test_feed_for_feeds_api", `{\"a\":1}`, []string{}, db, client)
	feedIdTwo, updatedTimeTwo := utils.TestCreateFeedAndValidate(t, userId, "test_feed_for_feeds_api", `{\"a\":1}`, []string{}, db, client)
	sourceId := utils.TestCreateSourceAndValidate(t, userId, "test_source_for_feeds_api", "test_domain", db, client)
	subSourceId := utils.TestCreateSubSourceAndValidate(t, userId, "test_source_for_feeds_api", "123123213123", sourceId, db, client)
	utils.TestCreateSubSourceAndValidate(t, userId, "test_subsource_for_feeds_api", "test_externalid", sourceId, db, client)
	utils.TestUserSubscribeFeedAndValidate(t, userId, feedIdOne, db, client)
	utils.TestUserSubscribeFeedAndValidate(t, userId, feedIdTwo, db, client)

	// 0 is oldest post, 6 is newest post
	utils.TestCreatePostAndValidate(t, "test_title_0", "test_content_0", subSourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_1", "test_content_1", subSourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_2", "test_content_2", subSourceId, feedIdOne, db, client)
	_, midCursorFirst := utils.TestCreatePostAndValidate(t, "test_title_3", "test_content_3", subSourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_4", "test_content_4", subSourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_5", "test_content_5", subSourceId, feedIdOne, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_6", "test_content_6", subSourceId, feedIdOne, db, client)

	// 0 is oldest post, 6 is newest post
	utils.TestCreatePostAndValidate(t, "test_title_0", "test_content_0", subSourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_1", "test_content_1", subSourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_2", "test_content_2", subSourceId, feedIdTwo, db, client)
	_, midCursorSecond := utils.TestCreatePostAndValidate(t, "test_title_3", "test_content_3", subSourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_4", "test_content_4", subSourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_5", "test_content_5", subSourceId, feedIdTwo, db, client)
	utils.TestCreatePostAndValidate(t, "test_title_6", "test_content_6", subSourceId, feedIdTwo, db, client)

	checkFeedPosts(t, userId, feedIdOne, midCursorFirst, 2, &updatedTimeOne, model.FeedRefreshDirectionNew,
		[]string{"test_content_5", "test_content_6"}, db, client)

	checkFeedPosts(t, userId, feedIdOne, midCursorSecond, 2, &updatedTimeOne, model.FeedRefreshDirectionOld,
		[]string{"test_content_2", "test_content_1"}, db, client)

	checkFeedTopPostsMultipleFeeds(t, userId, feedIdOne, feedIdTwo, midCursorFirst, midCursorSecond, updatedTimeOne, updatedTimeTwo, db, client)
	checkFeedBottomPostsMultipleFeeds(t, userId, feedIdOne, feedIdTwo, midCursorFirst, midCursorSecond, updatedTimeOne, updatedTimeTwo, db, client)
	checkFeedTopPostsWithoutSpecifyFeed(t, userId, feedIdOne, feedIdTwo, db, client)
	checkFeedTopPostsUpdateTimeChanged(t, userId, feedIdOne, midCursorFirst, "2021-08-24T21:57:15-07:00", db, client)
}

func checkFeedPosts(
	t *testing.T, userId string, feedId string, cursor int, limit int, updatedTime *string,
	direction model.FeedRefreshDirection, expectedPostsIds []string, db *gorm.DB, client *client.Client) {
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
	updatedTimeStr := `null`

	if updatedTime != nil {
		updatedTimeStr = fmt.Sprintf(`"%s"`, *updatedTime)
	}

	client.MustPost(fmt.Sprintf(`
	query{
		feeds (input : {
		  userId : "%s"
		  feedRefreshInputs : [
			{feedId: "%s", limit: %d, cursor: %d, direction: %s, feedUpdatedTime: %s}
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
	`, userId, feedId, limit, cursor, direction, updatedTimeStr), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	require.Equal(t, 1, len(resp.Feeds))
	require.Equal(t, feedId, resp.Feeds[0].Id)
	require.Equal(t, len(expectedPostsIds), len(resp.Feeds[0].Posts))

	var postIds []string
	for _, post := range resp.Feeds[0].Posts {
		postIds = append(postIds, post.Id)
	}
	utils.StringSlicesContainSameElements(postIds, expectedPostsIds)
}

func checkFeedTopPostsMultipleFeeds(
	t *testing.T, userId string, feedIdOne string, feedIdTwo string,
	cursorOne int, cursorTwo int, updatedTimeOne string, updatedTimeTwo string,
	db *gorm.DB, client *client.Client) {
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
			{feedId: "%s", limit: %d, cursor: %d, direction: %s, feedUpdatedTime: "%s"}
			{feedId: "%s", limit: %d, cursor: %d, direction: %s, feedUpdatedTime: "%s"}
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
	`, userId, feedIdOne, 2, cursorOne, model.FeedRefreshDirectionNew, updatedTimeOne,
		feedIdTwo, 2, cursorTwo, model.FeedRefreshDirectionNew, updatedTimeTwo), &resp)

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

func checkFeedBottomPostsMultipleFeeds(
	t *testing.T, userId string, feedIdOne string, feedIdTwo string,
	cursorOne int, cursorTwo int, updatedTimeOne string, updatedTimeTwo string,
	db *gorm.DB, client *client.Client) {
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
			{feedId: "%s", limit: %d, cursor: %d, direction: %s, feedUpdatedTime: "%s"}
			{feedId: "%s", limit: %d, cursor: %d, direction: %s, feedUpdatedTime: "%s"}
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
	`, userId, feedIdOne, 2, cursorOne, model.FeedRefreshDirectionOld, updatedTimeOne, feedIdTwo, 2, cursorTwo, model.FeedRefreshDirectionOld, updatedTimeTwo), &resp)

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

func checkFeedTopPostsUpdateTimeChanged(t *testing.T, userId string, feedId string, cursor int, wrongUpdatedTime string, db *gorm.DB, client *client.Client) {
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
				{feedId: "%s", limit: %d, cursor: %d, direction: %s, feedUpdatedTime: "%s"}
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
		`, userId, feedId, 7, cursor, model.FeedRefreshDirectionNew, wrongUpdatedTime), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	require.Equal(t, 1, len(resp.Feeds))
	require.Equal(t, feedId, resp.Feeds[0].Id)
	require.Equal(t, 7, len(resp.Feeds[0].Posts))
}

func TestUpSertFeeds(t *testing.T) {
	db, _ := utils.CreateTempDB(t)

	client := PrepareTestForGraphQLAPIs(db)

	userId := utils.TestCreateUserAndValidate(t, "test_user_for_feeds_api", "test_user_id", db, client)
	sourceId := utils.TestCreateSourceAndValidate(t, userId, "test_source_for_feeds_api", "test_domain", db, client)
	subSourceIdOne := utils.TestCreateSubSourceAndValidate(t, userId, "test_source_for_feeds_api", "1111", sourceId, db, client)
	subSourceIdTwo := utils.TestCreateSubSourceAndValidate(t, userId, "test_source_for_feeds_api_2", "2222", sourceId, db, client)
	feedIdOne, _ := utils.TestCreateFeedAndValidate(t, userId, "test_feed_for_feeds_api", ``, []string{}, db, client)

	postId1, _ := utils.TestCreatePostAndValidate(t, "test_title_1", "same_content_test", subSourceIdOne, feedIdOne, db, client)
	postId2, _ := utils.TestCreatePostAndValidate(t, "test_title_2", "test_content_2", subSourceIdOne, feedIdOne, db, client)

	postId3, _ := utils.TestCreatePostAndValidate(t, "test_title_3", "same_content_test", subSourceIdTwo, feedIdOne, db, client)
	postId4, _ := utils.TestCreatePostAndValidate(t, "test_title_4", "test_content_4", subSourceIdTwo, feedIdOne, db, client)
	postId5, cursor5 := utils.TestCreatePostAndValidate(t, "test_title_5", "test_content_5", subSourceIdTwo, feedIdOne, db, client)

	var feed model.Feed
	queryResult := db.Where("id = ?", feedIdOne).First(&feed)
	require.Equal(t, int64(1), queryResult.RowsAffected)

	// Re-publish when update subsource
	var subSourceOne model.SubSource
	db.Where("id = ?", subSourceIdOne).First(&subSourceOne)
	require.Equal(t, int64(1), queryResult.RowsAffected)
	feed.SubSources = []*model.SubSource{
		&subSourceOne,
	}

	_, posts := utils.TestUpdateFeedAndReturnPosts(t, feed, db, client)
	require.Equal(t, 0, len(posts))
	checkFeedPosts(t, userId, feedIdOne, 0, 999, nil, model.FeedRefreshDirectionOld,
		[]string{postId1, postId2}, db, client)

	// Re-publish when update to two subsources
	var subSourceTwo model.SubSource
	db.Where("id = ?", subSourceIdTwo).First(&subSourceTwo)
	require.Equal(t, int64(1), queryResult.RowsAffected)
	feed.SubSources = []*model.SubSource{
		&subSourceOne,
		&subSourceTwo,
	}
	updatedAt, posts := utils.TestUpdateFeedAndReturnPosts(t, feed, db, client)
	require.Equal(t, 0, len(posts))
	checkFeedPosts(t, userId, feedIdOne, 0, 999, nil, model.FeedRefreshDirectionOld,
		[]string{postId1, postId2, postId3, postId4, postId5}, db, client)

	// Re-publish when update data expression
	feed.FilterDataExpression = datatypes.JSON(
		`{
			"id":"1",
			"expr":{
				"pred":{
				"type":"LITERAL",
				"param":{
					"text":"same_content_test"
				}
				}
			}
	 	}`)
	updatedAt, posts = utils.TestUpdateFeedAndReturnPosts(t, feed, db, client)
	require.Equal(t, 0, len(posts))
	checkFeedPosts(t, userId, feedIdOne, 0, 999, nil, model.FeedRefreshDirectionOld,
		[]string{postId1, postId3}, db, client)

	// Re-publish when update data expression
	feed.FilterDataExpression = datatypes.JSON(``)
	updatedAt, posts = utils.TestUpdateFeedAndReturnPosts(t, feed, db, client)
	require.Equal(t, 0, len(posts))
	checkFeedPosts(t, userId, feedIdOne, 0, 1, nil, model.FeedRefreshDirectionOld,
		[]string{postId5}, db, client)
	// check only 1 post is published
	var count int64
	db.Model(&model.PostFeedPublish{}).Where("feed_id = ?", feedIdOne).Count(&count)
	require.Equal(t, int64(1), count)

	// publish more
	checkFeedPosts(t, userId, feedIdOne, cursor5, 2, &updatedAt, model.FeedRefreshDirectionOld,
		[]string{postId4, postId3}, db, client)
	// check only 3 post is published now after republishing
	db.Model(&model.PostFeedPublish{}).Where("feed_id = ?", feedIdOne).Count(&count)
	require.Equal(t, int64(3), count)
}
