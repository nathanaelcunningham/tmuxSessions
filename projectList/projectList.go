package projectList

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathanaelcunningham/tmuxSessions/commands"
	"github.com/nathanaelcunningham/tmuxSessions/projects"
	"github.com/nathanaelcunningham/tmuxSessions/tmux"
)

const listHeight = 14

type ProjectList struct {
	list list.Model
}

func loadProjects(projects []projects.Project) []list.Item {
	var proj []list.Item

	for _, p := range projects {
		proj = append(proj, ProjectItem(ProjectItem{
			Name: p.Name,
			Path: p.Filepath,
		}))
	}
	return proj
}
func New() ProjectList {
	projects, _ := projects.LoadProjects()
	items := loadProjects(projects)

	l := NewDelegate(items)
	currentProject := tmux.ActiveSessionIndex()
	l.Select(currentProject)

	return ProjectList{
		list: l,
	}
}

func (m ProjectList) Init() tea.Cmd {
	return nil
}

func (m ProjectList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		return m, nil
	case commands.SaveProjectCmd:
		err := tmux.SaveSession(string(msg))
		if err != nil {
			log.Fatal(err)
		}
		m.RefreshProjects()
	case tea.KeyMsg:
		if state := m.list.FilterState(); state != list.Filtering {
			switch {
			case key.Matches(msg, keys.Select):
				i, ok := m.list.SelectedItem().(ProjectItem)
				if ok {
					projects.RestoreProject(projects.Project{
						Name:     i.Name,
						Filepath: i.Path,
					})
				}
			case key.Matches(msg, keys.ViewSessions):
				cmd := commands.ViewSessions()
				cmds = tea.Batch(cmds, cmd)
			case key.Matches(msg, keys.Delete):
				// 	i, ok := m.list.SelectedItem().(ProjectItem)
				// 	index := m.list.Cursor()
				// 	if ok {
				// 		tmux.DeleteProject(string(i))
				// 		m.list.RemoveItem(index)
				// 		m.list.ResetSelected()
				// 	}
				// case key.Matches(msg, keys.Rename):
				// 	i, ok := m.list.SelectedItem().(ProjectItem)
				// 	index := m.list.Index()
				// 	if ok {
				// 		cmd := commands.RenameProject(index, string(i))
				// 		cmds = tea.Batch(cmds, cmd)
				// 	}
			}
		}
	}
	model, cmd := m.list.Update(msg)
	m.list = model
	cmds = tea.Batch(cmds, cmd)
	return m, cmds
}

func (m ProjectList) View() string {
	return m.list.View()
}

func (m ProjectList) RefreshProjects() ProjectList {
	projects, _ := projects.LoadProjects()
	items := loadProjects(projects)
	m.list.SetItems(items)
	return m
}

func (m ProjectList) Selected() string {
	i, ok := m.list.SelectedItem().(ProjectItem)
	if ok {
		return i.Name
	}
	return ""
}
