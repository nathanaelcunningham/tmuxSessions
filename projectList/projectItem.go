package projectList

type ProjectItem struct {
	Name string
	Path string
}

func (i ProjectItem) FilterValue() string { return i.Name }
