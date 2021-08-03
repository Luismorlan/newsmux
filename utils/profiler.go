package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func init() {
	// Datadog profiler
	if err := profiler.Start(
		profiler.WithService("apiserver"),
		profiler.WithEnv("development"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			// The profiles below are disabled by
			// default to keep overhead low, but
			// can be enabled as needed.
			// profiler.BlockProfile,
			// profiler.MutexProfile,
			// profiler.GoroutineProfile,
		),
	); err != nil {
		Logger.Fatal(err)
	}

	Logger.WithFields(logrus.Fields{"service": "api_server", "is_development": IsDevelopment}).Info("profiler initialized")
}

// Stop profiler, OK to be closed multiple times
func CloseProfiler() {
	// Datadog profiler
	profiler.Stop()
}
