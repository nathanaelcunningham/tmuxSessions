package tmux

import (
	"fmt"
	"strconv"
	"strings"
)

func NewPane(session string, windowIndex int, pane Pane) {
	cmd := []string{"split-window", "-c", pane.Path, "-t", fmt.Sprintf("%s:%d", session, windowIndex)}
	RunCommand(cmd)
}
func RespawnPane(session string, windowIndex int, pane Pane) {
	panes := GetPanes(session, windowIndex)
	NewPane(session, windowIndex, pane)
	KillPane(panes[0].ID)
}
func ResizePane(session string, windowIndex int) {
	RunCommand([]string{"resize-pane", "-t", fmt.Sprintf("%s:%d", session, windowIndex), "-U", "999"})
}

func KillPane(paneID string) {
	cmd := []string{"kill-pane", "-t", paneID}
	RunCommand(cmd)
}

func PaneExists(session string, windowIndex, paneIndex int) bool {
	panes := GetPanes(session, windowIndex)
	for _, p := range panes {
		if p.Index == paneIndex {
			return true
		}
	}
	return false
}

func PaneExistsByID(session string, windowIndex int, paneID string) bool {
	panes := GetPanes(session, windowIndex)
	for _, p := range panes {
		if p.ID == paneID {
			return true
		}
	}
	return false
}

func GetPanes(session string, windowIndex int) []Pane {
	out := RunCommand([]string{"list-panes", "-t", fmt.Sprintf("%s:%d", session, windowIndex), "-F", pane_format()})

	var panes []Pane

	lines := strings.Split(string(out), "\n")

	for _, line := range lines[:len(lines)-1] {
		split := strings.Split(line, "\t")
		index, _ := strconv.Atoi(split[0])
		title := split[1]
		path := split[2]
		active, _ := strconv.ParseBool(split[3])
		id := split[4]

		panes = append(panes, Pane{
			Index:  index,
			Title:  title,
			Path:   path,
			Active: active,
			ID:     id,
		})
	}
	return panes
}

func GetPaneID(session string, windowIndex int) string {
	return RunCommand([]string{"display-message", "-p", "-F", "#{pane_id}", "-t", fmt.Sprintf("%s:%d", session, windowIndex)})
}
