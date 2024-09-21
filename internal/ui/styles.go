package ui

import "github.com/charmbracelet/lipgloss"

// (centering with lipgloss) https://github.com/charmbracelet/bubbletea/discussions/818

// CHANGE: inherit style from style.go
var (
	ButtonStyle    = lipgloss.NewStyle().Padding(0, 2).Border(lipgloss.RoundedBorder()).Align(lipgloss.Center)
	ActiveButton   = ButtonStyle.Foreground(lipgloss.Color("#FFFFFF")).Background(lipgloss.Color("#5A56E0"))
	InactiveButton = ButtonStyle.Foreground(lipgloss.Color("#888888"))
	InputStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#5A56E0"))
	ContinueStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
)

var OolongStyle = lipgloss.NewStyle().Margin(1, 2).
	Foreground(lipgloss.Color("default"))
