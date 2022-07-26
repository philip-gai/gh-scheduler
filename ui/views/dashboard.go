package views

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	textinput "github.com/philip-gai/gh-scheduler/ui/components"
)

// https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics
// https://github.com/charmbracelet/bubbles
// https://www.inngest.com/blog/interactive-clis-with-bubbletea

func Init() {
	p := tea.NewProgram(textinput.InitialModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
