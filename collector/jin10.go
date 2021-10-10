package collector

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Luismorlan/newsmux/protocol"
	Logger "github.com/Luismorlan/newsmux/utils/log"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	jin10DateFormat = "20060102-15:04:05"
)

type Jin10Crawler struct {
	sink CollectedDataSink
}

func (collector Jin10Crawler) UpdateFileUrls(workingContext *CrawlerWorkingContext) error {
	return errors.New("UpdateFileUrls not implemented, should not be called")
}

func (collector Jin10Crawler) UpdateNewsType(workingContext *CrawlerWorkingContext) error {
	selection := workingContext.Element.DOM.Find(".jin-flash-item")
	if len(selection.Nodes) == 0 {
		workingContext.NewsType = protocol.PanopticSubSource_UNSPECIFIED
		return errors.New("Jin10 news item not found")
	}
	if selection.HasClass("is-important") {
		workingContext.NewsType = protocol.PanopticSubSource_KEYNEWS
		return nil
	}
	workingContext.NewsType = protocol.PanopticSubSource_FLASHNEWS
	return nil
}

func (collector Jin10Crawler) UpdateContent(workingContext *CrawlerWorkingContext) error {
	var sb strings.Builder
	selection := workingContext.Element.DOM.Find(".right-content > div")
	if len(selection.Nodes) == 0 {
		return errors.New("jin10 news DOM not found")
	}
	selection.Children().Each(func(_ int, s *goquery.Selection) {
		if len(s.Nodes) > 0 && s.Nodes[0].Data == "br" {
			sb.WriteString(" ")
		}
		sb.WriteString(s.Text())
	})

	// goquery don't have a good way to get text without child elements'
	// remove children's text manually
	remove := selection.Children().Text()
	text := selection.Text()
	result := strings.Replace(text, remove, "", -1)
	sb.WriteString(result)

	content := sb.String()
	workingContext.Result.Post.Content = content
	if len(content) == 0 {
		return errors.New("empty Content")
	}
	return nil
}

func (collector Jin10Crawler) UpdateGeneratedTime(workingContext *CrawlerWorkingContext) error {
	id := workingContext.Element.DOM.AttrOr("id", "")
	timeText := workingContext.Element.DOM.Find(".item-time").Text()
	if len(id) <= 13 {
		workingContext.Result.Post.ContentGeneratedAt = timestamppb.Now()
		return errors.New("Jin10 news DOM id length is smaller than expected")
	}

	dateStr := id[5:13] + "-" + timeText
	generatedTime, err := time.Parse(jin10DateFormat, dateStr)
	if err != nil {
		workingContext.Result.Post.ContentGeneratedAt = timestamppb.Now()
		return err
	}
	workingContext.Result.Post.ContentGeneratedAt = timestamppb.New(generatedTime.UTC())
	return nil
}

func (collector Jin10Crawler) UpdateExternalPostId(workingContext *CrawlerWorkingContext) error {
	id := workingContext.Element.DOM.AttrOr("id", "")
	if len(id) == 0 {
		return errors.New("can't get external post id for the news")
	}
	workingContext.ExternalPostId = id
	return nil
}

func (collector Jin10Crawler) UpdateDedupId(workingContext *CrawlerWorkingContext) error {
	hasher := md5.New()
	_, err := hasher.Write([]byte(workingContext.Task.TaskParams.SourceId + workingContext.ExternalPostId))
	if err != nil {
		return err
	}
	workingContext.Result.Post.DeduplicateId = hex.EncodeToString(hasher.Sum(nil))
	return nil
}

func (collector Jin10Crawler) UpdateImageUrls(workingContext *CrawlerWorkingContext) error {
	// there is only one image can be in jin10
	selection := workingContext.Element.DOM.Find(".img-container > img")
	workingContext.Result.Post.ImageUrls = []string{}
	if len(selection.Nodes) == 0 {
		return nil
	}

	imageUrl := selection.AttrOr("src", "")
	if len(imageUrl) == 0 {
		return errors.New("image DOM exist but src not found")
	}
	workingContext.Result.Post.ImageUrls = []string{imageUrl}
	return nil
}

// Process each html selection to get content
func (collector Jin10Crawler) IsRequested(workingContext *CrawlerWorkingContext) bool {
	requestedTypes := make(map[protocol.PanopticSubSource_SubSourceType]bool)

	for _, subsource := range workingContext.Task.TaskParams.SubSources {
		s := subsource
		requestedTypes[s.Type] = true
	}

	if _, ok := requestedTypes[workingContext.NewsType]; !ok {
		fmt.Println("Not requested, actual level: ", workingContext.NewsType, " requested ", requestedTypes)
		return false
	}

	return true
}

