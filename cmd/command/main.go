package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ycchuang99/ptt-image-crawler/internal/crawler"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)
)

type model struct {
	list   list.Model
	choice crawler.Board
}

func main() {
	m, err := initModel()
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

	if m, ok := finalModel.(model); ok && m.choice.Title() != "" {
		fmt.Printf("\nSelected board: %s\n", m.choice.Title())
		fmt.Printf("Starting crawler for %s...\n", m.choice.Title())
	}
}

func initModel() (model, error) {
	boards, err := crawler.CollectBoardList()
	if err != nil {
		return model{}, err
	}

	items := make([]list.Item, len(boards))
	for i, b := range boards {
		items[i] = b
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "PTT Boards"
	l.Styles.Title = titleStyle

	return model{
		list: l,
	}, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			if i, ok := m.list.SelectedItem().(crawler.Board); ok {
				m.choice = i
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}
