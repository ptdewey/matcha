package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/data"
	"github.com/ptdewey/oolong/internal/ui"
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
		return updateLanding(msg, m)
	}

	switch m.Mode {
	case CREATE:
		// TODO: handle note creation
		return m, nil
	case SEARCH:
		// TODO: search handler
		return m, nil
	}

	// TODO: handle different update modes post landing

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// TODO: wrap this entire section up in a function somewhere else ()
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.List.FilterState() == list.Filtering {
				m.List.ResetFilter()
				return m, nil
			}
		case "enter":
			selectedItem := m.List.SelectedItem()
			if selectedItem == nil {
				return m, nil
			}

			item, ok := selectedItem.(data.Note)
			if !ok {
				fmt.Println("Error: could not type assert to data.Item")
				return m, nil
			}

			// TODO: open/create note
			_ = item

			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := ui.Style.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	default:
		panic(fmt.Sprintf("unexpected tea.Msg: %#v", msg))
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}
