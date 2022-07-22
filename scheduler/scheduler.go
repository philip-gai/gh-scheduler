package scheduler

import (
	"fmt"
	"time"

	"github.com/gizak/termui/v3/widgets"
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

func ScheduleJob(opts ScheduleJobOptions, logs *widgets.Paragraph) {
	jobName := fmt.Sprintf("%d: %v in %s", len(jobs), opts.GhCliCmd, opts.In)
	logs.Text += jobName
	jobs = append(jobs, jobName)
	duration, err := time.ParseDuration(opts.In)
	if err != nil {
		logs.Text += fmt.Sprintln("Error:", err)
		return
	}
	time.Sleep(duration)
	scheduler.Every(opts.In).LimitRunsTo(1).Do(func() {
		logs.Text += fmt.Sprintln("Running job:", jobName)
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
