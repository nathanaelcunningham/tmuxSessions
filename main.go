package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathanaelcunningham/tmuxSessions/commands"
	"github.com/nathanaelcunningham/tmuxSessions/common"
	"github.com/nathanaelcunningham/tmuxSessions/sessionInput"
	"github.com/nathanaelcunningham/tmuxSessions/sessionList"
	"github.com/nathanaelcunningham/tmuxSessions/tmux"
)

type model struct {
	state        common.ViewState
	sessionList  sessionList.SessionList
	searchTerm   string
	sessionInput sessionInput.SessionInput
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case commands.NewSessionCmd:
		m.state = common.NewSession
	case commands.NewSessionDoneCmd:
		value := msg.Value
		tmux.NewSession(value)
		model := m.sessionList.RefreshSessions()
		m.sessionList = model
		m.state = common.SessionList
	case commands.RenameSessionDoneCmd:

	case commands.InputCancelCmd:
		m.state = common.SessionList
	}

	switch m.state {
	case common.SessionList:
		model, cmd := m.sessionList.Update(msg)
		m.sessionList = model.(sessionList.SessionList)
		cmds = tea.Batch(cmds, cmd)
	case common.NewSession:
		model, cmd := m.sessionInput.Update(msg)
		m.sessionInput = model.(sessionInput.SessionInput)
		cmds = tea.Batch(cmds, cmd)
	case common.RenameSession:
		model, cmd := m.sessionInput.Update(msg)
		m.sessionInput = model.(sessionInput.SessionInput)
		cmds = tea.Batch(cmds, cmd)
	}
	return m, cmds
}

func (m model) View() string {
	switch m.state {
	case common.NewSession:
		return m.sessionInput.View()
	case common.RenameSession:
		return m.sessionInput.View()
	case common.SessionList:
		return m.sessionList.View()
	default:
		return m.sessionList.View()
	}
}

func main() {
	l := sessionList.New()

	m := model{
		state:        common.SessionList,
		sessionList:  l,
		searchTerm:   "",
		sessionInput: sessionInput.New(),
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
