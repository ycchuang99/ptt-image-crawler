package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ycchuang99/ptt-image-crawler/internal/tui"
)

func main() {
	m, err := tui.NewBoardSelector()
	if err != nil {
		fmt.Println("Error initializing model:", err)
		return
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		return
	}

	if m, ok := finalModel.(tui.Model); ok && m.Choice.Title() != "" {
		fmt.Printf("\nSelected board: %s\n", m.Choice.Title())
		fmt.Printf("Starting crawler for %s...\n", m.Choice.Title())
	}
}
