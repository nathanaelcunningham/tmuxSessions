package tmux

import (
	"fmt"
	"testing"
)

func TestGetWindows(t *testing.T) {
	session := ActiveSession()
	windows := GetWindows(session)
	fmt.Printf("%+v\n", windows)
}

func TestGetPanes(t *testing.T) {
	session := ActiveSession()
	windows := GetWindows(session)

	panes := GetPanes(session, windows[0].Index)
	fmt.Printf("panes: \n %+v\n", panes)
}

func TestConvertSession(t *testing.T) {
	name := ActiveSession()
	session := ConvertSession(name)

	fmt.Printf("%+v\n", session)
}

func TestSaveSession(t *testing.T) {
	name := ActiveSession()
	err := SaveSession(name)
	if err != nil {
		t.Error(err)
	}
}

func TestNewWindow(t *testing.T) {
	name := ActiveSession()
	window := Window{
		Index:  2,
		Name:   "Test",
		Active: false,
		Flags:  "",
		Layout: "",
	}
	NewWindow(name, window)
}
