package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ptdewey/oolong/internal/config"
	"github.com/ptdewey/oolong/internal/data"
	"github.com/ptdewey/oolong/internal/model"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF00")) // Green text for the title

	itemStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(lipgloss.Color("#FFFFFF")) // White text for items

	selectedItemStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#7D56F4")). // Purple background for selected item
				Foreground(lipgloss.Color("#FFFFFF"))  // White text for selected item
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 } // Each item takes 1 row
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	note, ok := item.(data.Note) // Cast to your data.Note type
	if !ok {
		return
	}

	// Determine if this item is selected
	isSelected := index == m.Index()

	var title string
	if isSelected {
		title = selectedItemStyle.Render(note.Title()) // Use selected item style
	} else {
		title = itemStyle.Render(note.Title()) // Use regular item style
	}

	fmt.Fprintf(w, "%s", title)
}

func main() {
	cfg := config.ParseConfig()

	items := data.GetItems(cfg.NoteSources)

	// TODO: define entrypoint as a checkbox selector if an argument is not passed
	// - choose search/create note
	// - passing arg will choose one initially

	m := model.Model{
		List: list.New(items, list.NewDefaultDelegate(), 0, 0),
		// List:        list.New(items, itemDelegate{}, 50, 30),
		NoteSources: cfg.NoteSources,
	}
	m.List.Title = "Notes"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
