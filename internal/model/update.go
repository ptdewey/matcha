package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/editor"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		k := msg.String()
		if m.Mode == SEARCH {
			if m.List.SettingFilter() || m.List.IsFiltered() {
				if k == "esc" {
					m.List.ResetFilter()
					return m, nil
				}
			}
		}
		if k == "esc" || k == "ctrl+c" || k == "q" {
			return m, tea.Quit
		}
	case editor.EditorFinishedMsg:
		// m.Mode = BROWSE
		m.Mode = SEARCH
		m.List.ResetFilter()
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
	case BROWSE:
		// TODO: browse mode adaptations
		return m.updateSearch(msg)
	case SEARCH:
		// TODO: send a key message at some point
		return m.updateSearch(msg)
	}

	return m, nil
}
