package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ptdewey/oolong/internal/ui"
)

func (m Model) updateLanding(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h", "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right", "l", "down", "j":
			if m.cursor < 1 {
				m.cursor++
			}

		case "enter":
			m.ModeChosen = true
			m.Mode = m.cursor + 1

			// set size based on terminal to fix broken list size
			m.List.SetSize(m.width, m.height)

			return m, func() tea.Msg {
				return tea.WindowSizeMsg{Width: m.width, Height: m.height}
			}
		}
	}

	return m, nil
}

func (m Model) viewLanding() string {
	if m.ModeChosen {
		return ""
	}

	var createButton, searchButton string

	if m.cursor == 0 {
		createButton = ui.ActiveButton.Render("Create Note")
		searchButton = ui.InactiveButton.Render("Search Notes")
	} else {
		createButton = ui.InactiveButton.Render("Create Note")
		searchButton = ui.ActiveButton.Render("Search Notes")
	}

	gap := lipgloss.NewStyle().Width(5).Render(" ")

	buttons := lipgloss.JoinHorizontal(lipgloss.Top, createButton, gap, searchButton)

	ui := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, buttons)

	return ui
}
