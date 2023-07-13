package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 15

type session string
type sessionDelegate struct{}

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

func (i session) FilterValue() string { return string(i) }

func (d sessionDelegate) Height() int                             { return 1 }
func (d sessionDelegate) Spacing() int                            { return 0 }
func (d sessionDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d sessionDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(session)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}


func loadSessions() []list.Item {
	cmd := exec.Command("tmux", "list-sessions")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
	split := strings.Split(string(out), "\n")

	var sessions []list.Item

	for _, s := range split {
		filter := strings.SplitN(s, ":", 2)
		if len(filter[0]) > 0 {
			sessions = append(sessions, session(filter[0]))
		}
	}
	return sessions
}

func switchSession(session string) {
	cmd := exec.Command("tmux", "switch-client", "-t", session)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
	fmt.Println(string(out))
}

func deleteSession(session string) {
	cmd := exec.Command("tmux", "kill-session", "-t", session)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to run command")
	}
}
