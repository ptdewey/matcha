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
		switch msg.String() {
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

			// set selected note and change to edit mode
			m.SelectedNote = item

			// return m, nil
			return m, editor.OpenEditor(item.Path())
		}
	case tea.WindowSizeMsg:
		// dynamically handle window sizing (fixes no list items showing)
		m.width = msg.Width
		m.height = msg.Height

		listWidth := m.width - ui.OolongStyle.GetHorizontalFrameSize()
		listHeight := m.height - ui.OolongStyle.GetVerticalFrameSize()

		m.List.SetSize(listWidth, listHeight)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func (m Model) viewSearch() string {
	// TODO:
	return ui.OolongStyle.Render(m.List.View())
}
