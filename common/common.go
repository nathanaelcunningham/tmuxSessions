package common

type ViewState int

const (
	NewSession ViewState = iota
	RenameSession
	SessionList
	ProjectList
)
