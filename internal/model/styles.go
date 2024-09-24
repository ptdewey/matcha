package model

import "github.com/charmbracelet/lipgloss"

// (centering with lipgloss) https://github.com/charmbracelet/bubbletea/discussions/818

var (
	ButtonStyle = lipgloss.NewStyle().Padding(0, 2).Border(lipgloss.RoundedBorder()).Align(lipgloss.Center)
	// ActiveButton   = ButtonStyle.Foreground(lipgloss.Color("#FFFFFF")).Background(lipgloss.Color("#5A56E0"))
	ActiveButton = ButtonStyle.
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("2")).
			Foreground(lipgloss.Color("default")).
			Background(lipgloss.Color("2"))
	InactiveButton = ButtonStyle.Foreground(lipgloss.Color("default"))

	HeaderStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color("2")).
			Foreground(lipgloss.Color("default"))

	InputHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("4"))
		// Foreground(lipgloss.Color("32"))

	InputStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("3"))
		// BorderForeground(lipgloss.Color("36"))

	ContinueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
)

var MatchaStyle = lipgloss.NewStyle().Margin(1, 2).
	Foreground(lipgloss.Color("default"))

// TODO: handle all the style stuff better
type Styles struct {
	Foreground  lipgloss.Color
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.Foreground = lipgloss.Color("default")
	s.BorderColor = lipgloss.Color("36")

	return s
}
