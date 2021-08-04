package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/Luismorlan/newsmux/model"
	"github.com/Luismorlan/newsmux/server/graph/generated"
)

func (r *feedResolver) DeletedAt(ctx context.Context, obj *model.Feed) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Feed returns generated.FeedResolver implementation.
func (r *Resolver) Feed() generated.FeedResolver { return &feedResolver{r} }

type feedResolver struct{ *Resolver }
