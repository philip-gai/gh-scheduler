package tui

import (
	"log"
	"strings"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/philip-gai/gh-scheduler/scheduler"
	"github.com/philip-gai/gh-scheduler/utils"
)

var grid *termui.Grid
var console *widgets.List
var logs *widgets.List
var jobTable *widgets.Table
var userInput string

func Render() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

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

		if args[0] == "gh" {
			args = args[1:]
		}

		argLen := len(args)

		if argLen == 0 {
			utils.PushListRow("No command provided", logs)
			return
		} else if argLen < 3 {
			utils.PushListRow("Command must contain \"<cmd> in <duration>\"", logs)
			return
		}
		timeDuration := args[len(args)-1]

		// Remove scheduling info from command
		ghCliArgs := args[:len(args)-2]

		scheduler.ScheduleJob(scheduler.ScheduleJobOptions{
			In:       timeDuration,
			GhCliCmd: ghCliArgs,
		}, logs)
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
		termui.NewRow(2.0/4, termui.NewCol(1.0/2, console), termui.NewCol(1.0/2, logs)),
	)
	termui.Render(grid)
	return grid
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
	logs.WrapText = true
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
