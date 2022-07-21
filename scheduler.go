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
	jobName := fmt.Sprintf("Scheduling job %d to run %v in %s", len(jobs), opts.GhCliCmd, opts.In)
	fmt.Println(jobName)
	jobs = append(jobs, jobName)
	job, _ := scheduler.Every(opts.In).Do(func() {
		fmt.Println("Running job:", jobName)
		gh(opts.GhCliCmd...)
	})
	job.LimitRunsTo(1)
}

func startScheduler() {
	if scheduler == nil {
		// fmt.Println("Starting scheduler")
		scheduler = gocron.NewScheduler(time.Local)
		scheduler.StartAsync()
		scheduler.Every("1m").Do(func() {
			fmt.Println("\nScheduler is running")
			fmt.Print("Enter command: ")
		})
	} else {
		// fmt.Println("Scheduler already started")
	}
}
