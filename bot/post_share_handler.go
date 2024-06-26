package bot

// This handler is for slack slash commands
// https://api.slack.com/interactivity/slash-commands

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Luismorlan/newsmux/model"
	Logger "github.com/Luismorlan/newsmux/utils/log"
)

const (
	SimilarityThreshold   = 37
	SimilarityWindowHours = 1
)

var PostsSent map[string][]postMeta
var Mutex sync.Mutex

func init() {
	PostsSent = map[string][]postMeta{}
}

type postMeta struct {
	Id                 string `json:"id"`
	SemanticHash       string `json:"semantic_hash"`
	ContentGeneratedAt time.Time
}

type SharePostPayload struct {
	model.Post
	WebhookUrl string `json:"webhook_url"`
}

func isHashingSemanticallyIdentical(h1 string, h2 string) bool {
	// If the hashing is invalid, or not of same length, they cannot be considered
	// as the semantically identical.
	if h1 == "" || h2 == "" || len(h1) != len(h2) {
		return false
	}

	// Calculate hamming distance by counting how many different bits in total.
	count := 0
	for idx := 0; idx < len(h1); idx++ {
		if h1[idx] != h2[idx] {
			count++
		}
	}

	return count <= SimilarityThreshold
}

func isPostDuplicated(
	post model.Post,
	channelId string,
) bool {
	Mutex.Lock()
	defer Mutex.Unlock()

	_, ok := PostsSent[channelId]
	if !ok {
		return false
	}

	for i := len(PostsSent[channelId]) - 1; i >= 0; i-- {
		p := PostsSent[channelId][i]
		// the collector has some interval(up to 12 hours for zsxq) to collect the data
		// we will keep the cache for two days
		if math.Abs(time.Since(p.ContentGeneratedAt).Hours()) > 48 {
			PostsSent[channelId] = append(PostsSent[channelId][:i], PostsSent[channelId][i+1:]...)
		}

		if post.SemanticHashing == "" ||
			p.SemanticHash == "" {
			return false
		}

		if (math.Abs(post.ContentGeneratedAt.Sub(p.ContentGeneratedAt).Hours())) < SimilarityWindowHours {
			if isHashingSemanticallyIdentical(post.SemanticHashing, p.SemanticHash) {
				return true
			}
		}
	}

	return false
}

func parsePostSharePayload(body io.ReadCloser) (*SharePostPayload, error) {
	bodybytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	payload := SharePostPayload{}

	err = json.Unmarshal(bodybytes, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func PostShareHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := parsePostSharePayload(c.Request.Body)
		if err != nil {
			bodybytes, _ := ioutil.ReadAll(c.Request.Body)
			Logger.Log.Error("invalid post share payload", err, string(bodybytes))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}

		if isPostDuplicated(payload.Post, payload.WebhookUrl) {
			c.Data(202, "application/json; charset=utf-8", []byte("Post duplicated"))
			return
		}

		if err := PushPostViaWebhook(payload.Post, payload.WebhookUrl); err != nil {
			Logger.Log.Error("Fail to post via webhook", err)
		}

		Mutex.Lock()
		defer Mutex.Unlock()

		if posts, ok := PostsSent[payload.WebhookUrl]; ok {
			PostsSent[payload.WebhookUrl] = append(posts,
				postMeta{
					Id:                 payload.Post.Id,
					SemanticHash:       payload.Post.SemanticHashing,
					ContentGeneratedAt: payload.Post.ContentGeneratedAt,
				})
		} else {
			PostsSent[payload.WebhookUrl] = []postMeta{
				{
					Id:                 payload.Post.Id,
					SemanticHash:       payload.Post.SemanticHashing,
					ContentGeneratedAt: payload.Post.ContentGeneratedAt,
				},
			}
		}

		c.Data(200, "application/json; charset=utf-8", []byte("Post sent"))
	}
}
