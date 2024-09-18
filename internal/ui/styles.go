package ui

import "github.com/charmbracelet/lipgloss"

// (centering with lipgloss) https://github.com/charmbracelet/bubbletea/discussions/818

var Style = lipgloss.NewStyle().Margin(1, 2).
	Foreground(lipgloss.Color("default"))
