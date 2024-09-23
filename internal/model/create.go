package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ptdewey/oolong/internal/editor"
	"github.com/ptdewey/oolong/internal/notes"
)

const (
	name = iota
	dir
	ext
)

// TODO: inherit from config? (also add selector to change between defaults and templates)
// - options for filename, dir, filetype?
// Include file content options?
// - My note template with parent dir + grandparent dir + date
//   - this should probably be defined in a config file or template of some kind
// (text multi-input) https://github.com/charmbracelet/bubbletea/blob/main/examples/credit-card-form/main.go

func (m Model) updateCreate(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.Inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.Inputs)-1 {
				notePath := m.checkJoinInputs()

				if err := notes.CreateNote(notePath); err != nil {
					fmt.Println("Error creating note:", err)
					return m, nil
				}

				return m, editor.OpenEditor(notePath)
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP, tea.KeyUp:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN, tea.KeyDown:
			m.nextInput()
		}
		for i := range m.Inputs {
			m.Inputs[i].Blur()
		}
		m.Inputs[m.focused].Focus()
	case error:
		m.err = msg
		return m, nil
	}

	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) viewCreate() string {
	// TODO: make this look way better
	// - maybe use lipgloss.table?

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
		fmt.Sprintf(
			`%s

%s
%s

%s
%s

%s
%s

%s
`,
			HeaderStyle.Render("Create Note:"),
			InputHeaderStyle.Render("Filename"),
			InputStyle.Render(m.Inputs[name].View()),
			InputHeaderStyle.Render("Directory"),
			InputStyle.Render(m.Inputs[dir].View()),
			InputHeaderStyle.Render("Extension"),
			InputStyle.Render(m.Inputs[ext].View()),
			ContinueStyle.Render("Continue ->"),
		))
}

func initTextInput() []textinput.Model {
	defaultName, defaultRelPath := notes.GenFileNamePath()
	// REFACTOR: change defaultpath to be relative to notes directory, not system
	defaultPath := filepath.Join(cfg.NoteSources[0], defaultRelPath)

	var inputs []textinput.Model = make([]textinput.Model, 3)

	// filename
	inputs[name] = textinput.New()
	inputs[name].Placeholder = defaultName
	inputs[name].Focus()
	inputs[name].CharLimit = 0
	inputs[name].Width = 20

	// directory
	inputs[dir] = textinput.New()
	inputs[dir].Placeholder = defaultPath
	inputs[dir].Width = 40

	// file extension
	inputs[ext] = textinput.New()
	inputs[ext].Placeholder = cfg.DefaultExt
	inputs[ext].Width = 10

	// TODO: field for templates to use? (might need to be a selector of some kind to allow multiple)

	return inputs
}

func (m *Model) nextInput() {
	m.focused = (m.focused + 1) % len(m.Inputs)
}

func (m *Model) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.Inputs) - 1
	}
}

func (m *Model) checkJoinInputs() string {
	// check name, set to default if empty
	n := m.Inputs[name]
	if n.Value() == "" {
		n.SetValue(n.Placeholder)
	}

	// check directory, set to default if empty
	// FIX: add handling of ~/ in input
	d := m.Inputs[dir]
	if d.Value() == "" {
		d.SetValue(d.Placeholder)
	}
	if _, err := os.Stat(d.Value()); os.IsNotExist(err) {
		// create directories that don't exist
		if err = os.MkdirAll(d.Value(), os.ModePerm); err != nil {
			// TODO: switch print to logging?
			fmt.Printf("Error creating directory(ies) %s: %v", d.Value(), err)
			d.SetValue(d.Placeholder)
		}
	}

	// check file extenson, set to default if empty
	e := m.Inputs[ext]
	if e.Value() == "" {
		e.SetValue(e.Placeholder)
	} else if !strings.HasPrefix(e.Value(), ".") {
		e.SetValue("." + e.Value())
	}

	return filepath.Join(d.Value(), n.Value()+e.Value())
}
