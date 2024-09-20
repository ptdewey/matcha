package main

import (
	"fmt"
	"os"

	"github.com/ptdewey/oolong/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// TODO: define entrypoint as a checkbox selector if an argument is not passed
	// - choose search/create note
	// - passing arg will choose one initially
	// other options?
	// - option for quickopening a note (pass current dir into filter with date?)
	//    - this will require some config file setup

	m := model.InitialModel()
	m.List.Title = "Notes"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
