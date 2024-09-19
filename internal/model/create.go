package model

import tea "github.com/charmbracelet/bubbletea"

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
	_ = msg
	// TODO:
	return m, nil
}

func (m Model) viewCreate() string {
	return "TODO: create view"
}
