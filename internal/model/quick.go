package model

import (
	"fmt"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/matcha/internal/editor"
	"github.com/ptdewey/matcha/internal/notes"
)

func (m Model) updateQuick(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		m.err = msg
		return m, nil
	}

	notePath := m.checkJoinInputs()

	ext := filepath.Ext(notePath)
	templateSource := ""
	for _, template := range m.Templates {
		fmt.Println(template.Ext, ext)
		if template.Ext == ext {
			templateSource = template.Name
			fmt.Println(templateSource)
			break
		}
	}

	if err := notes.CreateNote(notePath, templateSource); err != nil {
		fmt.Println("Error creating note:", err)
		return m, nil
	}

	// HACK: see if there is a better way of handling this (to avoid editor loop)
	// - open editor message should be sent but seems to be delayed?
	// - maybe this could be handled with a message in this function?
	//   - not an issue in create.go, likely because wrapped in keymsg handler
	m.ModeChosen = false

	return m, editor.OpenEditor(notePath)
}

func (m Model) viewQuick() string {
	return ""
}
