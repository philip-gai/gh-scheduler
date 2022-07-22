package tui

import (
	"fmt"
	"log"
	"strings"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/philip-gai/gh-scheduler/scheduler"
)

var grid *termui.Grid
var actions *widgets.Table
var console *widgets.List
var logs *widgets.Paragraph
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
	logs.Text += fmt.Sprintf("Running command \"%s\"\n", userInput)

	// Execute action
	userInput = strings.TrimRight(userInput, "\n")

	if len(userInput) > 0 {
		args := strings.Split(userInput, " ")

		if len(args) == 0 {
			logs.Text += fmt.Sprintln("Error: no commands given")
		} else {
			command := args[0]
			if command == "merge" {
				if len(args) == 3 {
					opts := mergeOptions{}
					opts.PullUrl = args[1]
					opts.In = args[3]
					runMerge(opts)
				} else {
					logs.Text += fmt.Sprintln("Error: not enough arguments")
				}
			} else {
				logs.Text += fmt.Sprintf("Error: unknown command \"%s\"\n", command)
			}
		}
	}
	pushConsoleRow("$ ")
	console.ScrollBottom()
}

func pushConsoleRow(text string) {
	console.Rows = append(console.Rows, text)
	termui.Render(console)
}

func appendToCurrentConsoleRow(text string) {
	console.Rows[len(console.Rows)-1] += text
	termui.Render(console)
}

func backspaceCurrentConsoleRow() {
	currentRow := console.Rows[len(console.Rows)-1]
	console.Rows[len(console.Rows)-1] = currentRow[:len(currentRow)-1]
	termui.Render(console)
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
					appendToCurrentConsoleRow(e.ID)
					userInput += e.ID
				} else if e.ID == "<Backspace>" {
					if userInput != "" {
						backspaceCurrentConsoleRow()
						userInput = userInput[:len(userInput)-1]
					}
				} else if e.ID == "<Space>" {
					appendToCurrentConsoleRow(" ")
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

func createLogsSection() *widgets.Paragraph {
	logs := widgets.NewParagraph()
	logs.Title = "Logs"
	return logs
}

func createConsole() *widgets.List {
	// Information and user input
	console := widgets.NewList()
	console.Title = "Console"
	// console.Text = "$ "
	console.Rows = []string{
		"$ ",
	}
	return console
}

func runMerge(opts mergeOptions) error {
	logs.Text += fmt.Sprintf("Scheduling merge of %s in %s\n", opts.PullUrl, opts.In)
	ghCliCmd := []string{"pr", "merge", opts.PullUrl}
	go scheduler.ScheduleJob(scheduler.ScheduleJobOptions{
		In:       opts.In,
		GhCliCmd: ghCliCmd,
	}, logs)
	return nil
}
