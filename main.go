package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathanaelcunningham/tmuxSessions/commands"
	"github.com/nathanaelcunningham/tmuxSessions/common"
	"github.com/nathanaelcunningham/tmuxSessions/sessionInput"
	"github.com/nathanaelcunningham/tmuxSessions/sessionList"
)

type model struct {
	state         common.ViewState
	sessionList   sessionList.SessionList
	searchTerm    string
	activeSession string
	addNew        bool
	rename        bool
	sessionInput  sessionInput.SessionInput
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	switch msg.(type) {
	case commands.NewSession:
		m.state = common.NewSession
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
	}
	return m, cmds
}

func (m model) View() string {
	switch m.state {
	case common.SessionList:
		return m.sessionList.View() + fmt.Sprintf(" State: %d\n", m.state)
	default:
		return m.sessionList.View() + fmt.Sprintf(" State: %d\n", m.state)
	}
}

func main() {
	l := sessionList.New()

	m := model{
		state:        common.SessionList,
		sessionList:  l,
		searchTerm:   "",
		addNew:       false,
		sessionInput: sessionInput.New(),
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
