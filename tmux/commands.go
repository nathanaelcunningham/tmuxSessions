package tmux

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetSessions() []string {
	cmd := exec.Command("tmux", "list-sessions")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
	split := strings.Split(string(out), "\n")

	var sessions []string

	for _, s := range split {
		filter := strings.SplitN(s, ":", 2)
		if len(filter[0]) > 0 {
			sessions = append(sessions, filter[0])
		}
	}
	return sessions
}

func ActiveSession() string {
	cmd := exec.Command("tmux", "list-sessions")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
	split := strings.Split(string(out), "\n")

	activeSession := ""

	for _, s := range split {
		if strings.Contains(s, "attached") {
			filter := strings.SplitN(s, ":", 2)
			activeSession = filter[0]
		}
	}

	return activeSession
}
func ActiveSessionIndex() int {
	cmd := exec.Command("tmux", "list-sessions")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
	split := strings.Split(string(out), "\n")

	index := 0

	for i, s := range split {
		if strings.Contains(s, "attached") {
			index = i
		}
	}

	return index
}

func SwitchSession(session string) {
	cmd := exec.Command("tmux", "switch-client", "-t", session)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
	fmt.Println(string(out))
}

func NewSession(session string) {
	cmd := exec.Command("tmux", "new", "-d", "-s", session)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
}

func DeleteSession(session string) {
	cmd := exec.Command("tmux", "kill-session", "-t", session)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
}

func RenameSession(session, newSession string) {
	cmd := exec.Command("tmux", "rename-session", "-t", session, newSession)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
}
func GetWindows(session string) []Window {
	out, err := exec.Command("tmux", "list-windows", "-t", session).Output()
	if err != nil {
		return []Window{}
	}

	// Match groups: index, name, layout, ID
	re := regexp.MustCompile(`^(\d+): (\w+[\*\-]?).*?(\[layout [^\]]+\]*) (@\d+)`)

	var windows []Window
	lines := strings.Split(string(out), "\n")
	for _, line := range lines[:len(lines)-1] {
		groups := re.FindStringSubmatch(line)
		if groups != nil {
			index, _ := strconv.Atoi(groups[1])
			active := false
			if strings.Contains(groups[2], "*") {
				active = true
			}
			name := strings.Trim(groups[2], "* ") // remove '*' sign
			name = strings.Trim(groups[2], "- ")  // remove '-' sign

			layout := groups[3]
			id := groups[4]

			windows = append(windows, Window{Index: index, Name: name, Active: active, Layout: layout, ID: id})
		}
	}
	return windows
}

func GetPanes(session string, windowIndex int) []Pane {
	out, err := exec.Command("tmux", "list-panes", "-t", fmt.Sprintf("%s:%d", session, windowIndex)).Output()
	if err != nil {
		return []Pane{}
	}

	re := regexp.MustCompile(`^(\d+): (\[\d+x\d+\]) \[history (\d+/\d+), \d+ bytes\] (%\d+)(?: \((\w+)\))?$`)

	var panes []Pane
	lines := strings.Split(string(out), "\n")
	for _, line := range lines[:len(lines)-1] {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			// Index
			indexStr := matches[1]
			index, _ := strconv.Atoi(indexStr)
			// Size
			sizeStr := matches[2]
			// ID
			idStr := matches[4]
			// Active
			activeStr := matches[5]
			active := false
			if activeStr != "" {
				active = true
			}
			panes = append(panes, Pane{
				Index:  index,
				Size:   sizeStr,
				ID:     idStr,
				Active: active,
			})
		}
	}
	return panes
}
