package sessionInput

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Done   key.Binding
	Cancel key.Binding
}

var keys = KeyMap{
	Done: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Enter"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Esc"),
	),
}
