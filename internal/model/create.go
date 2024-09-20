package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ptdewey/oolong/internal/ui"
)

const (
	name = iota
	dir
	ext
)

// TODO: see credit-card example for text multi-input (with defaults)
// - options for filename, dir, filetype?
// - default filename is parent dirname + date
// - default dir is current dir
// - default filetype is .md
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
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.Inputs {
			m.Inputs[i].Blur()
		}
		m.Inputs[m.focused].Focus()

	// We handle errors just like any other message
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
	return fmt.Sprintf(
		`Create Note:
%s
%s

%s
%s

%s
%s

%s
`,
		ui.InputStyle.Width(30).Render("Filename"),
		m.Inputs[name].View(),
		ui.InputStyle.Width(30).Render("Directory"),
		m.Inputs[dir].View(),
		ui.InputStyle.Width(15).Render("Extension"),
		m.Inputs[ext].View(),
		ui.ContinueStyle.Render("Continue ->"),
	) + "\n"
}

func initTextInput() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 3)
	inputs[name] = textinput.New()
	inputs[name].Placeholder = "TODO: populate with date-dir"
	inputs[name].Focus()
	inputs[name].CharLimit = 0
	inputs[name].Width = 30
	inputs[name].Prompt = ""

	inputs[dir] = textinput.New()
	inputs[dir].Placeholder = "TODO: populate with current directory"
	inputs[dir].CharLimit = 0
	inputs[dir].Width = 30
	inputs[dir].Prompt = ""

	inputs[ext] = textinput.New()
	inputs[ext].Placeholder = ".md"
	inputs[ext].CharLimit = 0
	inputs[ext].Width = 15
	inputs[ext].Prompt = ""

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
