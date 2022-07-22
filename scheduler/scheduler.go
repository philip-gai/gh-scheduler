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
var jobs = []string{}

type ScheduleJobOptions struct {
	In       string
	GhCliCmd []string
}

func ScheduleJob(opts ScheduleJobOptions, logs *widgets.List) {
	humanReadableArgs := fmt.Sprintf("gh %s", strings.Join(opts.GhCliCmd, " "))
	jobName := fmt.Sprintf("%d: %s in %s", len(jobs), humanReadableArgs, opts.In)
	jobs = append(jobs, jobName)
	duration, err := time.ParseDuration(opts.In)
	if err != nil {
		utils.PushListRow(fmt.Sprint("Error:", err), logs)
		return
	}
	utils.PushListRow(fmt.Sprintf("Scheduled to run \"%s\" in %s", humanReadableArgs, opts.In), logs)
	time.Sleep(duration)
	scheduler.Every(opts.In).LimitRunsTo(1).Do(func() {
		gh.Exec(logs, opts.GhCliCmd...)
	})
	scheduler.StartAsync()
}

func Start() {
	if scheduler == nil {
		scheduler = gocron.NewScheduler(time.Local)
		scheduler.StartAsync()
	}
}
