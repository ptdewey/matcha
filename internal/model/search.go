package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/data"
	"github.com/ptdewey/oolong/internal/editor"
	"github.com/ptdewey/oolong/internal/ui"
)

func (m Model) updateSearch(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch k := msg.String(); k {
		case "enter":
			if !m.List.SettingFilter() {
				selectedItem := m.List.SelectedItem()
				if selectedItem == nil {
					return m, nil
				}

				item, ok := selectedItem.(data.Note)
				if !ok {
					fmt.Println("Error: could not type assert to data.Note")
					return m, nil
				}

				// set selected note and change to edit mode
				m.SelectedNote = item

				// open not for editing
				return m, editor.OpenEditor(item.Path())
			}
		default:
			if !m.List.SettingFilter() && (k == "q" || k == "esc") {
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		// dynamically handle window sizing (fixes no list items showing)
		h, v := ui.OolongStyle.GetFrameSize()
		m.width, m.height = msg.Width-h, msg.Height-v
		m.List.SetSize(m.width, m.height)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func (m Model) viewSearch() string {
	// TODO: customize views
	return ui.OolongStyle.Render(m.List.View())
}
