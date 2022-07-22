package main

import (
	scheduler "github.com/philip-gai/gh-scheduler/scheduler"
	tui "github.com/philip-gai/gh-scheduler/tui"
)

func main() {
	scheduler.Start()
	tui.Render()
}
