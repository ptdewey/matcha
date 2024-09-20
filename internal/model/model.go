package model

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/config"
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
	// model mode
	Mode int

	// landing page fields
	ModeChosen bool
	cursor     int
	width      int
	height     int

	// browse/search mode fields
	List         list.Model
	NoteSources  []string
	SelectedNote data.Note

	// inputs for create mode
	Inputs  []textinput.Model
	focused int

	err error
}

func (m Model) Init() tea.Cmd {
	// return nil
	return textinput.Blink
}

func InitialModel() Model {
	cfg := config.ParseConfig()

	// TODO: possibly move this somewhere else to avoid loading notes during quick-launch/create
	items := data.GetItems(cfg.NoteSources)
	inputs := initTextInput()

	return Model{
		List:        list.New(items, list.NewDefaultDelegate(), 0, 0),
		NoteSources: cfg.NoteSources,
		Inputs:      inputs,
		focused:     0,
		err:         nil,
	}
}
