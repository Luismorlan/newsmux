package modules

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Luismorlan/newsmux/protocol"
	Logger "github.com/Luismorlan/newsmux/utils/log"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/encoding/prototext"
)

// A valid job batch must not contains duplicate job name.
func ValidateJobs(jobs []*SchedulerJob) error {
	seen := make(map[string]bool)
	for _, job := range jobs {
		if _, ok := seen[job.panopticConfig.Name]; ok {
			return fmt.Errorf("duplicate scheduler job name: %s", job.panopticConfig.Name)
		}
		seen[job.panopticConfig.Name] = true
	}
	return nil
}

type SchedulerConfig struct {
	// Name of the scheduler.
	Name string

	// Path of the schedules configs. In dev this is a local filepath, while in
	// prod this points to a S3 path.
	PanopticConfigPath string
}

type Scheduler struct {
	m sync.RWMutex

	// Config for this scheduler.
	Config SchedulerConfig

	// Context of this Scheduler.
	ctx context.Context

	// A list of SchedulerJobs that this scheduler is managing.
	Jobs []*SchedulerJob

	// Whether this scheduler is running.
	running bool

	Doer JobDoer

	EventBus *gochannel.GoChannel
}

// Return a new instance of Scheduler.
func NewScheduler(
	config SchedulerConfig, e *gochannel.GoChannel, doer JobDoer, ctx context.Context) *Scheduler {
	scheduler := &Scheduler{
		Config:   config,
		ctx:      ctx,
		EventBus: e,
		Doer:     doer,
		running:  false,
	}

	err := scheduler.ParseAndUpsertJobs()
	if err != nil {
		log.Fatalf("cannot initialize scheduler: %v", err)
	}

	return scheduler
}

// For existing jobs, only job's PanopticConfig is updated. Otherwise remove
// from the job list. If the job is already in pending state, cancel it
// proactively. For all new jobs, append to the end of job lists.
func (s *Scheduler) UpsertJobs(jobs []*SchedulerJob) {
	s.m.Lock()
	defer s.m.Unlock()

	nameToJobMap := make(map[string]*SchedulerJob)

	// Index all jobs by it's config name.
	for idx := range jobs {
		nameToJobMap[jobs[idx].panopticConfig.Name] = jobs[idx]
	}

	// Existing Jobs.
	idx := 0
	for idx < len(s.Jobs) {
		existingJob := s.Jobs[idx]
		if v, ok := nameToJobMap[existingJob.panopticConfig.Name]; ok {
			delete(nameToJobMap, existingJob.panopticConfig.Name)
			existingJob.panopticConfig = v.panopticConfig
			idx += 1
		} else {
			s.Jobs = append(s.Jobs[:idx], s.Jobs[idx+1:]...)
			existingJob.cancel()
		}
	}

	// New Jobs.
	for _, v := range nameToJobMap {
		s.Jobs = append(s.Jobs, v)
	}
}

// Read config either from local workspace (dev) or from Github (production)
func (s *Scheduler) ReadConfig() (*protocol.PanopticConfigs, error) {
	env := os.Getenv("NEWSMUX_ENV")
	configs := &protocol.PanopticConfigs{}

	if env == "prod" {
		Logger.Log.Infoln("read config from Github project: https://github.com/Luismorlan/panoptic_config")
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
		)
		tc := oauth2.NewClient(s.ctx, ts)
		client := github.NewClient(tc)
		content, _, res, err := client.Repositories.GetContents(s.ctx, "Luismorlan", "panoptic_config", "config.textproto", nil)
		if err != nil {
			return nil, err
		}
		if res.StatusCode != 200 {
			return nil, fmt.Errorf("fail to get config from Github, http code %d", res.StatusCode)
		}
		decode, _ := base64.StdEncoding.DecodeString(*content.Content)
		if err := prototext.Unmarshal(decode, configs); err != nil {
			return nil, err
		}
	} else {
		Logger.Log.Infoln("read config from local workspace, file panoptic/data/testing_panoptic_config.textproto")
		in, err := ioutil.ReadFile(s.Config.PanopticConfigPath)
		if err != nil {
			return nil, err
		}
		if err := prototext.Unmarshal(in, configs); err != nil {
			return nil, err
		}
	}

	return configs, nil
}

func (s *Scheduler) ParseAndUpsertJobs() error {
	configs, err := s.ReadConfig()

	Logger.Log.Infof("initial PanopticConfigs: %s", configs.String())

	if err != nil {
		return err
	}
	jobs := NewSchedulerJobs(configs, s.ctx)
	err = ValidateJobs(jobs)
	if err != nil {
		return err
	}

	s.UpsertJobs(jobs)
	return nil
}

func (s *Scheduler) DoSingleJob(job *SchedulerJob) {
	err := s.Doer.Do(job)
	if err != nil {
		log.Printf(
			"Job execution failed. Name: %s, err: %v",
			job.panopticConfig.Name,
			err,
		)
	}
}

func (s *Scheduler) ScheduleSingleJob(job *SchedulerJob) {
	// Start immediately if required and never ran before.
	if !job.HasRunBefore() && job.panopticConfig.TaskSchedule.StartImmediatly {
		job.UpdateLastAndNextTime()
		// Execute the job in a non-blocking way so that we the execution time will
		// not skew the next run time.
		go s.DoSingleJob(job)
	}

	for {
		durationTillNextRun := job.DurationTillNextRun()
		select {
		// Scheduler's lifecycle is managed by engine's context, cancelling engine
		// should also shutdown scheduler.
		case <-s.ctx.Done():
			log.Printf("Job %s cancelled by scheduler.", job.panopticConfig.Name)
			return
		// In the future, a job could cancel itself under some condition (e.g. keep
		// failing, reach max run count)
		case <-job.ctx.Done():
			log.Printf("Job %s cancelled by itself.", job.panopticConfig.Name)
			return
		case <-time.After(durationTillNextRun):
			job.UpdateLastAndNextTime()
			go s.DoSingleJob(job)
		}
	}
}

// A blocking call that returns once all jobs finished running.
func (s *Scheduler) ScheduleJobs() {
	log.Println("start scheduler jobs")
	var wg sync.WaitGroup

	for _, j := range s.Jobs {
		wg.Add(1)
		go func(job *SchedulerJob) {
			defer wg.Done()
			s.ScheduleSingleJob(job)
		}(j)
	}

	wg.Wait()
	log.Println("finished scheduler")
}

func (s *Scheduler) RunModule(ctx context.Context) error {
	s.ScheduleJobs()
	return nil
}

func (s *Scheduler) Name() string {
	return s.Config.Name
}
