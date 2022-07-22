package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

var scheduler *gocron.Scheduler
var jobs = []string{}

type scheduleJobOptions struct {
	// At   string
	// On 	string
	In       string
	GhCliCmd []string
}

func scheduleJob(opts scheduleJobOptions) {
	jobName := fmt.Sprintf("%d: %v in %s", len(jobs), opts.GhCliCmd, opts.In)
	fmt.Println(jobName)
	jobs = append(jobs, jobName)
	scheduler.Every(opts.In).LimitRunsTo(1).Do(func() {
		fmt.Println("Running job:", jobName)
		// gh(opts.GhCliCmd...)
	})
	scheduler.StartAsync()
}

func startScheduler() {
	if scheduler == nil {
		// fmt.Println("Starting scheduler")
		scheduler = gocron.NewScheduler(time.Local)
		scheduler.StartAsync()
	}
}
