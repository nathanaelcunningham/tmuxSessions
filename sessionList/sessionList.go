package sessionList

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathanaelcunningham/tmuxSessions/commands"
	"github.com/nathanaelcunningham/tmuxSessions/tmux"
)

const listHeight = 14

type SessionList struct {
	list list.Model
}

func loadSessions(sessions []string) []list.Item {
	var sess []list.Item

	for _, s := range sessions {
		sess = append(sess, SessionItem(s))
	}
	return sess
}
func New() SessionList {
	sess := tmux.GetSessions()
	items := loadSessions(sess)

	l := NewDelegate(items)
	currentSession := tmux.ActiveSessionIndex()
	l.Select(currentSession)

	return SessionList{
		list: l,
	}
}

func (m SessionList) Init() tea.Cmd {
	return nil
}

func (m SessionList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		return m, nil
	case tea.KeyMsg:
		if state := m.list.FilterState(); state != list.Filtering {
			switch {
			case key.Matches(msg, keys.Select):
				i, ok := m.list.SelectedItem().(SessionItem)
				if ok {
					tmux.SwitchSession(string(i))
					return m, tea.Quit
				}
			case key.Matches(msg, keys.ViewProjects):
				cmd := commands.ViewProjects()
				cmds = tea.Batch(cmds, cmd)
			case key.Matches(msg, keys.Delete):
				i, ok := m.list.SelectedItem().(SessionItem)
				index := m.list.Cursor()
				if ok {
					tmux.DeleteSession(string(i))
					m.list.RemoveItem(index)
					m.list.ResetSelected()
				}
			case key.Matches(msg, keys.Rename):
				i, ok := m.list.SelectedItem().(SessionItem)
				index := m.list.Index()
				if ok {
					cmd := commands.RenameSession(index, string(i))
					cmds = tea.Batch(cmds, cmd)
				}
			case key.Matches(msg, keys.New):
				cmd := commands.NewSession()
				cmds = tea.Batch(cmds, cmd)
			case key.Matches(msg, keys.SaveProject):
				i, ok := m.list.SelectedItem().(SessionItem)
				if ok {
					cmd := commands.SaveProject(string(i))
					cmds = tea.Batch(cmds, cmd)
				}
			}
		}
	}
	model, cmd := m.list.Update(msg)
	m.list = model
	cmds = tea.Batch(cmds, cmd)
	return m, cmds
}

func (m SessionList) View() string {
	return m.list.View()
}

func (m SessionList) RefreshSessions() SessionList {
	sessions := tmux.GetSessions()
	items := loadSessions(sessions)
	m.list.SetItems(items)
	return m
}

func (m SessionList) Selected() string {
	i, ok := m.list.SelectedItem().(SessionItem)
	if ok {
		return string(i)
	}
	return ""
}
