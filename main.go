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
}

func mergeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "merge <pull_url>",
		Short: "Merge a pull request at the scheduled time",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := mergeOptions{}
			if len(args) == 0 {
				return fmt.Errorf("pull_url is required")
			} else {
				opts.PullUrl = args[0]
			}
			return runMerge(opts)
		},
	}
}

func runMerge(opts mergeOptions) error {
	fmt.Printf("Merging %s\n", opts.PullUrl)
	return nil
}

func main() {
	rc := rootCmd()
	rc.AddCommand(mergeCmd())
	if err := rc.Execute(); err != nil {
		// TODO not bothering as long as cobra is also printing error
		//fmt.Println(err)
		os.Exit(1)
	}
}
