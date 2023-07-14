package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getSessions() []string {
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

func activeSession() string {
	cmd := exec.Command("tmux", "list-sessions")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
	split := strings.Split(string(out), "\n")

	activeSession := ""

	for _, s := range split {
		if strings.Contains(s, "active") {
			filter := strings.SplitN(s, ":", 2)
			activeSession = filter[0]
		}
	}

	return activeSession
}
func activeSessionIndex() int {
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

func switchSession(session string) {
	cmd := exec.Command("tmux", "switch-client", "-t", session)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
	fmt.Println(string(out))
}

func newSession(session string) {
	cmd := exec.Command("tmux", "new", "-d", "-s", session)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
}

func deleteSession(session string) {
	cmd := exec.Command("tmux", "kill-session", "-t", session)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
}

func renameSession(session, newSession string) {
	cmd := exec.Command("tmux", "rename-session", "-t", session, newSession)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
}
