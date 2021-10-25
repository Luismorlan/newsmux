package collector_job_handler

import (
	"errors"
	"fmt"
	"sync"

	. "github.com/Luismorlan/newsmux/collector"
	. "github.com/Luismorlan/newsmux/collector/builder"
	"github.com/Luismorlan/newsmux/collector/file_store"
	. "github.com/Luismorlan/newsmux/collector/instances"
	"github.com/Luismorlan/newsmux/collector/sink"
	"github.com/Luismorlan/newsmux/protocol"
	"github.com/Luismorlan/newsmux/utils"
	Logger "github.com/Luismorlan/newsmux/utils/log"
	"github.com/golang/protobuf/proto"
)

type DataCollectJobHandler struct{}

func UpdateIpAddressesInTasks(ip string, job *protocol.PanopticJob) {
	for _, task := range job.Tasks {
		if task.TaskMetadata == nil {
			task.TaskMetadata = &protocol.TaskMetadata{}
		}
		task.TaskMetadata.IpAddr = ip
	}
}

// This is the entry point to data collector.
func (handler DataCollectJobHandler) Collect(job *protocol.PanopticJob) (err error) {
	Logger.Log.Info("Collect() with request: \n", proto.MarshalTextString(job))
	var (
		s          sink.CollectedDataSink
		imageStore file_store.CollectedFileStore
		wg         sync.WaitGroup
		httpClient HttpClient
	)
	ip, err := GetCurrentIpAddress(httpClient)
	if err == nil {
		if job == nil {
			fmt.Println("============= nil job ============")
		}
		UpdateIpAddressesInTasks(ip, job)
	} else {
		Logger.Log.Error("ip fetching error: ", err)
	}

	if !utils.IsProdEnv() || job.Debug {
		s = sink.NewStdErrSink()
		// Debug job + Prod env still download files.
		// TODO(chenweilunster): Clean up to a simpler logic. Maybe using a Noop
		// file store.
		if utils.IsProdEnv() {
			if imageStore, err = file_store.NewS3FileStore(file_store.ProdS3ImageBucket); err != nil {
				return err
			}
		} else {
			if imageStore, err = file_store.NewLocalFileStore("test"); err != nil {
				return err
			}
		}

		defer imageStore.CleanUp()
	} else {
		s, err = sink.NewSnsSink()
		if err != nil {
			return err
		}
		if imageStore, err = file_store.NewS3FileStore(file_store.ProdS3ImageBucket); err != nil {
			return err
		}
	}

	for ind := range job.Tasks {
		t := job.Tasks[ind]
		wg.Add(1)
		go func(t *protocol.PanopticTask) {
			defer wg.Done()
			if err := handler.processTask(t, s, imageStore); err != nil {
				Logger.Log.Errorf("fail to process task: %s", err)
			}
			t.TaskMetadata.IpAddr = ip
		}(t)
	}
	wg.Wait()
	Logger.Log.Info("Collect() response: \n", proto.MarshalTextString(job))
	return nil
}

func (hanlder DataCollectJobHandler) processTask(t *protocol.PanopticTask, sink sink.CollectedDataSink, imageStore file_store.CollectedFileStore) error {
	var (
		collector DataCollector
		builder   CollectorBuilder
	)
	zsxqFileStore, err := GetZsxqS3FileStore(t, utils.IsProdEnv())
	if err != nil {
		return err
	}
	// forward task to corresponding collector
	switch t.DataCollectorId {
	case protocol.PanopticTask_COLLECTOR_JINSHI:
		// please follow this patter to get collector
		collector = builder.NewJin10Crawler(sink)
	case protocol.PanopticTask_COLLECTOR_WEIBO:
		collector = builder.NewWeiboApiCollector(sink, imageStore)
	case protocol.PanopticTask_COLLECTOR_ZSXQ:
		collector = builder.NewZsxqApiCollector(sink, imageStore, zsxqFileStore)
	case protocol.PanopticTask_COLLECTOR_WALLSTREET_NEWS:
		collector = builder.NewWallstreetNewsApiCollector(sink)
	case protocol.PanopticTask_COLLECTOR_KUAILANSI:
		collector = builder.NewKuailansiApiCollector(sink)
	case protocol.PanopticTask_COLLECTOR_JINSE:
		collector = builder.NewJinseApiCollector(sink)
	default:
		return errors.New("unknown task data collector id")
	}
	RunCollectorForTask(collector, t)
	return nil
}
