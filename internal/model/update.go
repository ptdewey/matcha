package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/editor"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		k := msg.String()
		if k == "esc" || k == "ctrl+c" || k == "q" {
			return m, tea.Quit
		}
	case editor.EditorFinishedMsg:
		// m.Mode = BROWSE
		m.Mode = SEARCH
		if msg.Err != nil {
			m.err = msg.Err
			return m, tea.Quit
		}
	}

	// mode is chosen on landing page
	if !m.ModeChosen {
		return m.updateLanding(msg)
	}

	switch m.Mode {
	case CREATE:
		// TODO: handle note creation
		return m, nil
	case SEARCH:
		return m.updateSearch(msg)
	}

	return m, nil
}
