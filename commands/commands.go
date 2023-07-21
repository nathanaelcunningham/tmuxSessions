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

type InputDoneCmd struct {
	Value string
}

func InputDone(value string) tea.Cmd {
	return func() tea.Msg {
		return InputDoneCmd{
			Value: value,
		}
	}
}

type InputCancelCmd bool

func InputCancel() tea.Cmd {
	return func() tea.Msg {
		return InputCancelCmd(true)
	}
}
