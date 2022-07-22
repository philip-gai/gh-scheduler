package tui

import (
	"fmt"
	"log"
	"strings"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/philip-gai/gh-scheduler/scheduler"
	"github.com/philip-gai/gh-scheduler/utils"
)

var grid *termui.Grid
var actions *widgets.Table
var console *widgets.List
var logs *widgets.List
var jobTable *widgets.Table
var userInput string

type mergeOptions struct {
	PullUrl string
	In      string
}

func Render() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	actions = createActionsSection()
	console = createConsole()
	logs = createLogsSection()
	jobTable = createJobTable()
	grid = initializeGrid()
	startEventPolling()
}

func runCommand(userInput string) {
	// Re-enable once we have debug logging
	// utils.PushListRow(fmt.Sprintf("Running command \"%s\"", userInput), logs)

	// Execute action
	userInput = strings.TrimRight(userInput, "\n")

	if len(userInput) > 0 {
		args := strings.Split(userInput, " ")

		if len(args) == 0 {
			utils.PushListRow("Error: no commands given", logs)
		} else {
			command := args[0]
			if command == "merge" {
				if len(args) == 4 {
					opts := mergeOptions{}
					opts.PullUrl = args[1]
					opts.In = args[3]
					runMerge(opts)
				} else {
					utils.PushListRow("Error: not enough arguments", logs)
				}
			} else {
				utils.PushListRow(fmt.Sprintf("Error: unknown command \"%s\"", command), logs)
			}
		}
	}
	utils.PushListRow("$ ", console)
}

func startEventPolling() {
	uiEvents := termui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {

		// Exit on Escape or ctrl-c.
		case "<Escape>", "<C-c>":
			return

		// Redraw grid on window resize
		case "<Resize>":
			payload := e.Payload.(termui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			termui.Clear()

		case "<Enter>":
			runCommand(userInput)
			userInput = ""

		default:
			// This is the user regularly typing in the console
			if e.Type == termui.KeyboardEvent {
				if len(e.ID) == 1 {
					utils.ConcatListRow(e.ID, console)
					userInput += e.ID
				} else if e.ID == "<Backspace>" {
					if userInput != "" {
						utils.BackspaceListRow(console)
						userInput = userInput[:len(userInput)-1]
					}
				} else if e.ID == "<Space>" {
					utils.ConcatListRow(" ", console)
					userInput += " "
				}
			}
		}
		// Rerender the grid on any event to make sure it's up to date
		termui.Render(grid)
	}
}

func initializeGrid() *termui.Grid {
	grid := termui.NewGrid()
	termWidth, termHeight := termui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		termui.NewRow(2.0/4, termui.NewCol(1.0, jobTable)),
		termui.NewRow(1.0/4, termui.NewCol(1.0, actions)),
		termui.NewRow(1.0/4, termui.NewCol(1.0/2, console), termui.NewCol(1.0/2, logs)),
	)
	termui.Render(grid)
	return grid
}

func createActionsSection() *widgets.Table {
	actions := widgets.NewTable()
	actions.Title = "Actions"
	actions.Rows = [][]string{
		{"Action", "Example"},
		{"Merge a pull request", "merge https://github.com/philip-gai/gh-scheduler/pull/1 in 1h30m"},
	}
	actions.TextAlignment = termui.AlignCenter
	return actions
}

func createJobTable() *widgets.Table {
	// Job table
	jobTable := widgets.NewTable()
	jobTable.Title = "Jobs"
	jobTable.Rows = [][]string{
		{"Job", "Created", "Scheduled", "Status"},
	}
	jobTable.TextStyle = termui.NewStyle(termui.ColorWhite)
	jobTable.TextAlignment = termui.AlignCenter
	return jobTable
}

func createLogsSection() *widgets.List {
	logs := widgets.NewList()
	logs.Title = "Logs"
	return logs
}

func createConsole() *widgets.List {
	// Information and user input
	console := widgets.NewList()
	console.Title = "Console"
	console.Rows = []string{
		"$ ",
	}
	return console
}

func runMerge(opts mergeOptions) error {
	utils.PushListRow(fmt.Sprintf("Scheduling merge of %s in %s\n", opts.PullUrl, opts.In), logs)
	ghCliCmd := []string{"pr", "merge", opts.PullUrl}
	scheduler.ScheduleJob(scheduler.ScheduleJobOptions{
		In:       opts.In,
		GhCliCmd: ghCliCmd,
	}, logs)
	return nil
}
