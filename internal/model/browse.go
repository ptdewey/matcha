package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/data"
	"github.com/ptdewey/oolong/internal/editor"
)

func (m Model) updateBrowse(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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
		}
	case tea.WindowSizeMsg:
		// dynamically handle window sizing (fixes no list items showing)
		h, v := OolongStyle.GetFrameSize()
		m.width, m.height = msg.Width-h, msg.Height-v
		m.List.SetSize(m.width, m.height)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func (m Model) viewBrowse() string {
	// TODO: customize views
	return OolongStyle.Render(m.List.View())
}
