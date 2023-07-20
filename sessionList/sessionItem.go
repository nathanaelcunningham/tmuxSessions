package sessionList

type SessionItem string

func (i SessionItem) FilterValue() string { return string(i) }
