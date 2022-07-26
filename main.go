package main

import (
	scheduler "github.com/philip-gai/gh-scheduler/scheduler"
	ui "github.com/philip-gai/gh-scheduler/ui"
)

func main() {
	scheduler.Start()
	ui.Render()
}
