// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type FeedSeedStateInterface interface {
	IsFeedSeedStateInterface()
}

type UserSeedStateInterface interface {
	IsUserSeedStateInterface()
}

type FeedRefreshInput struct {
	FeedID          string               `json:"feedId"`
	Limit           int                  `json:"limit"`
	Cursor          int                  `json:"cursor"`
	Direction       FeedRefreshDirection `json:"direction"`
	FeedUpdatedTime *time.Time           `json:"feedUpdatedTime"`
}

type FeedSeedState struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (FeedSeedState) IsFeedSeedStateInterface() {}

type FeedSeedStateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FeedsForUserInput struct {
	UserID            string              `json:"userId"`
	FeedRefreshInputs []*FeedRefreshInput `json:"feedRefreshInputs"`
}

type NewFeedInput struct {
	UserID        string  `json:"userId"`
	Name          string  `json:"name"`
	FilterSetting *string `json:"filterSetting"`
}

type NewPostInput struct {
	Title            string   `json:"title"`
	Content          string   `json:"content"`
	SourceID         string   `json:"sourceId"`
	SubSourceID      *string  `json:"subSourceId"`
	FeedsIDPublishTo []string `json:"feedsIdPublishTo"`
	SharedFromPostID *string  `json:"sharedFromPostId"`
}

type NewSourceInput struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type NewSubSourceInput struct {
	UserID             string `json:"userId"`
	Name               string `json:"name"`
	ExternalIdentifier string `json:"externalIdentifier"`
	SourceID           string `json:"sourceId"`
}

type NewUserInput struct {
	Name string `json:"name"`
}

type PostInFeedOutput struct {
	Post   *Post `json:"post"`
	Cursor int   `json:"cursor"`
}

type SeedStateInput struct {
	UserSeedState *UserSeedStateInput   `json:"userSeedState"`
	FeedSeedState []*FeedSeedStateInput `json:"feedSeedState"`
}

type SubscribeInput struct {
	UserID string `json:"userId"`
	FeedID string `json:"feedId"`
}

type UserSeedState struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

func (UserSeedState) IsUserSeedStateInterface() {}

type UserSeedStateInput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

type FeedRefreshDirection string

const (
	FeedRefreshDirectionNew FeedRefreshDirection = "NEW"
	FeedRefreshDirectionOld FeedRefreshDirection = "OLD"
)

var AllFeedRefreshDirection = []FeedRefreshDirection{
	FeedRefreshDirectionNew,
	FeedRefreshDirectionOld,
}

func (e FeedRefreshDirection) IsValid() bool {
	switch e {
	case FeedRefreshDirectionNew, FeedRefreshDirectionOld:
		return true
	}
	return false
}

func (e FeedRefreshDirection) String() string {
	return string(e)
}

func (e *FeedRefreshDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FeedRefreshDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FeedRefreshDirection", str)
	}
	return nil
}

func (e FeedRefreshDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
