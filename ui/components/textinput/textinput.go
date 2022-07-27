package textinput

// A simple program demonstrating the text input component from the Bubbles
// component library.
// https://github.com/charmbracelet/bubbletea/blob/master/examples/textinput/main.go

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type Model struct {
	textInput textinput.Model
	err       error
}

func New() tea.Model {
	ti := textinput.New()
	ti.Placeholder = "gh pr merge <url> in 1h"
	ti.Focus()

	return Model{
		textInput: ti,
		err:       nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:

		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf(
		"Command to schedule: \n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

// func runCommand(m Model) {
// 	// Re-enable once we have debug logging
// 	// utils.PushListRow(fmt.Sprintf("Running command \"%s\"", userInput), logs)

// 	// Execute action
// 	userInput := m.textInput.Value()

// 	if len(userInput) > 0 {
// 		args := strings.Split(userInput, " ")

// 		if args[0] == "gh" {
// 			args = args[1:]
// 		}

// 		argLen := len(args)

// 		if argLen == 0 {
// 			utils.PushListRow("No command provided", logs)
// 			utils.PushListRow("$ ", console)
// 			return
// 		}

// 		hasTime := len(args) >= 2 && args[len(args)-2] == "in"
// 		timeDuration := "0s"

// 		ghCliArgs := args

// 		if hasTime {
// 			timeDuration = args[len(args)-1]
// 			// Remove scheduling info from command
// 			ghCliArgs = args[:len(args)-2]
// 		}

// 		scheduler.ScheduleJob(scheduler.ScheduleJobOptions{
// 			In:       timeDuration,
// 			GhCliCmd: ghCliArgs,
// 		}, logs, jobTable)
// 	}
// 	utils.PushListRow("$ ", console)
// }
