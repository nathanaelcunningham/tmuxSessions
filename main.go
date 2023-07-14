package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	sessionList  list.Model
	searchTerm   string
	addNew       bool
	rename       bool
	sessionInput textinput.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.sessionList.SetSize(msg.Width-h, msg.Height-v)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.addNew = false
			m.rename = false
			return m, nil
		case "enter":
			if m.addNew {
				newSession(m.sessionInput.Value())
				cmd = m.sessionList.SetItems(loadSessions())
				m.addNew = false
				return m, cmd
			}
			if m.rename {
				i, ok := m.sessionList.SelectedItem().(session)
				if ok {
					renameSession(string(i), m.sessionInput.Value())
					cmd = m.sessionList.SetItems(loadSessions())
					m.rename = false
					return m, cmd
				}
			}
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
		case "r":
			i, ok := m.sessionList.SelectedItem().(session)
			if ok {
				m.rename = true
				m.sessionInput.SetValue(string(i))
				m.sessionInput.Focus()
			}
		case "n":
			m.addNew = true
			m.sessionInput.Focus()
			return m, nil
		}
	}
	if m.addNew || m.rename {
		m.sessionInput, cmd = m.sessionInput.Update(msg)
	} else {
		m.sessionList, cmd = m.sessionList.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	newSession := fmt.Sprintf("New Session Name\n\n%s\n\n%s\n", m.sessionInput.View(), "(esc to cancel)")
	renameSession := fmt.Sprintf("Rename Session\n\n%s\n\n%s\n", m.sessionInput.View(), "(esc to cancel)")
	if m.addNew {
		return newSession
	}
	if m.rename {
		return renameSession
	}
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
		sessionList:  l,
		searchTerm:   "",
		addNew:       false,
		sessionInput: textinput.New(),
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
