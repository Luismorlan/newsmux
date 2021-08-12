package resolver

import (
	"fmt"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Luismorlan/newsmux/model"
	"github.com/Luismorlan/newsmux/server/graph/generated"
	"github.com/Luismorlan/newsmux/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	t.Run("Test User Creation", func(t *testing.T) {
		createUserAndValidate(t, "test_user_name")
	})
}

func TestCreateFeed(t *testing.T) {
	t.Run("Test Feed Creation", func(t *testing.T) {
		uid := createUserAndValidate(t, "test_user_name")
		feedId := createFeedAndValidate(t, uid, "test_feed_for_feeds_api")
		require.NotEmpty(t, feedId)
	})
}

func TestCreateSource(t *testing.T) {
	t.Run("Test Source Creation", func(t *testing.T) {
		uid := createUserAndValidate(t, "test_user_name")
		sourceId := createSourceAndValidate(t, uid, "test_source_for_feeds_api", "test_domain")
		require.NotEmpty(t, sourceId)
	})
}

func TestCreateSubSource(t *testing.T) {
	t.Run("Test Source Creation", func(t *testing.T) {
		uid := createUserAndValidate(t, "test_user_name")
		sourceId := createSourceAndValidate(t, uid, "test_source_for_feeds_api", "test_domain")
		subSourceId := createSubSourceAndValidate(t, uid, "test_subsource_for_feeds_api", "test_externalid", sourceId)
		require.NotEmpty(t, subSourceId)
	})
}

func TestUserSubscribeFeed(t *testing.T) {
	t.Run("Test User subscribe Feed", func(t *testing.T) {
		uid := createUserAndValidate(t, "test_user_name")
		feedId := createFeedAndValidate(t, uid, "test_feed_for_feeds_api")
		userSubscribeFeedAndValidate(t, uid, feedId)
	})
}

func TestQueryFeeds(t *testing.T) {

	userId := createUserAndValidate(t, "test_user_for_feeds_api")
	feedIdOne := createFeedAndValidate(t, userId, "test_feed_for_feeds_api")
	feedIdTwo := createFeedAndValidate(t, userId, "test_feed_for_feeds_api")
	sourceId := createSourceAndValidate(t, userId, "test_source_for_feeds_api", "test_domain")
	createSubSourceAndValidate(t, userId, "test_subsource_for_feeds_api", "test_externalid", sourceId)
	userSubscribeFeedAndValidate(t, userId, feedIdOne)
	userSubscribeFeedAndValidate(t, userId, feedIdTwo)

	// 0 is oldest post, 6 is newest post
	createPostAndValidate(t, "test_title_0", "test_content_0", sourceId, feedIdOne)
	createPostAndValidate(t, "test_title_1", "test_content_1", sourceId, feedIdOne)
	createPostAndValidate(t, "test_title_2", "test_content_2", sourceId, feedIdOne)
	_, midCursorFirst := createPostAndValidate(t, "test_title_3", "test_content_3", sourceId, feedIdOne)
	createPostAndValidate(t, "test_title_4", "test_content_4", sourceId, feedIdOne)
	createPostAndValidate(t, "test_title_5", "test_content_5", sourceId, feedIdOne)
	createPostAndValidate(t, "test_title_6", "test_content_6", sourceId, feedIdOne)

	// 0 is oldest post, 6 is newest post
	createPostAndValidate(t, "test_title_0", "test_content_0", sourceId, feedIdTwo)
	createPostAndValidate(t, "test_title_1", "test_content_1", sourceId, feedIdTwo)
	createPostAndValidate(t, "test_title_2", "test_content_2", sourceId, feedIdTwo)
	_, midCursorSecond := createPostAndValidate(t, "test_title_3", "test_content_3", sourceId, feedIdTwo)
	createPostAndValidate(t, "test_title_4", "test_content_4", sourceId, feedIdTwo)
	createPostAndValidate(t, "test_title_5", "test_content_5", sourceId, feedIdTwo)
	createPostAndValidate(t, "test_title_6", "test_content_6", sourceId, feedIdTwo)

	checkFeedTopPosts(t, userId, feedIdOne, midCursorFirst)
	checkFeedBottomPosts(t, userId, feedIdOne, midCursorFirst)

	checkFeedTopPostsMultipleFeeds(t, userId, feedIdOne, feedIdTwo, midCursorFirst, midCursorSecond)
	checkFeedBottomPostsMultipleFeeds(t, userId, feedIdOne, feedIdTwo, midCursorFirst, midCursorSecond)

	checkFeedTopPostsWithoutSpecifyFeed(t, userId, feedIdOne, feedIdTwo)
}

