package sessionList

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type SessionListDelegate struct{}

func (d SessionListDelegate) Height() int                             { return 1 }
func (d SessionListDelegate) Spacing() int                            { return 0 }
func (d SessionListDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d SessionListDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(SessionItem)
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

func NewDelegate(items []list.Item) list.Model {
	l := list.New(items, SessionListDelegate{}, 20, listHeight)
	l.Title = "Session List"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	l.AdditionalShortHelpKeys = keys.OverrideKeys

	return l
}
