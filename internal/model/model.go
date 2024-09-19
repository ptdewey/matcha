package model

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/data"
)

const (
	LANDING = iota
	CREATE
	SEARCH
	BROWSE
	EDIT
)

// TODO: add new fields to allow multi-stage application
type Model struct {
	List         list.Model
	NoteSources  []string
	Mode         int
	ModeChosen   bool
	SelectedNote data.Note
	cursor       int
	width        int
	height       int
	err          error
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		func() tea.Msg {
			return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}}
		},
	)
}
