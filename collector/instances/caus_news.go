package collector_instances

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/Luismorlan/newsmux/collector"
	"github.com/Luismorlan/newsmux/collector/sink"
	"github.com/Luismorlan/newsmux/collector/working_context"
	"github.com/Luismorlan/newsmux/protocol"
	"github.com/Luismorlan/newsmux/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CaUsNewsCrawler struct {
	Sink sink.CollectedDataSink
}

type CaUsNewsResponseItem struct {
	ContentID   int    `json:"contentId"`
	Title       string `json:"title"`
	PublishTime int64  `json:"publishTime"`
	Content     string `json:"content"`
	Lanmus      []struct {
		ID            int         `json:"id"`
		Name          string      `json:"name"`
		Description   interface{} `json:"description"`
		CreateTime    int64       `json:"createTime"`
		CreateTimeStr interface{} `json:"createTimeStr"`
	} `json:"lanmus"`
	Type       string `json:"type"`
	CreateTime int64  `json:"createTime"`
	CountLike  int    `json:"countLike"`
}

type CaUsNewsResponse struct {
	CurrentTime     int64  `json:"currentTime"`
	ErrorCode       int    `json:"errorCode"`
	APIErrorMessage string `json:"apiErrorMessage"`
	Data            struct {
		ArticleList []CaUsNewsResponseItem `json:"articleList"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (caus CaUsNewsCrawler) ConstructUrl() string {
	return "https://api.caus.money/toronto/display/lanmuArticlelistNew"
}

func (caus CaUsNewsCrawler) UpdateDedupId(post *protocol.CrawlerMessage_CrawledPost) error {
	md5, err := utils.TextToMd5Hash(post.Content)
	if err != nil {
		return utils.ImmediatePrintError(err)
	}
	post.DeduplicateId = md5
	return nil
}

func (caus CaUsNewsCrawler) UpdateResult(wc *working_context.ApiCollectorWorkingContext) error {
	item := wc.ApiResponseItem.(CaUsNewsResponseItem)
	post := wc.Result.Post
	generatedTime := time.UnixMilli(item.PublishTime)
	post.ContentGeneratedAt = timestamppb.New(generatedTime)

	post.OriginUrl = ""

	post.SubSource.Name = wc.SubSource.Name
	post.SubSource.ExternalId = fmt.Sprint(item.ContentID)

	post.Content = item.Content
	post.Title = item.Title

	err := caus.UpdateDedupId(post)
	if err != nil {
		return utils.ImmediatePrintError(err)
	}
	return nil
}

// Caus news is one subsource
func (caus CaUsNewsCrawler) CollectOneSubsourceOnePage(
	task *protocol.PanopticTask,
	paginationInfo *working_context.PaginationInfo,
) error {
	lanmuId := 3
	if task.TaskParams.GetCausNewsTaskParams() != nil {
		lanmuId = int(task.TaskParams.GetCausNewsTaskParams().LanmuId)
	}
	bodyStr := fmt.Sprintf(`{"lanmuId": %d, "filterIds": []}`, lanmuId)
	client := collector.NewHttpClientFromTaskParams(task)
	url := caus.ConstructUrl()
	resp, err := client.Post(url, strings.NewReader(bodyStr))
	if err != nil {
		utils.ImmediatePrintError(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	res := &CaUsNewsResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		utils.ImmediatePrintError(err)
	}

	if res.Success != true {
		utils.ImmediatePrintError(errors.New(fmt.Sprintf("response not success: %+v", res)))
	}
	for _, item := range res.Data.ArticleList {
		// working context for each message
		workingContext := &working_context.ApiCollectorWorkingContext{
			SharedContext: working_context.SharedContext{
				Task:                 task,
				Result:               &protocol.CrawlerMessage{},
				IntentionallySkipped: false,
			},
			SubSource:       task.TaskParams.SubSources[0],
			ApiUrl:          url,
			NewsType:        protocol.PanopticSubSource_UNSPECIFIED,
			ApiResponseItem: item,
		}
		collector.InitializeApiCollectorResult(workingContext)
		err := caus.UpdateResult(workingContext)
		if err != nil {
			task.TaskMetadata.TotalMessageFailed++
			return utils.ImmediatePrintError(err)
		}

		if workingContext.SharedContext.Result != nil {
			sink.PushResultToSinkAndRecordInTaskMetadata(caus.Sink, workingContext)
		}

		cursor := fmt.Sprint(item.PublishTime)
		if strings.Compare(paginationInfo.NextPageId, cursor) > 0 {
			paginationInfo.NextPageId = cursor
		}
	}

	return nil
}

// Support configable multi-page API call
func (caus CaUsNewsCrawler) CollectOneSubsource(task *protocol.PanopticTask, subsource *protocol.PanopticSubSource) error {
	// NextPageId is set from API
	paginationInfo := working_context.PaginationInfo{
		CurrentPageCount: 1,
		NextPageId:       "",
	}

	maxPages := 1
	if task.TaskParams.GetCausNewsTaskParams() != nil {
		maxPages = int(task.TaskParams.GetCausNewsTaskParams().MaxPages)
	}

	for {
		err := caus.CollectOneSubsourceOnePage(task, &paginationInfo)
		if err != nil {
			return err
		}
		paginationInfo.CurrentPageCount++
		if task.GetTaskParams() == nil || paginationInfo.CurrentPageCount > maxPages {
			break
		}
	}

	collector.SetErrorBasedOnCounts(task, "caus_news")
	return nil
}

func (caus CaUsNewsCrawler) CollectAndPublish(task *protocol.PanopticTask) {
	if len(task.TaskParams.SubSources) != 1 {
		utils.ImmediatePrintError(errors.New("subsource length is not 1"))
		return
	}

	caus.CollectOneSubsource(task, task.TaskParams.SubSources[0])
}