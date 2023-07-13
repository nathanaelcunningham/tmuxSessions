package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	sessionList list.Model
	searchTerm  string
	addNew      bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.sessionList.SetSize(msg.Width-h, msg.Height-v)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i, ok := m.sessionList.SelectedItem().(session)
			if ok {
				switchSession(string(i))
				return m, tea.Quit
			}
		case "d":
			i, ok := m.sessionList.SelectedItem().(session)
			index := m.sessionList.Cursor()
			if ok {
				deleteSession(string(i))
				m.sessionList.RemoveItem(index)
				m.sessionList.ResetSelected()
				return m, nil
			}
		}
	}
	var cmd tea.Cmd
	m.sessionList, cmd = m.sessionList.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.sessionList.View()
}

func additionalKeys() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete session"),
		),
		key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "new session"),
		),
	}
}

func main() {
	sessions := loadSessions()

	const defaultWidth = 20
	l := list.New(sessions, sessionDelegate{}, 20, listHeight)
	l.AdditionalShortHelpKeys = additionalKeys

	l.Title = "Session List"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{
		sessionList: l,
		searchTerm:  "",
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