func checkFeedTopPosts(t *testing.T, userId string, feedId string, cursor int) {
	client := prepareTestForGraphQLAPIs()
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

func checkFeedBottomPosts(t *testing.T, userId string, feedId string, cursor int) {
	client := prepareTestForGraphQLAPIs()
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

func checkFeedTopPostsMultipleFeeds(t *testing.T, userId string, feedIdOne string, feedIdTwo string, cursorOne int, cursorTwo int) {
	client := prepareTestForGraphQLAPIs()
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

func checkFeedBottomPostsMultipleFeeds(t *testing.T, userId string, feedIdOne string, feedIdTwo string, cursorOne int, cursorTwo int) {
	client := prepareTestForGraphQLAPIs()
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

func checkFeedTopPostsWithoutSpecifyFeed(t *testing.T, userId string, feedIdOne string, feedIdTwo string) {
	client := prepareTestForGraphQLAPIs()
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

func prepareTestForGraphQLAPIs() *client.Client {
	db, err := utils.GetDBLocalTest()
	db = db.Debug()
	utils.DatabaseSetupAndMigration(db)
	if err != nil {
		panic("failed to connect database")
	}

	client := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
		DB:             db,
		SeedStateChans: NewSeedStateChannels(),
	}})))
	return client
}

// create user with name, do sanity checks and returns its Id
func createUserAndValidate(t *testing.T, name string) (id string) {
	client := prepareTestForGraphQLAPIs()

	var resp struct {
		CreateUser struct {
			Id         string `json:"id"`
			Name       string `json:"name"`
			CreatedAt  string `json:"createdAt"`
			DeletedAt  string `json:"deletedAt"`
			SavedPosts []struct {
				Id string `json:"id"`
			}
			SubscribedFeeds []struct {
				Id string `json:"id"`
			}
		} `json:"createUser"`
	}

	client.MustPost(fmt.Sprintf(`mutation {
		createUser(input:{name:"%s"}) {
		  id
		  name
		  createdAt
		  deletedAt
		  savedPosts {
			  id
		  }
		  subscribedFeeds {
			  id
		  }
		}
	  }
	  `, name), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	createTime, _ := time.Parse("2021-08-08T14:32:50-07:00", resp.CreateUser.CreatedAt)

	require.NotEmpty(t, resp.CreateUser.Id)
	require.Equal(t, name, resp.CreateUser.Name)
	require.Equal(t, 0, len(resp.CreateUser.SavedPosts))
	require.Equal(t, 0, len(resp.CreateUser.SubscribedFeeds))
	require.Truef(t, time.Now().UnixNano() > createTime.UnixNano(), "time created wrong")
	require.Equal(t, "", resp.CreateUser.DeletedAt)

	return resp.CreateUser.Id
}

// create feed with name, do sanity checks and returns its Id
func createFeedAndValidate(t *testing.T, userId string, name string) (id string) {
	client := prepareTestForGraphQLAPIs()

	var resp struct {
		CreateFeed struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			CreatedAt string `json:"createdAt"`
			DeletedAt string `json:"deletedAt"`
		} `json:"createFeed"`
	}

	client.MustPost(fmt.Sprintf(`mutation {
		createFeed(input:{userId:"%s" name:"%s"}) {
		  id
		  name
		  createdAt
		  deletedAt
		}
	  }
	  `, userId, name), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	createTime, _ := time.Parse("2021-08-08T14:32:50-07:00", resp.CreateFeed.CreatedAt)

	require.NotEmpty(t, resp.CreateFeed.Id)
	require.Equal(t, name, resp.CreateFeed.Name)
	require.Truef(t, time.Now().UnixNano() > createTime.UnixNano(), "time created wrong")
	require.Equal(t, "", resp.CreateFeed.DeletedAt)

	return resp.CreateFeed.Id
}

// create source with name, do sanity checks and returns its Id
func createSourceAndValidate(t *testing.T, userId string, name string, domain string) (id string) {
	client := prepareTestForGraphQLAPIs()

	var resp struct {
		CreateSource struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			Domain    string `json:"domain"`
			CreatedAt string `json:"createdAt"`
			DeletedAt string `json:"deletedAt"`
		} `json:"createSource"`
	}

	client.MustPost(fmt.Sprintf(`mutation {
		createSource(input:{userId:"%s" name:"%s" domain:"%s"}) {
		  id
		  name
		  domain
		  createdAt
		  deletedAt
		}
	  }
	  `, userId, name, domain), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)
	createTime, _ := time.Parse("2021-08-08T14:32:50-07:00", resp.CreateSource.CreatedAt)

	require.NotEmpty(t, resp.CreateSource.Id)
	require.Equal(t, name, resp.CreateSource.Name)
	require.Equal(t, domain, resp.CreateSource.Domain)
	require.Truef(t, time.Now().UnixNano() > createTime.UnixNano(), "time created wrong")
	require.Equal(t, "", resp.CreateSource.DeletedAt)

	return resp.CreateSource.Id
}

