package views

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	textinput "github.com/philip-gai/gh-scheduler/ui/components"
)

func Init() {
	p := tea.NewProgram(textinput.InitialModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
