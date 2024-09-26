package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
			// adjust this number if buttons are added/removed
			if m.cursor < 2 {
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

	var quickButton, createButton, searchButton string

	if m.cursor == 0 {
		quickButton = ActiveButton.Render("Quick Note")
		createButton = InactiveButton.Render("Create Note")
		searchButton = InactiveButton.Render("Search Notes")
	} else if m.cursor == 1 {
		quickButton = InactiveButton.Render("Quick Note")
		createButton = ActiveButton.Render("Create Note")
		searchButton = InactiveButton.Render("Search Notes")
	} else {
		quickButton = InactiveButton.Render("Quick Note")
		createButton = InactiveButton.Render("Create Note")
		searchButton = ActiveButton.Render("Search Notes")
	}

	gap := lipgloss.NewStyle().Width(5).Render(" ")

	buttons := lipgloss.JoinHorizontal(lipgloss.Top, quickButton, gap, createButton, gap, searchButton)

	ui := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, buttons)

	return ui
}
