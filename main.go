package main

import (
	scheduler "github.com/philip-gai/gh-scheduler/scheduler"
	dashboard "github.com/philip-gai/gh-scheduler/ui/views"
)

func main() {
	scheduler.Start()
	dashboard.Init()
}
