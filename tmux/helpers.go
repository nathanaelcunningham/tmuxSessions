package tmux

import "os/exec"

const delimiter = "\t"

func RunCommand(args []string) string {
	out, _ := exec.Command("tmux", args...).Output()
	return string(out)
}
func pane_format() string {
	format := ""
	format += "#{pane_index}"
	format += delimiter
	format += "#{pane_title}"
	format += delimiter
	format += "#{pane_current_path}"
	format += delimiter
	format += "#{pane_active}"
	format += delimiter
	format += "#{pane_id}"

	return format
}

func window_format() string {
	format := ""
	format += "#{window_index}"
	format += delimiter
	format += "#{window_name}"
	format += delimiter
	format += "#{window_active}"
	format += delimiter
	format += "#{window_flags}"
	format += delimiter
	format += "#{window_layout}"
	return format
}
