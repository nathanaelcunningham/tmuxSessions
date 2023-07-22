package tmux

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load from file
func LoadSession(name string) Session {
	return Session{}
}

// Save to file
func SaveSession(name string) error {
	session := ConvertSession(name)
	err := writeToFile(session)
	if err != nil {
		return err
	}
	return nil
}

// Get from tmux server
func ConvertSession(name string) Session {
	session := Session{
		Name: name,
	}
	windows := GetWindows(session.Name)

	for i, window := range windows {
		panes := GetPanes(session.Name, window.Index)
		windows[i].Panes = panes
	}

	session.Windows = windows

	return session
}

func writeToFile(session Session) error {
	filename := fmt.Sprintf("%s.json", session.Name)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(session)
}
