package sessionList

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	CursorUp   key.Binding
	CursorDown key.Binding
	Filter     key.Binding
	Delete     key.Binding
	New        key.Binding
	Rename     key.Binding
	Select     key.Binding
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
		key.WithHelp("d", "delete session"),
	),
	New: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new session"),
	),
	Rename: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "rename session"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select session"),
	),
}

func (k KeyMap) OverrideKeys() []key.Binding {
	return []key.Binding{
		k.New,
		k.Rename,
		k.Delete,
	}
}
