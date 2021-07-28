package models

import "gorm.io/gorm"

// User is a data model for users
// UserId: a uuid generated to identify a user internally
// AuthId: generated by external authentication service "cognito"
// Name: user's name to display
type User struct {
	gorm.Model
	AuthenticationId string `json:"authentication_id"`
	Name             string
	Email            string
	SubscribedFeeds  []*Feed `gorm:"many2many:user_feed_subscription;"`
}
