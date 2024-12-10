package main

import (
	"journal/journal"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := journal.NewModel("")

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run tui %v", err)
	}
}


