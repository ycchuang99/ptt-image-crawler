package tui

import (
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

type Model struct {
	list   list.Model
	Choice crawler.Board
}

func NewBoardSelector() (Model, error) {
	boards, err := crawler.CollectBoardList()
	if err != nil {
		return Model{}, err
	}

	items := make([]list.Item, len(boards))
	for i, b := range boards {
		items[i] = b
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "PTT Boards"
	l.Styles.Title = titleStyle

	return Model{
		list: l,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			if i, ok := m.list.SelectedItem().(crawler.Board); ok {
				m.Choice = i
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

func (m Model) View() string {
	return appStyle.Render(m.list.View())
}
