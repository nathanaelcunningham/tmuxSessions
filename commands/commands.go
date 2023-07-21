package commands

import tea "github.com/charmbracelet/bubbletea"

type NewSessionCmd bool

func NewSession() tea.Cmd {
	return func() tea.Msg {
		return true
	}
}
