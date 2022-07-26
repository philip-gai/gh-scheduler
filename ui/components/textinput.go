package components

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

func InitialModel() Model {
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
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
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
