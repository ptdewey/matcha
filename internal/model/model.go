package model

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/matcha/internal/config"
	"github.com/ptdewey/matcha/internal/data"
)

const (
	LANDING = iota
	QUICK
	CREATE
	SEARCH
	BROWSE
	EDIT
	QUIT
)

var cfg config.Config

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
	Inputs           []textinput.Model
	focused          int
	Templates        []config.NoteTemplate
	selectedTemplate int

	err error
}

func init() {
	cfg = config.ParseConfig()
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func InitialModel() Model {
	// TODO: possibly move this somewhere else to avoid loading notes during quick-launch/create
	items := data.GetItems(cfg.NoteSources, cfg.TemplateDir)
	inputs := initTextInput()
	templates := config.ReadTemplates(cfg)

	// TODO: populate templates from config?
	// - config dir option (assume text/template formatting)
	return Model{
		List:             list.New(items, list.NewDefaultDelegate(), 0, 0),
		NoteSources:      cfg.NoteSources,
		Inputs:           inputs,
		focused:          0,
		err:              nil,
		Templates:        templates,
		selectedTemplate: -1,
	}
}
