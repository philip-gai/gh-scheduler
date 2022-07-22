package scheduler

import (
	"fmt"
	"time"

	"github.com/gizak/termui/v3/widgets"
	"github.com/go-co-op/gocron"
	"github.com/philip-gai/gh-scheduler/gh"
	"github.com/philip-gai/gh-scheduler/utils"
)

var scheduler *gocron.Scheduler
var jobs = []string{}

type ScheduleJobOptions struct {
	In       string
	GhCliCmd []string
}

func ScheduleJob(opts ScheduleJobOptions, logs *widgets.List) {
	jobName := fmt.Sprintf("%d: %v in %s", len(jobs), opts.GhCliCmd, opts.In)
	utils.PushListRow(jobName, logs)
	jobs = append(jobs, jobName)
	duration, err := time.ParseDuration(opts.In)
	if err != nil {
		utils.PushListRow(fmt.Sprintln("Error:", err), logs)
		return
	}
	time.Sleep(duration)
	scheduler.Every(opts.In).LimitRunsTo(1).Do(func() {
		utils.PushListRow(fmt.Sprintln("Running job:", jobName), logs)
		gh.Exec(opts.GhCliCmd...)
	})
	scheduler.StartAsync()
}

func Start() {
	if scheduler == nil {
		scheduler = gocron.NewScheduler(time.Local)
		scheduler.StartAsync()
	}
}
