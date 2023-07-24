package projects

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/nathanaelcunningham/tmuxSessions/tmux"
)

type Project struct {
	Name     string
	Filepath string
}

func LoadProjects() ([]Project, error) {
	configDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/.config/tmuxSessions", configDir)

	filenames := []Project{}

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			name := fileWithoutExtension(info.Name())
			filenames = append(filenames, Project{
				Name:     name,
				Filepath: path,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filenames, nil
}

func fileWithoutExtension(fileName string) string {
	return path.Base(fileName[:len(fileName)-len(path.Ext(fileName))])
}

func RestoreProject(project Project) error {
	err := tmux.RestoreSession(project.Filepath)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProject(project Project) error {
	err := os.Remove(project.Filepath)
	if err != nil {
		return err
	}

	return nil
}
