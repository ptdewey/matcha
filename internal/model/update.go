package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "esc" || k == "ctrl+c" || k == "q" {
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
		// TODO: search handler
		// FIX: there is something wrong with the list view
		return m.updateSearch(msg)
	case EDIT:
		// TODO: open for editing
		return m, nil
	}

	// TODO: handle different update modes post landing

	return m, nil
}
