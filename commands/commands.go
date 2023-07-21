package commands

import tea "github.com/charmbracelet/bubbletea"

type NewSessionCmd bool

func NewSession() tea.Cmd {
	return func() tea.Msg {
		return NewSessionCmd(true)
	}
}

type RenameSessionCmd struct {
	Index int
	Value string
}

func RenameSession(index int, value string) tea.Cmd {
	return func() tea.Msg {
		return RenameSessionCmd{
			Index: index,
			Value: value,
		}
	}
}

type NewSessionDoneCmd struct {
	Value string
}

func NewSessionDone(value string) tea.Cmd {
	return func() tea.Msg {
		return NewSessionDoneCmd{
			Value: value,
		}
	}
}

type RenameSessionDoneCmd struct {
	Index int
	Value string
}

type InputCancelCmd bool

func InputCancel() tea.Cmd {
	return func() tea.Msg {
		return InputCancelCmd(true)
	}
}
