package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

var scheduler *gocron.Scheduler

// Create messages string array
var jobs = []string{}

type scheduleJobOptions struct {
	// At   string
	// On 	string
	In       string
	GhCliCmd []string
}

func scheduleJob(opts scheduleJobOptions) {
	jobName := fmt.Sprintf("Scheduled job %d to run %v in %s", len(jobs), opts.GhCliCmd, opts.In)
	fmt.Println(jobName)
	jobs = append(jobs, jobName)
}

func startScheduler() {
	if scheduler == nil {
		fmt.Println("Starting scheduler")
		scheduler = gocron.NewScheduler(time.Local)
		scheduler.StartBlocking()
	} else {
		fmt.Println("Scheduler already started")
	}
}
