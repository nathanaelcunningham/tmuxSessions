package tmux

import (
	"fmt"
	"testing"
)

func TestGetWindows(t *testing.T) {
	session := ActiveSession()
	GetWindows(session)
}

func TestGetPanes(t *testing.T) {
	session := ActiveSession()
	windows := GetWindows(session)

	panes := GetPanes(session, windows[0].Index)
	fmt.Println("panes: \n", panes)
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

func TestLoadSession(t *testing.T){

}
