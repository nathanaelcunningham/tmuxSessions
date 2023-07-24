package projects

import (
	"fmt"
	"testing"
)

func TestLoadProjects(t *testing.T) {
	projects, err := LoadProjects()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(projects)
}
