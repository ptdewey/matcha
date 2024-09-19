package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/data"
	"github.com/ptdewey/oolong/internal/ui"
)

func (m Model) updateSearch(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// TODO: wrap this entire section up in a function somewhere else ()
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

			// TODO: open/create note
			fmt.Println(item.Title())

			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		// Handle resizing dynamically
		m.width = msg.Width
		m.height = msg.Height

		// Calculate the dynamic size for the list, considering styling
		listWidth := m.width - ui.OolongStyle.GetHorizontalFrameSize()
		listHeight := m.height - ui.OolongStyle.GetVerticalFrameSize()

		m.List.SetSize(listWidth, listHeight)

		// h, v := ui.OolongStyle.GetFrameSize()
		// m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func (m Model) viewSearch() string {
	// TODO:
	return ui.OolongStyle.Render(m.List.View())
}
