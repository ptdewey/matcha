package main

import (
	"fmt"
	"os"

	"github.com/ptdewey/oolong/internal/config"
	"github.com/ptdewey/oolong/internal/data"
	"github.com/ptdewey/oolong/internal/model"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg := config.ParseConfig()

	// TODO: possibly move this somewhere else to avoid loading notes during quick-launch/create
	items := data.GetItems(cfg.NoteSources)

	// TODO: define entrypoint as a checkbox selector if an argument is not passed
	// - choose search/create note
	// - passing arg will choose one initially
	// other options?
	// - option for quickopening a note (pass current dir into filter with date?)
	//    - this will require some config file setup

	m := model.Model{
		List:        list.New(items, list.NewDefaultDelegate(), 0, 0),
		NoteSources: cfg.NoteSources,
	}
	m.List.Title = "Notes"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
