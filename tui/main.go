package tui

import (
	"fmt"
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/philip-gai/gh-schedule/scheduler"
)

var grid *ui.Grid
var actions *widgets.Table
var console *widgets.Paragraph
var logs *widgets.Paragraph
var userInput string

var CurrentState State

type State int64

const (
	SelectAction                State = 0
	MergePull_PullRequestPrompt State = 1
	MergePull_TimePrompt        State = 2
)

type mergeOptions struct {
	PullUrl string
	In      string
}

func Render() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	actions = createActionsSection()
	console = createConsole()
	logs = createLogsSection()
	initializeGrid()
}

// func handleEvents() {
// 	uiEvents := ui.PollEvents()
// 	for {
// 		e := <-uiEvents
// 		switch e.ID {

// 		// Exit on Escape or ctrl-c.
// 		case "<Escape>", "<C-c>":
// 			return

// 		// Redraw grid on window resize
// 		case "<Resize>":
// 			payload := e.Payload.(ui.Resize)
// 			grid.SetRect(0, 0, payload.Width, payload.Height)
// 			ui.Clear()
// 			ui.Render(grid)

// 		// Action list navigation
// 		case "<Down>":
// 			actions.ScrollDown()
// 		case "<Up>":
// 			actions.ScrollUp()
// 		case "<C-d>":
// 			actions.ScrollHalfPageDown()
// 		case "<C-u>":
// 			actions.ScrollHalfPageUp()
// 		case "<C-f>":
// 			actions.ScrollPageDown()
// 		case "<C-b>":
// 			actions.ScrollPageUp()
// 		case "<Home>":
// 			actions.ScrollTop()
// 		case "<End>":
// 			actions.ScrollBottom()
// 		case "<Enter>":
// 			logs.Text += fmt.Sprintf("You selected: %s\n", actions.Rows[actions.SelectedRow])
// 		default:
// 			console.Text += e.ID
// 		}
// 		ui.Render(grid)
// 	}
// }

func initializeGrid() {
	grid = ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	// Add widgets to grid and render
	// CurrentState = SelectAction
	grid.Set(
		ui.NewRow(3.0/4, ui.NewCol(1.0, actions)),
		ui.NewRow(1.0/4, ui.NewCol(1.0/2, console), ui.NewCol(1.0, logs)),
	)
	uiEvents := ui.PollEvents()
	ui.Render(grid)
	for {
		e := <-uiEvents
		switch e.ID {

		// Exit on Escape or ctrl-c.
		case "<Escape>", "<C-c>":
			return

		// Redraw grid on window resize
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(grid)

		case "<Enter>":
			// Execute action
			userInput = strings.TrimRight(userInput, "\n")
			fmt.Println("Command:", userInput)

			if len(userInput) > 0 {
				args := strings.Split(userInput, " ")

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
		default:
			if e.Type == ui.KeyboardEvent && len(e.ID) == 1 && e.ID[0] != '<' {
				console.Text += e.ID
				userInput += e.ID
			}
		}
		ui.Render(grid)
	}
}

func createActionsSection() *widgets.Table {
	// List of actions to perform
	actions := widgets.NewTable()
	actions.Title = "Actions"
	actions.Rows = [][]string{
		{"Action", "Example"},
		{"Merge a pull request", "\"merge https://github.com/philip-gai/gh-schedule/pull/1 in 1h30m\""},
	}
	return actions
}

func createJobTable() {
	// Job table
	jobTable := widgets.NewTable()
	jobTable.Title = "Jobs"
	jobTable.Rows = [][]string{
		{"Job", "Created", "Scheduled", "Status"},
	}
	jobTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	jobTable.TextAlignment = ui.AlignCenter
}

func createLogsSection() *widgets.Paragraph {
	logs = widgets.NewParagraph()
	logs.Title = "Logs"
	return logs
}

func createConsole() *widgets.Paragraph {
	// Information and user input
	console = widgets.NewParagraph()
	console.Title = "Console"
	return console
}

// func PopulateGrid() {
// 	// Add widgets to grid and render
// 	grid.Set(
// 		// ui.NewRow(2.0/4, ui.NewCol(1.0, jobTable)),
// 		ui.NewRow(1.0/4, ui.NewCol(1.0, actions)),
// 		// ui.NewRow(1.0/4, ui.NewCol(1.0/2, console), ui.NewCol(1.0/2, logs)),
// 	)
// 	ui.Render(grid)
// }

func runMerge(opts mergeOptions) error {
	fmt.Printf("Scheduling merge of %s in %s\n", opts.PullUrl, opts.In)
	ghCliCmd := []string{"pr", "merge", opts.PullUrl}
	go scheduler.ScheduleJob(scheduler.ScheduleJobOptions{
		In:       opts.In,
		GhCliCmd: ghCliCmd,
	})
	return nil
}