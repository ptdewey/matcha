package model

import "github.com/ptdewey/oolong/internal/ui"

func (m Model) View() string {
	return ui.Style.Render(m.List.View())
}
