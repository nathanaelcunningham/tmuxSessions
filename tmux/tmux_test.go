package tmux

import (
	"fmt"
	"testing"
)

func TestSessionExists(t *testing.T) {
	session := "tmuxSessions"

	got := SessionExists(session)
	want := true

	if got != want {
		t.Error("failed")
	}

}
func TestWindowExists(t *testing.T) {
	session := ActiveSession()
	index := 0

	got := WindowExists(session, index)
	want := true

	if got != want {
		t.Error("failed")
	}
}

func TestPaneExists(t *testing.T) {
	session := ActiveSession()
	windowIndex := 0
	paneIndex := 0

	got := PaneExists(session, windowIndex, paneIndex)
	want := true

	if got != want {
		t.Error("failed")
	}
}

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

func TestRestoreSession(t *testing.T) {
	path := "/users/nathanael/.config/tmuxSessions/tmuxSessions.json"
	err := RestoreSession(path)
	if err != nil {
		t.Error(err)
	}
}
