package tmux

type Session struct {
	Name    string
	Windows []Window
}

type Window struct {
	Index  int
	Name   string
	Active bool
	Layout string
	ID     string
	Panes  []Pane
}

type Pane struct {
	Index  int
	Size   string
	ID     string
	Active bool
}