func (collector Jin10Crawler) UpdateSubsourceName(workingContext *CrawlerWorkingContext) error {
	workingContext.Result.Post.SubSource.Name = SubsourceTypeToName(workingContext.NewsType)
	return nil
}

func (collector Jin10Crawler) GetMessage(workingContext *CrawlerWorkingContext) error {
	workingContext.Result = &protocol.CrawlerMessage{Post: &protocol.CrawlerMessage_CrawledPost{}}
	// context per element crawl
	workingContext.Result.Post.SubSource = &protocol.CrawledSubSource{}
	workingContext.Result.Post.SubSource.SourceId = workingContext.Task.TaskParams.SourceId
	//todo: put in central place
	workingContext.Result.Post.SubSource.AvatarUrl = "https://newsfeed-logo.s3.us-west-1.amazonaws.com/jin10.png"
	workingContext.Result.CrawledAt = &timestamp.Timestamp{}
	workingContext.Result.CrawlerVersion = "1"
	workingContext.Result.IsTest = false
	workingContext.Result.Post.OriginUrl = workingContext.OriginUrl

	err := collector.UpdateContent(workingContext)
	if err != nil {
		return err
	}

	err = collector.UpdateExternalPostId(workingContext)
	if err != nil {
		return err
	}

	err = collector.UpdateDedupId(workingContext)
	if err != nil {
		return err
	}

	err = collector.UpdateNewsType(workingContext)
	if err != nil {
		return err
	}

	if !collector.IsRequested(workingContext) {
		return errors.New("Not requested level")
	}

	err = collector.UpdateGeneratedTime(workingContext)
	if err != nil {
		return err
	}

	err = collector.UpdateImageUrls(workingContext)
	if err != nil {
		return err
	}

	err = collector.UpdateSubsourceName(workingContext)
	if err != nil {
		return err
	}

	return nil
}

func (collector Jin10Crawler) GetQueryPath() string {
	return `#jin_flash_list > .jin-flash-item-container`
}

func (collector Jin10Crawler) GetStartUri() string {
	return "https://www.jin10.com/index.html"
}

// todo: mock http response and test end to end Collect()
func (collector Jin10Crawler) CollectAndPublish(task *protocol.PanopticTask) {
	metadata := task.TaskMetadata
	metadata.ResultState = protocol.TaskMetadata_STATE_SUCCESS

	c := colly.NewCollector()
	Logger.Log.Info("Starting crawl Jin10, Task ", task.String())
	// each crawled card(news) will go to this
	// for each page loaded, there are multiple calls into this func
	c.OnHTML(collector.GetQueryPath(), func(elem *colly.HTMLElement) {
		var (
			err error
		)
		workingContext := &CrawlerWorkingContext{Task: task, Element: elem, OriginUrl: collector.GetStartUri()}
		if err = collector.GetMessage(workingContext); err != nil {
			metadata.TotalMessageFailed++
			LogHtmlParsingError(task, elem, err)
			return
		}
		if err = collector.sink.Push(workingContext.Result); err != nil {
			task.TaskMetadata.ResultState = protocol.TaskMetadata_STATE_FAILURE
			metadata.TotalMessageFailed++
			Logger.Log.Errorf("fail to publish message %s to SNS. Task: %s", workingContext.Result.String(), task.String())
			return
		}
		metadata.TotalMessageCollected++
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		// todo: error should be put into metadata
		task.TaskMetadata.ResultState = protocol.TaskMetadata_STATE_FAILURE
		Logger.Log.Error("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err, " path ", collector.GetQueryPath())
	})

	c.OnResponse(func(_ *colly.Response) {
		Logger.Log.Info("Finished crawl one page for Jin10, Task ", task.String())
	})

	c.OnScraped(func(_ *colly.Response) {
		if task.TaskMetadata.TotalMessageCollected == 0 {
			task.TaskMetadata.ResultState = protocol.TaskMetadata_STATE_FAILURE
			Logger.Log.Error("Finished crawl Jin10 with 0 success msg, Task ", task.String(), " failCount ", task.TaskMetadata.TotalMessageFailed, " path ", collector.GetQueryPath())
			return
		}
		if task.TaskMetadata.TotalMessageFailed > 0 {
			task.TaskMetadata.ResultState = protocol.TaskMetadata_STATE_FAILURE
		}
	})

	c.Visit(collector.GetStartUri())
	return
}
