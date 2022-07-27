package views

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/philip-gai/gh-scheduler/ui/components/textinput"
)

// https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics
// https://github.com/charmbracelet/bubbles
// https://www.inngest.com/blog/interactive-clis-with-bubbletea
// https://github.com/dlvhdr/gh-dash/blob/main/ui/ui.go

func Init() {
	p := tea.NewProgram(textinput.New())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
