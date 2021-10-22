package collector_instances

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/Luismorlan/newsmux/collector"
	sink "github.com/Luismorlan/newsmux/collector/sink"
	"github.com/Luismorlan/newsmux/collector/working_context"
	"github.com/Luismorlan/newsmux/protocol"
	"github.com/Luismorlan/newsmux/utils"
	Logger "github.com/Luismorlan/newsmux/utils/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const ()

type WallstreetApiCollector struct {
	Sink sink.CollectedDataSink
}

type WallstreetItem struct {
	Content     string `json:"content"`
	DisplayTime int    `json:"display_time"`
	ID          int    `json:"id"`
	Score       int    `json:"score"`
}

type WallstreetApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []WallstreetItem `json:"items"`
	} `json:"data"`
}

func (w WallstreetApiCollector) UpdateFileUrls(workingContext *working_context.ApiCollectorWorkingContext) error {
	return errors.New("UpdateFileUrls not implemented, should not be called")
}

func (w WallstreetApiCollector) ConstructUrl(task *protocol.PanopticTask, subsource *protocol.PanopticSubSource, paginationInfo *working_context.PaginationInfo) string {
	return fmt.Sprintf("https://api.wallstcn.com/apiv1/content/lives?channel=%s&client=pc&limit=%d",
		paginationInfo.NextPageId,
		task.TaskParams.GetWallstreetNewsTaskParams().Limit,
	)
}

func (w WallstreetApiCollector) UpdateDedupId(post *protocol.CrawlerMessage_CrawledPost) error {
	md5, err := utils.TextToMd5Hash(post.SubSource.SourceId + post.SubSource.ExternalId)
	if err != nil {
		return utils.ImmediatePrintError(err)
	}
	post.DeduplicateId = md5
	return nil
}

func (w WallstreetApiCollector) UpdateResultFromItem(item *WallstreetItem, workingContext *working_context.ApiCollectorWorkingContext) error {
	generatedTime := time.Unix(int64(item.DisplayTime), 0)
	workingContext.Result.Post.ContentGeneratedAt = timestamppb.New(generatedTime)
	workingContext.Result.Post.SubSource.ExternalId = fmt.Sprint(item.ID)
	if err := w.UpdateDedupId(workingContext.Result.Post); err != nil {
		return utils.ImmediatePrintError(err)
	}
	workingContext.Result.Post.Content = item.Content
	newsType := protocol.PanopticSubSource_FLASHNEWS
	if item.Score != 1 {
		newsType = protocol.PanopticSubSource_KEYNEWS
	}
	workingContext.NewsType = newsType
	workingContext.Result.Post.SubSource.Name = collector.SubsourceTypeToName(newsType)
	return nil
}

func (w WallstreetApiCollector) CollectOneSubsourceOnePage(
	task *protocol.PanopticTask,
	subsource *protocol.PanopticSubSource,
	paginationInfo *working_context.PaginationInfo,
) error {
	var client collector.HttpClient
	url := w.ConstructUrl(task, subsource, paginationInfo)
	resp, err := client.Get(url)
	if err != nil {
		return utils.ImmediatePrintError(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := &WallstreetApiResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return utils.ImmediatePrintError(err)
	}
	if res.Message != "OK" {
		return fmt.Errorf("response not success: %v", res)
	}

	for _, item := range res.Data.Items {
		// working context for each message
		workingContext := &working_context.ApiCollectorWorkingContext{
			SharedContext:  working_context.SharedContext{Task: task, Result: &protocol.CrawlerMessage{}},
			PaginationInfo: paginationInfo,
			ApiUrl:         url,
			SubSource:      subsource,
		}
		collector.InitializeApiCollectorResult(workingContext)
		err := w.UpdateResultFromItem(&item, workingContext)
		if err != nil {
			task.TaskMetadata.TotalMessageFailed++
			return utils.ImmediatePrintError(err)
		} else {
			if !collector.IsRequestedNewsType(workingContext.Task.TaskParams.SubSources, workingContext.NewsType) {
				workingContext.Result = nil
				return nil
			}
			if err = w.Sink.Push(workingContext.Result); err != nil {
				task.TaskMetadata.ResultState = protocol.TaskMetadata_STATE_FAILURE
				task.TaskMetadata.TotalMessageFailed++
				return utils.ImmediatePrintError(err)
			}
		}
		task.TaskMetadata.TotalMessageCollected++
		Logger.Log.Debug(workingContext.Result.Post.Content)
	}

	collector.SetErrorBasedOnCounts(task, url, fmt.Sprintf("subsource: %s, body: %s", subsource.Name, string(body)))
	return nil
}

// Support configable multi-page API call
func (w WallstreetApiCollector) CollectOneSubsource(task *protocol.PanopticTask, subsource *protocol.PanopticSubSource) error {
	// Wallstreet uses channels and only know subsource after each message if fetched
	for ind, channel := range task.TaskParams.GetWallstreetNewsTaskParams().Channels {
		w.CollectOneSubsourceOnePage(task, subsource, &working_context.PaginationInfo{
			CurrentPageCount: ind,
			NextPageId:       channel,
		})
	}
	return nil
}

func (w WallstreetApiCollector) CollectAndPublish(task *protocol.PanopticTask) {
	task.TaskMetadata.ResultState = protocol.TaskMetadata_STATE_SUCCESS
	w.CollectOneSubsource(task, &protocol.PanopticSubSource{})
}
