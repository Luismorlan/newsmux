package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Luismorlan/newsmux/model"
	"github.com/Luismorlan/newsmux/server/graph/generated"
)

func (r *feedResolver) FilterDataExpression(ctx context.Context, obj *model.Feed) (string, error) {
	return string(obj.FilterDataExpression), nil
}

func (r *feedResolver) SubscriberCount(ctx context.Context, obj *model.Feed) (*int, error) {
	var count int64
	r.DB.Model(&model.UserFeedSubscription{}).
		Where("feed_id = ?", obj.Id).
		Count(&count)
	res := int(count)
	return &res, nil
}

// Feed returns generated.FeedResolver implementation.
func (r *Resolver) Feed() generated.FeedResolver { return &feedResolver{r} }

type feedResolver struct{ *Resolver }
