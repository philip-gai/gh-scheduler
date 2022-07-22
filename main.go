package main

import (
	scheduler "github.com/philip-gai/gh-scheduler/scheduler"
	tui "github.com/philip-gai/gh-scheduler/tui"
)

func main() {
	scheduler.Start()
	tui.Render()

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("\nEnter command: ")
	// text, _ := reader.ReadString('\n')
	// text = strings.TrimRight(text, "\n")
	// fmt.Println("Command:", text)

	// TODO - replace with TUI

	// fmt.Println("Welcome to gh-scheduler!")
	// fmt.Println("Available Commands:\n * merge <pull_url> in <time_string>\n * Ctrl-C: Exits the scheduler")

	// for {
	// 	// TODO - Press enter to exit, otherwise enter more schedule commands
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("\nEnter command: ")
	// 	text, _ := reader.ReadString('\n')
	// 	text = strings.TrimRight(text, "\n")
	// 	fmt.Println("Command:", text)

	// 	args := strings.Split(text, " ")

	// 	if len(args) == 0 {
	// 		fmt.Println("Error: invalid arguments")
	// 	} else {
	// 		command := args[0]
	// 		if command == "merge" {
	// 			// Example: merge https://github.com/philip-gai/gh-scheduler/pull/1 in 5s
	// 			opts := mergeOptions{}
	// 			opts.PullUrl = args[1]
	// 			opts.In = args[3]
	// 			runMerge(opts)
	// 		} else {
	// 			fmt.Println("Unknown command:", command)
	// 		}
	// 	}
	// }
}
