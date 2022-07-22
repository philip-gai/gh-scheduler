package scheduler

import (
	"fmt"
	"strings"
	"time"

	"github.com/gizak/termui/v3/widgets"
	"github.com/go-co-op/gocron"
	"github.com/philip-gai/gh-scheduler/gh"
	"github.com/philip-gai/gh-scheduler/utils"
)

var scheduler *gocron.Scheduler
var jobs = []utils.JobInfo{}

type ScheduleJobOptions struct {
	In       string
	GhCliCmd []string
}

func ScheduleJob(opts ScheduleJobOptions, logs *widgets.List, jobTable *widgets.Table) {
	humanReadableArgs := fmt.Sprintf("gh %s", strings.Join(opts.GhCliCmd, " "))

	duration, err := time.ParseDuration(opts.In)
	if err != nil {
		utils.PushListRow(fmt.Sprint("Error:", err), logs)
		return
	}

	scheduledAtTime := time.Now().Add(duration)
	scheduledAtFormatted := scheduledAtTime.Format(time.UnixDate)

	jobInfo := utils.JobInfo{
		ID:           len(jobs) + 1,
		Action:       humanReadableArgs,
		Status:       "Pending",
		ScheduledFor: scheduledAtFormatted,
		CreatedAt:    time.Now().Format(time.Stamp),
	}
	jobs = append(jobs, jobInfo)
	utils.PushJobRow(jobInfo, jobTable)
	utils.PushListRow(fmt.Sprintf("Scheduled to run \"%s\" in %s", humanReadableArgs, opts.In), logs)

	go func() {
		time.Sleep(duration)
		scheduler.Every(opts.In).LimitRunsTo(1).Do(func() {
			gh.Exec(logs, opts.GhCliCmd...)
			jobInfo.Status = "Completed"
			utils.UpdateJobRow(jobInfo, jobTable)
		})
	}()
}

func Start() {
	if scheduler == nil {
		scheduler = gocron.NewScheduler(time.Local)
		scheduler.StartAsync()
	}
}
