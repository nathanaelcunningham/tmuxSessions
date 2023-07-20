package commands

import tea "github.com/charmbracelet/bubbletea"

type NewSession interface{}

func NewSessionCmd() tea.Cmd {
	return func() tea.Msg {
		return "new state"
	}
}
