package tmux

import (
	"fmt"
	"strconv"
	"strings"
)

func WindowExists(session string, windowIndex int) bool {
	windows := GetWindows(session)
	for _, win := range windows {
		if win.Index == windowIndex {
			return true
		}
	}
	return false
}

func GetWindows(session string) []Window {
	out := RunCommand([]string{"list-windows", "-t", session, "-F", window_format()})

	var windows []Window
	lines := strings.Split(string(out), "\n")
	for _, line := range lines[:len(lines)-1] {
		split := strings.Split(line, "\t")
		index, _ := strconv.Atoi(split[0])
		name := split[1]
		active, _ := strconv.ParseBool(split[2])
		flags := split[3]
		layout := split[4]

		windows = append(windows, Window{Index: index, Name: name, Active: active, Flags: flags, Layout: layout})
	}
	return windows
}

func NewWindow(session string, window Window) {
	cmd := []string{"new-window", "-d", "-t", fmt.Sprintf("%s:%d", session, window.Index), "-n", window.Name}
	RunCommand(cmd)
}
