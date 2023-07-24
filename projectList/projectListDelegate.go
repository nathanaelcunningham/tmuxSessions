package projectList

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ProjectListDelegate struct{}

func (d ProjectListDelegate) Height() int                             { return 1 }
func (d ProjectListDelegate) Spacing() int                            { return 0 }
func (d ProjectListDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ProjectListDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(ProjectItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func NewDelegate(items []list.Item) list.Model {
	l := list.New(items, ProjectListDelegate{}, 20, listHeight)
	l.Title = "Project List"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	l.AdditionalShortHelpKeys = keys.OverrideKeys

	return l
}
