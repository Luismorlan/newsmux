package validation

import (
	"github.com/Luismorlan/newsmux/collector/working_context"
	"github.com/Luismorlan/newsmux/protocol"
	"github.com/pkg/errors"
)

// Intentionally duplicate the code in collector_utils.go so that there's no
// interference between the validation code and the generation code.
const (
	jin10SourceId          = "a882eb0d-0bde-401a-b708-a7ce352b7392"
	weiboSourceId          = "0129417c-4987-45c9-86ac-d6a5c89fb4f7"
	kuailansiSourceId      = "6e1f6734-985b-4a52-865f-fc39a9daa2e8"
	wallstreetNewsSourceId = "66251821-ef9a-464c-bde9-8b2fd8ef2405"
	JinseSourceId          = "5891f435-d51e-4575-b4af-47cd4ede5607"
)

func getSourceLogoUrl(sourceId string) string {
	switch sourceId {
	case jin10SourceId:
		return "https://newsfeed-logo.s3.us-west-1.amazonaws.com/jin10.png"
	case weiboSourceId:
		return ""
	case wallstreetNewsSourceId:
		return "https://newsfeed-logo.s3.us-west-1.amazonaws.com/wallstrt.png"
	case kuailansiSourceId:
		return "https://newsfeed-logo.s3.us-west-1.amazonaws.com/kuailansi.png"
	case JinseSourceId:
		return "https://newsfeed-logo.s3.us-west-1.amazonaws.com/jinse.png"
	default:
		return ""
	}
}

func getSourceIdFromDataCollectorId(collectorId protocol.PanopticTask_DataCollectorId) string {
	switch collectorId {
	case protocol.PanopticTask_COLLECTOR_JINSHI:
		return jin10SourceId
	case protocol.PanopticTask_COLLECTOR_KUAILANSI:
		return kuailansiSourceId
	case protocol.PanopticTask_COLLECTOR_WEIBO:
		return weiboSourceId
	case protocol.PanopticTask_COLLECTOR_WALLSTREET_NEWS:
		return wallstreetNewsSourceId
	case protocol.PanopticTask_COLLECTOR_JINSE:
		return JinseSourceId
	default:
		return ""
	}
}

// Validate a message before it get published to sink.
// Validator applied only to the shared context, where it contains the task to
// be returned back to Panoptic, as well as the crawled messages.
// A non valid shared context must not be pushed to sink.
func ValidateSharedContext(sharedContext *working_context.SharedContext) error {
	validators := []func(*working_context.SharedContext) error{
		crawledMessageValidation,
		panopticTaskValidation,
		crossTaskMessageValidation,
	}

	for _, v := range validators {
		if err := v(sharedContext); err != nil {
			return errors.Wrap(err, sharedContext.String())
		}
	}

	return nil
}

// Validate CrawledMessage is set correctly before pushing to sink. This type of
// validation only looks at CrawledMessage without any context about the
// original task itself. It's a stateless validation.
func crawledMessageValidation(sharedContext *working_context.SharedContext) error {
	messageValidators := []func(*protocol.CrawlerMessage) error{
		validateMessageSubSourceIsSetCorrectly,
		validateMessagePostIsSetCorrectly,
		validateMessageMetadataIsSetCorrectly,
	}
	for _, v := range messageValidators {
		if err := v(sharedContext.Result); err != nil {
			return err
		}
	}
	return nil
}

// Validate PanopticTask is set correctly before pushing to sink. This type of
// validation only looks at PanopticTask without any context about the
// original task itself. It's a stateless validation.
func panopticTaskValidation(sharedContext *working_context.SharedContext) error {
	taskValidators := []func(*protocol.PanopticTask) error{
		validateTaskMetadataIsSetCorrectly,
	}
	for _, v := range taskValidators {
		if err := v(sharedContext.Task); err != nil {
			return err
		}
	}
	return nil
}

// Validate CrawledMessage indeed matches the task specification.
func crossTaskMessageValidation(sharedContext *working_context.SharedContext) error {
	task := sharedContext.Task
	msg := sharedContext.Result

	if msg.CrawlerIp != task.TaskMetadata.IpAddr {
		return errors.New("crawled message mismatch task's IP address")
	}

	if msg.Post.SubSource.SourceId != task.TaskParams.SourceId {
		return errors.New("crawled message mismatch task's source id")
	}

	defaultUrl := getSourceLogoUrl(task.TaskParams.SourceId)
	if defaultUrl != "" && msg.Post.SubSource.AvatarUrl != defaultUrl {
		return errors.New("crawled message's avatar doesn't match the source's default avatar url: " + defaultUrl)
	}

	if msg.Post.SubSource.SourceId != getSourceIdFromDataCollectorId(task.DataCollectorId) {
		return errors.New("crawled message's source id doesn't match the data collector id")
	}

	return nil
}

// SubSource on Post is set correctly. A subsource is valid iff:
// - Has AvatarUrl
// - Has SourceId
// - Has SubSourceId
func validateMessageSubSourceIsSetCorrectly(msg *protocol.CrawlerMessage) error {
	if msg.Post.SubSource.AvatarUrl == "" {
		return errors.New("crawled post must have avatar url")
	}

	if msg.Post.SubSource.Name == "" {
		return errors.New("crawled post must have name")
	}

	if msg.Post.SubSource.SourceId == "" {
		return errors.New("crawled post must be associated with a source id")
	}

	return nil
}

// A Post is valid iff:
// - It has content
// - It is associated with a generated time to render correct timestamp
// - It has a deduplicateId
func validateMessagePostIsSetCorrectly(msg *protocol.CrawlerMessage) error {
	if msg.Post.Content == "" {
		return errors.New("crawled post must have Content at least")
	}

	if msg.Post.ContentGeneratedAt == nil {
		return errors.New("crawled post must be associated with a generated time")
	}

	if msg.Post.DeduplicateId == "" {
		return errors.New("crawled post must have a deduplicateId")
	}

	return nil
}

// A message's metadata is valid iff:
// - It has crawler ip
// - It is associated with a crawled time
func validateMessageMetadataIsSetCorrectly(msg *protocol.CrawlerMessage) error {
	if msg.CrawlerIp == "" {
		return errors.New("crawled message must be associated with an IP address")
	}

	if msg.CrawledAt == nil {
		return errors.New("crawled message must be associated with a crawled time")
	}

	return nil
}

// A PanopticTask's metadata is set correctly iff:
// - It must have IP address of the Lambda
// - It must have the config name that generated it
// - It must be associated with an end state, and the state is correct
func validateTaskMetadataIsSetCorrectly(task *protocol.PanopticTask) error {
	if task.TaskMetadata.ConfigName == "" {
		return errors.New("PanopticTask must have config name populated")
	}

	if task.TaskMetadata.IpAddr == "" {
		return errors.New("PanopticTask must have IP address populated")
	}

	if task.TaskMetadata.ResultState == protocol.TaskMetadata_STATE_UNSPECIFIED {
		return errors.New("PanopticTask must be associated with a result state")
	}

	if task.TaskMetadata.TotalMessageFailed > 0 &&
		task.TaskMetadata.ResultState == protocol.TaskMetadata_STATE_SUCCESS {
		return errors.New("PanopticTask must be at failure state if it has non zero failure message")
	}

	return nil
}