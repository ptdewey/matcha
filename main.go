package main

import (
	"fmt"
	"os"

	"github.com/ptdewey/notes-manager/internal/config"
	"github.com/ptdewey/notes-manager/internal/data"
	"github.com/ptdewey/notes-manager/internal/model"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg := config.ParseConfig()

	items := data.GetItems(cfg.NoteSources)

	// TODO: define entrypoint as a checkbox selector if an argument is not passed
	// - choose search/create note
	// - passing arg will choose one initially

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
