package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type mergeOptions struct {
	PullUrl string
	In      string
}

func runMerge(opts mergeOptions) error {
	fmt.Printf("Scheduling merge of %s in %s", opts.PullUrl, opts.In)
	ghCliCmd := []string{"pr", "merge", opts.PullUrl}
	scheduleJob(scheduleJobOptions{
		In:       opts.In,
		GhCliCmd: ghCliCmd,
	})
	return nil
}

func main() {
	startScheduler()

	fmt.Println("Welcome to gh-schedule!")
	fmt.Println("Available Commands:\n * merge <pull_url> in <time_string>\n * Ctrl-C: Exits the scheduler")

	for {
		// TODO - Press enter to exit, otherwise enter more schedule commands
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")
		fmt.Println("Command:", text)

		args := strings.Split(text, " ")

		if len(args) == 0 {
			fmt.Println("Error: invalid arguments")
		} else {
			command := args[0]
			if command == "merge" {
				// Example: merge https://github.com/philip-gai/gh-schedule/pull/1 in 5s
				opts := mergeOptions{}
				opts.PullUrl = args[1]
				opts.In = args[3]
				runMerge(opts)
			} else {
				fmt.Println("Unknown command:", command)
			}
		}
	}
}
