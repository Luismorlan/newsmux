package main

import (
	"context"
	"log"

	"github.com/Luismorlan/newsmux/panoptic"
	"github.com/Luismorlan/newsmux/panoptic/modules"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main() {
	eventbus := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	// Initialize all engine modules here.
	modules := []panoptic.Module{
		// Scheduler parses data collector configs, fanout into multiple tasks and
		// pushes onto EventBus.
		modules.NewScheduler(
			modules.SchedulerConfig{Name: "scheduler"},
			eventbus,
		),
		// Orchestrator listens tasks on EventBus, maintains an active Lambda pool
		// and wrap Lambda result in a tasks and publish to the exporter for
		// monitoring.
		modules.NewOrchestrator(
			modules.OrchestratorConfig{Name: "orchestrator"},
			eventbus,
		),
	}

	engine := panoptic.NewEngine(modules, eventbus)
	ctx := context.Background()

	// blocking call.
	engine.Run(ctx)

	log.Println("engine stopped execution.")
}
