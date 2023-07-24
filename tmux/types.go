package tmux

type Session struct {
	Name    string
	Windows []Window
}

type Window struct {
	Index  int
	Name   string
	Active bool
	Flags  string
	Layout string
	Panes  []Pane
}

type Pane struct {
	Index  int
	Title  string
	Path   string
	Active bool
}
