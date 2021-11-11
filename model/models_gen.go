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

type AddWeiboSubSourceInput struct {
	Name string `json:"name"`
}

type DeleteFeedInput struct {
	UserID string `json:"userId"`
	FeedID string `json:"feedId"`
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

type FeedsGetPostsInput struct {
	UserID            string              `json:"userId"`
	FeedRefreshInputs []*FeedRefreshInput `json:"feedRefreshInputs"`
}

type NewPostInput struct {
	Title            string   `json:"title"`
	Content          string   `json:"content"`
	SubSourceID      string   `json:"subSourceId"`
	FeedsIDPublishTo []string `json:"feedsIdPublishTo"`
	SharedFromPostID *string  `json:"sharedFromPostId"`
}

type NewSourceInput struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type NewUserInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PostInFeedOutput struct {
	Post   *Post `json:"post"`
	Cursor int   `json:"cursor"`
}

type PostInput struct {
	ID string `json:"id"`
}

type SeedStateInput struct {
	UserSeedState *UserSeedStateInput   `json:"userSeedState"`
	FeedSeedState []*FeedSeedStateInput `json:"feedSeedState"`
}

type SourcesInput struct {
	SubSourceFromSharedPost bool `json:"subSourceFromSharedPost"`
}

type SubscribeInput struct {
	UserID string `json:"userId"`
	FeedID string `json:"feedId"`
}

type SubsourcesInput struct {
	IsFromSharedPost bool `json:"isFromSharedPost"`
}

type UpsertFeedInput struct {
	UserID               string     `json:"userId"`
	FeedID               *string    `json:"feedId"`
	Name                 string     `json:"name"`
	FilterDataExpression string     `json:"filterDataExpression"`
	SubSourceIds         []string   `json:"subSourceIds"`
	Visibility           Visibility `json:"visibility"`
}

type UpsertSubSourceInput struct {
	Name               string `json:"name"`
	ExternalIdentifier string `json:"externalIdentifier"`
	SourceID           string `json:"sourceId"`
	AvatarURL          string `json:"avatarUrl"`
	OriginURL          string `json:"originUrl"`
	IsFromSharedPost   bool   `json:"isFromSharedPost"`
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

type SignalType string

const (
	SignalTypeSeedState SignalType = "SEED_STATE"
)

var AllSignalType = []SignalType{
	SignalTypeSeedState,
}

func (e SignalType) IsValid() bool {
	switch e {
	case SignalTypeSeedState:
		return true
	}
	return false
}

func (e SignalType) String() string {
	return string(e)
}

func (e *SignalType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SignalType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SignalType", str)
	}
	return nil
}

func (e SignalType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Visibility string

const (
	VisibilityGlobal  Visibility = "GLOBAL"
	VisibilityPrivate Visibility = "PRIVATE"
)

var AllVisibility = []Visibility{
	VisibilityGlobal,
	VisibilityPrivate,
}

func (e Visibility) IsValid() bool {
	switch e {
	case VisibilityGlobal, VisibilityPrivate:
		return true
	}
	return false
}

func (e Visibility) String() string {
	return string(e)
}

func (e *Visibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Visibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Visibility", str)
	}
	return nil
}

func (e Visibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
