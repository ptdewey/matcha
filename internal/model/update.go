package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/editor"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		k := msg.String()
		// NOTE: mode should always be browse at this point
		// if m.Mode == SEARCH || m.Mode == BROWSE {
		if m.Mode == BROWSE {
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
		m.Mode = BROWSE
		m.List.ResetFilter()
		if msg.Err != nil {
			return m, tea.Quit
		}
	}

	// mode is chosen on landing page
	if !m.ModeChosen {
		return m.updateLanding(msg)
	}

	switch m.Mode {
	case CREATE:
		return m.updateCreate(msg)
	case BROWSE:
		return m.updateBrowse(msg)
	case SEARCH:
		// set mode to browse for next iteration
		m.Mode = BROWSE

		// call update with keymsg to trigger filter mode on enter
		m, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}})

		// call update and ensure windowsize event is called
		return m.(Model).updateBrowse(msg)
	}

	return m, nil
}
