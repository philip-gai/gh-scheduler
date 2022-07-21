package main

import (
	"fmt"
	"os"

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
	ghCliCmd := []string{"pr", "merge", opts.PullUrl}
	scheduleJob(scheduleJobOptions{
		In:       opts.In,
		GhCliCmd: ghCliCmd,
	})
	return nil
}

func main() {
	go startScheduler()
	rc := rootCmd()
	rc.AddCommand(mergeCmd())
	if err := rc.Execute(); err != nil {
		// TODO not bothering as long as cobra is also printing error
		//fmt.Println(err)
		os.Exit(1)
	}
}
