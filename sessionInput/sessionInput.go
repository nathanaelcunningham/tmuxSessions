package sessionInput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
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
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m SessionInput) View() string {
	return m.input.View()
}
