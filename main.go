package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathanaelcunningham/tmuxSessions/commands"
	"github.com/nathanaelcunningham/tmuxSessions/common"
	"github.com/nathanaelcunningham/tmuxSessions/projectList"
	"github.com/nathanaelcunningham/tmuxSessions/sessionInput"
	"github.com/nathanaelcunningham/tmuxSessions/sessionList"
	"github.com/nathanaelcunningham/tmuxSessions/tmux"
)

type model struct {
	state        common.ViewState
	sessionList  sessionList.SessionList
	searchTerm   string
	sessionInput sessionInput.SessionInput
	projectList  projectList.ProjectList
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		pModel, pCmd := m.projectList.Update(msg)
		sModel, sCmd := m.sessionList.Update(msg)
		m.projectList = pModel.(projectList.ProjectList)
		m.sessionList = sModel.(sessionList.SessionList)
		cmds = tea.Batch(cmds, pCmd, sCmd)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case commands.NewSessionCmd:
		m.state = common.NewSession
	case commands.RenameSessionCmd:
		m.state = common.RenameSession
		model := m.sessionInput.SetValue(msg.Value)
		m.sessionInput = model
	case commands.InputDoneCmd:
		if m.state == common.NewSession {
			tmux.NewSession(msg.Value)
			model := m.sessionList.RefreshSessions()
			m.sessionList = model
			m.state = common.SessionList
		} else {
			selected := m.sessionList.Selected()
			tmux.RenameSession(selected, msg.Value)
			model := m.sessionList.RefreshSessions()
			m.sessionList = model
			m.state = common.SessionList
		}
	case commands.InputCancelCmd:
		m.state = common.SessionList
	case commands.ViewSessionsCmd:
		m.state = common.SessionList
	case commands.ViewProjectsCmd:
		m.state = common.ProjectList
	case commands.SaveProjectCmd:
		tmux.SaveSession(string(msg))
		model := m.projectList.RefreshProjects()
		m.projectList = model
	case commands.RestoreProjectCmd:
		model := m.sessionList.RefreshSessions()
		m.sessionList = model
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
	case common.ProjectList:
		model, cmd := m.projectList.Update(msg)
		m.projectList = model.(projectList.ProjectList)
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
	case common.ProjectList:
		return m.projectList.View()
	default:
		return m.sessionList.View()
	}
}

func main() {
	l := sessionList.New()
	p := projectList.New()

	m := model{
		state:        common.SessionList,
		sessionList:  l,
		searchTerm:   "",
		sessionInput: sessionInput.New(),
		projectList:  p,
	}
	prog := tea.NewProgram(m)
	if _, err := prog.Run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
