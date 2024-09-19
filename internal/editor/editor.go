package editor

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type EditorFinishedMsg struct {
	Err error
}

func OpenEditor(notePath string) tea.Cmd {
	ed := os.Getenv("EDITOR")
	if ed == "" {
		// TODO: possibly return errors instead of setting default?
		ed = "vim"
	}

	c := exec.Command(ed, notePath)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return EditorFinishedMsg{err}
	})
}