// create subsource with name, do sanity checks and returns its Id
func createSubSourceAndValidate(t *testing.T, userId string, name string, externalIdentifier string, sourceId string) (id string) {
	client := prepareTestForGraphQLAPIs()

	var resp struct {
		CreateSubSource struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			CreatedAt string `json:"createdAt"`
			DeletedAt string `json:"deletedAt"`
		} `json:"createSubSource"`
	}

	client.MustPost(fmt.Sprintf(`mutation {
		createSubSource(input:{userId:"%s" name:"%s" externalIdentifier:"%s" sourceId:"%s"}) {
		  id
		  name
		  createdAt
		  deletedAt
		}
	  }
	  `, userId, name, externalIdentifier, sourceId), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	createTime, _ := time.Parse("2021-08-08T14:32:50-07:00", resp.CreateSubSource.CreatedAt)

	require.NotEmpty(t, resp.CreateSubSource.Id)
	require.Equal(t, name, resp.CreateSubSource.Name)
	require.Truef(t, time.Now().UnixNano() > createTime.UnixNano(), "time created wrong")
	require.Equal(t, "", resp.CreateSubSource.DeletedAt)

	return resp.CreateSubSource.Id
}

// create subsource with title,content, do sanity checks and returns its Id
func createPostAndValidate(t *testing.T, title string, content string, sourceId string, publishFeedId string) (id string, cursor int) {
	client := prepareTestForGraphQLAPIs()

	var resp struct {
		CreatePost struct {
			Id        string `json:"id"`
			Title     string `json:"title"`
			Content   string `json:"content"`
			Cursor    int    `json:"cursor"`
			CreatedAt string `json:"createdAt"`
			DeletedAt string `json:"deletedAt"`
		} `json:"createPost"`
	}

	client.MustPost(fmt.Sprintf(`mutation {
		createPost(
			input: {
				title: "%s"
				content: "%s"
				sourceId: "%s"
				feedsIdPublishTo: ["%s"]
			}
		) {
		  id
		  title
		  content
		  cursor
		  createdAt
		  deletedAt
		}
	  }
	  `, title, content, sourceId, publishFeedId), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	createTime, _ := time.Parse("2021-08-08T14:32:50-07:00", resp.CreatePost.CreatedAt)

	require.NotEmpty(t, resp.CreatePost.Id)
	require.Equal(t, title, resp.CreatePost.Title)
	require.Equal(t, content, resp.CreatePost.Content)
	require.Truef(t, time.Now().UnixNano() > createTime.UnixNano(), "time created wrong")
	require.Equal(t, "", resp.CreatePost.DeletedAt)

	return resp.CreatePost.Id, resp.CreatePost.Cursor
}

// create user to feed subscription, do sanity checks
func userSubscribeFeedAndValidate(t *testing.T, userId string, feedId string) {
	client := prepareTestForGraphQLAPIs()

	var resp struct {
		Subscribe struct {
			Id string `json:"id"`
		} `json:"subscribe"`
	}

	client.MustPost(fmt.Sprintf(`mutation {
		subscribe(
			input: {
				userId: "%s"
				feedId: "%s"
			}
		) {
		  id
		}
	  }
	  `, userId, feedId), &resp)

	fmt.Printf("\nResponse from resolver: %+v\n", resp)

	require.Equal(t, userId, resp.Subscribe.Id)
}