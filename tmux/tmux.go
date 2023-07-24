package tmux

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load from file
func RestoreSession(filepath string) error {
	session, err := loadFromFile(filepath)
	if err != nil {
		return err
	}
	fmt.Println(session)

	return nil
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

func loadFromFile(filepath string) (Session, error) {
	var session Session
	file, err := os.Open(filepath)
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeToFile(session Session) error {
	filename := fmt.Sprintf("%s.json", session.Name)

	configDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(fmt.Sprintf("%s/.config/tmuxSessions", configDir), os.ModePerm)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/.config/tmuxSessions/%s", configDir, filename)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(session)
}
