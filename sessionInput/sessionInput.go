package sessionInput

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathanaelcunningham/tmuxSessions/commands"
)

type SessionInput struct {
	input textinput.Model
}

func New() SessionInput {
	i := textinput.New()

	return SessionInput{
		input: i,
	}
}

func (m SessionInput) Init() tea.Cmd {
	return nil
}

func (m SessionInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.input.Focus()
	var cmds tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Done):
			cmd := commands.InputDone(m.input.Value())
			m.input.Reset()
			cmds = tea.Batch(cmds, cmd)
		case key.Matches(msg, keys.Cancel):
			m.input.Reset()
			cmd := commands.InputCancel()
			cmds = tea.Batch(cmds, cmd)
		}
	}

	model, cmd := m.input.Update(msg)
	m.input = model
	cmds = tea.Batch(cmds, cmd)
	return m, cmds
}

func (m SessionInput) View() string {
	return m.input.View()
}

func (m SessionInput) SetValue(value string) SessionInput {
	m.input.SetValue(value)
	return m
}
