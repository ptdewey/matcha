package notes

import (
	"errors"
	"fmt"
	"os"
	"text/template"

	tea "github.com/charmbracelet/bubbletea"
)

type CreateFinishedMsg struct {
	Err error
}

func createFromTemplate(notePath string, templatePath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	f, err := os.Create(notePath)
	if err != nil {
		return err
	}

	tmpl.ExecuteTemplate()

	// TODO: templates (as defined in config)

	return nil
}

func CreateNote(notePath string) tea.Msg {
	f, err := os.Stat(notePath)
	if errors.Is(err, os.ErrNotExist) {
		// TODO: create file? open editor
	} else if err != nil {
		return CreateFinishedMsg{err}
	}

	if f.IsDir() {
		return CreateFinishedMsg{
			errors.New(fmt.Sprintf("Error: %s is a directory.", notePath)),
		}
	}

	return CreateFinishedMsg{nil}
}
