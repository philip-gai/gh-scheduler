package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use: "schedule",
	}
}

type mergeOptions struct {
	PullUrl string
	In      string
}

func mergeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "merge <pull_url> in <time_string>",
		Short:   "Merge a pull request at a future time",
		Example: "gh schedule merge https://github.com/philip-gai/gh-schedule/pull/1 in 10m",
		Args:    cobra.MaximumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := mergeOptions{}
			if len(args) == 0 {
				return fmt.Errorf("invalid arguments")
			} else if len(args) == 3 {
				opts.PullUrl = args[0]
				opts.In = args[2]
			}
			return runMerge(opts)
		},
	}
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

	// TODO: Figure out how to pass stdin to cobra
	// rc := rootCmd()
	// rc.AddCommand(mergeCmd())
	// if err := rc.Execute(); err != nil {
	// 	// TODO not bothering as long as cobra is also printing error
	// 	//fmt.Println(err)
	// 	os.Exit(1)
	// }

	fmt.Println("Welcome to gh-schedule!")
	fmt.Println("Available Commands:\n * merge <pull_url> in <time_string>\n * <Enter>: Exits the scheduler")

	for {
		// TODO - Press enter to exit, otherwise enter more schedule commands
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			fmt.Println("Stopping the scheduler")
			break
		}
		fmt.Println("Running command:", text)

		args := strings.Split(text, " ")

		if len(args) == 0 {
			fmt.Println("Error: invalid arguments")
		} else {
			command := args[0]
			if command == "merge" {
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
