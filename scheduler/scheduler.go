package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	gh "github.com/philip-gai/gh-schedule/gh"
)

var scheduler *gocron.Scheduler
var jobs = []string{}

type ScheduleJobOptions struct {
	// At   string
	// On 	string
	In       string
	GhCliCmd []string
}

func ScheduleJob(opts ScheduleJobOptions) {
	jobName := fmt.Sprintf("%d: %v in %s", len(jobs), opts.GhCliCmd, opts.In)
	fmt.Println("\n" + jobName)
	jobs = append(jobs, jobName)
	duration, err := time.ParseDuration(opts.In)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	time.Sleep(duration)
	scheduler.Every(opts.In).LimitRunsTo(1).Do(func() {
		fmt.Println("Running job:", jobName)
		gh.Exec(opts.GhCliCmd...)
	})
	scheduler.StartAsync()
}

func Start() {
	if scheduler == nil {
		// fmt.Println("Starting scheduler")
		scheduler = gocron.NewScheduler(time.Local)
		scheduler.StartAsync()
	}
}
