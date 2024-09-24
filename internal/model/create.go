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
	"github.com/ptdewey/oolong/internal/utils"
)

const (
	NAME = iota
	DIR
	EXT
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

				ext := filepath.Ext(notePath)
				templateSource := ""
				for _, template := range m.Templates {
					fmt.Println(template.Ext, ext)
					if template.Ext == ext {
						templateSource = template.Name
						fmt.Println(templateSource)
						break
					}
				}

				if err := notes.CreateNote(notePath, templateSource); err != nil {
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
			InputStyle.Render(m.Inputs[NAME].View()),
			InputHeaderStyle.Render("Directory"),
			InputStyle.Render(m.Inputs[DIR].View()),
			InputHeaderStyle.Render("Extension"),
			InputStyle.Render(m.Inputs[EXT].View()),
			ContinueStyle.Render("Continue ->"),
		))
}

func initTextInput() []textinput.Model {
	defaultName, defaultRelPath := notes.GenFileNamePath()
	// REFACTOR: change defaultpath to be relative to notes directory, not system
	defaultPath := filepath.Join(cfg.NoteSources[0], defaultRelPath)

	var inputs []textinput.Model = make([]textinput.Model, 3)

	// filename
	inputs[NAME] = textinput.New()
	inputs[NAME].Placeholder = defaultName
	inputs[NAME].Focus()
	inputs[NAME].CharLimit = 0
	inputs[NAME].Width = 20

	// directory
	inputs[DIR] = textinput.New()
	inputs[DIR].Placeholder = defaultPath
	inputs[DIR].Width = 40

	// file extension
	inputs[EXT] = textinput.New()
	inputs[EXT].Placeholder = cfg.DefaultExt
	inputs[EXT].Width = 10

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
	n := m.Inputs[NAME]
	if n.Value() == "" {
		n.SetValue(n.Placeholder)
	}

	// check directory, set to default if empty
	d := m.Inputs[DIR]
	if d.Value() == "" {
		d.SetValue(d.Placeholder)
	} else {
		// if dir is valid, attempt to convert ~/ to /home/username if necessary
		path, err := utils.TildeToHome(d.Value())
		if err != nil {
			d.SetValue(d.Placeholder)
		} else {
			d.SetValue(path)
		}
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
	e := m.Inputs[EXT]
	if e.Value() == "" {
		e.SetValue(e.Placeholder)
	} else if !strings.HasPrefix(e.Value(), ".") {
		e.SetValue("." + e.Value())
	}

	return filepath.Join(d.Value(), n.Value()+e.Value())
}
