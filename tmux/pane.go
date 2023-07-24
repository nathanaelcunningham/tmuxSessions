package tmux

import (
	"fmt"
	"strconv"
	"strings"
)

func NewPane(session string, windowIndex int, pane Pane) {
	cmd := []string{"split-window", "-t", fmt.Sprintf("%s:%d", session, windowIndex)}
	if pane.Path != "" {
		cmd = append(cmd, "-c", pane.Path)
	}
	RunCommand(cmd)
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

		panes = append(panes, Pane{
			Index:  index,
			Title:  title,
			Path:   path,
			Active: active,
		})
	}
	return panes
}
