package projectList

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	CursorUp     key.Binding
	CursorDown   key.Binding
	Filter       key.Binding
	Delete       key.Binding
	Select       key.Binding
	ViewSessions key.Binding
}

var keys = KeyMap{
	CursorUp: key.NewBinding(
		key.WithKeys("k"),
		key.WithHelp("↑/k", "up"),
	),
	CursorDown: key.NewBinding(
		key.WithKeys("j"),
		key.WithHelp("↓/j", "down"),
	),
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete project"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "load project"),
	),
	ViewSessions: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "View Sessions"),
	),
}

func (k KeyMap) OverrideKeys() []key.Binding {
	return []key.Binding{
		k.Select,
		k.Delete,
		k.ViewSessions,
	}
}
