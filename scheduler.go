package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func scheduleJob() {
	fmt.Println("Creating scheduler")
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every("5s").Do(func() {
		fmt.Println("Testing...")
	})
	scheduler.StartBlocking()
}
