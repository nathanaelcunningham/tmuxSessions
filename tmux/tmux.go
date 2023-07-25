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

	//Create session
	if !SessionExists(session.Name) {
		NewSession(session.Name)
	}
	//Create Windows
	for _, window := range session.Windows {
		if !WindowExists(session.Name, window.Index) {
			NewWindow(session.Name, window)
		}
		//Create Panes
		for _, pane := range window.Panes {
			paneExists := PaneExists(session.Name, window.Index, pane.Index)
			if paneExists == false {
				NewPane(session.Name, window.Index, pane)
			} else {
				RespawnPane(session.Name, window.Index, pane)
			}
			ResizePane(session.Name,window.Index)
		}
		//Restore Window Layout
		RestoreLayout(session.Name, window)
	}

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
