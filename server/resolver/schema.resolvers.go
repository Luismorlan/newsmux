package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/Luismorlan/newsmux/model"
	"github.com/Luismorlan/newsmux/server/graph/generated"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	uuid := uuid.New().String()

	t := model.User{
		Id:              uuid,
		Name:            input.Name,
		CreatedAt:       time.Now(),
		SubscribedFeeds: []*model.Feed{},
	}
	r.DB.Create(&t)
	r.DB.Save(&t)
	return &t, nil
}

func (r *mutationResolver) CreateFeed(ctx context.Context, input model.NewFeedInput) (*model.Feed, error) {
	uuid := uuid.New().String()

	t := model.Feed{
		Id:          uuid,
		Name:        input.Name,
		CreatedAt:   time.Now(),
		Subscribers: []*model.User{},
		Posts:       []*model.Post{},
	}
	r.DB.Create(&t)

	var user model.User
	r.DB.Where("id = ?", input.UserID).First(&user)
	r.DB.Model(&t).Association("Creator").Append(&user)

	r.DB.Save(&t)
	return &t, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPostInput) (*model.Post, error) {
	uuid := uuid.New().String()

	// TODO: clean up this logic
	var source model.Source
	r.DB.Where("id = ?", input.SourceID).First(&source)

	var subSource model.SubSource
	r.DB.Where("id = ?", input.SubSourceID).First(&subSource)

	var sourcePost model.Post
	r.DB.Where("id = ?", input.SharedFromPostID).First(&sourcePost)

	post := model.Post{
		Id:             uuid,
		Title:          input.Title,
		Content:        input.Content,
		CreatedAt:      time.Now(),
		Source:         source,
		SubSource:      &subSource,
		SharedFromPost: &sourcePost,
		SavedByUser:    []*model.User{},
		PublishedFeeds: []*model.Feed{},
	}
	r.DB.Create(&post)

	//TODO: test Publish post to feed
	for _, feedId := range input.FeedsIDPublishTo {
		err := r.DB.Transaction(func(tx *gorm.DB) error {
			var feed model.Feed
			r.DB.Where("id = ?", feedId).First(&feed)

			if e := r.DB.Model(&post).Association("PublishedFeeds").Append(&feed); e != nil {
				return e
			}
			if e := tx.Model(&feed).Association("Subscribers").Append(&post); e != nil {
				return e
			}
			// return nil will commit the whole transaction
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	r.DB.Save(&post)
	return &post, nil
}

func (r *mutationResolver) Subscribe(ctx context.Context, input model.SubscribeInput) (*model.User, error) {
	userId := input.UserID
	feedId := input.FeedID

	var user model.User
	var feed model.Feed

	result := r.DB.First(&user, "id = ?", userId)
	if result.RowsAffected != 1 {
		return nil, errors.New("No valid user found")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.DB.First(&feed, "id = ?", feedId)
	if result.RowsAffected != 1 {
		return nil, errors.New("No valid feed found")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&user).Association("SubscribedFeeds").Append(&feed); err != nil {
			return err
		}
		if err := tx.Model(&feed).Association("Subscribers").Append(&user); err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) CreateSource(ctx context.Context, input model.NewSourceInput) (*model.Source, error) {
	uuid := uuid.New().String()

	t := model.Source{
		Id:        uuid,
		Name:      input.Name,
		Domain:    input.Domain,
		CreatedAt: time.Now(),
	}
	r.DB.Create(&t)

	var user model.User
	r.DB.Where("id = ?", input.UserID).First(&user)
	r.DB.Model(&t).Association("Creator").Append(&user)

	r.DB.Save(&t)
	return &t, nil
}

func (r *mutationResolver) CreateSubSource(ctx context.Context, input model.NewSubSourceInput) (*model.SubSource, error) {
	uuid := uuid.New().String()

	t := model.SubSource{
		Id:                 uuid,
		Name:               input.Name,
		ExternalIdentifier: input.ExternalIdentifier,
		CreatedAt:          time.Now(),
		SourceID:           input.SourceID,
	}
	r.DB.Create(&t)

	var user model.User
	r.DB.Where("id = ?", input.UserID).First(&user)
	r.DB.Model(&t).Association("Creator").Append(&user)

	r.DB.Save(&t)
	return &t, nil
}

func (r *queryResolver) AllFeeds(ctx context.Context) ([]*model.Feed, error) {
	var feeds []*model.Feed
	result := r.DB.Preload(clause.Associations).Find(&feeds)
	return feeds, result.Error
}

func (r *queryResolver) Sources(ctx context.Context) ([]*model.Source, error) {
	var sources []*model.Source
	result := r.DB.Preload(clause.Associations).Find(&sources)
	return sources, result.Error
}

func (r *queryResolver) SubSources(ctx context.Context) ([]*model.SubSource, error) {
	var subSources []*model.SubSource
	result := r.DB.Preload(clause.Associations).Find(&subSources)
	return subSources, result.Error
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var posts []*model.Post
	result := r.DB.Preload(clause.Associations).Find(&posts)
	return posts, result.Error
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	result := r.DB.Preload(clause.Associations).Find(&users)
	return users, result.Error
}

func (r *queryResolver) Feeds(ctx context.Context, input *model.FeedsForUserInput) ([]*model.Feed, error) {
	// TODO: Here we always return all feed, should respect user input
	var feeds []*model.Feed
	result := r.DB.Preload(clause.Associations).Find(&feeds)
	return feeds, result.Error
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
