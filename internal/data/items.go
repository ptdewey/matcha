package data

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
)

type Note struct {
	title string
	desc  string
	path  string
}

func (i Note) Title() string {
	return i.title
}

func (i Note) Description() string {
	return i.desc
}

func (i Note) FilterValue() string {
	return i.path
}

func (i Note) Path() string {
	return i.path
}

func GetItems(noteSources []string) []list.Item {
	var out []list.Item
	for _, src := range noteSources {
		items, err := getDirContents(src)
		if err != nil {
			return nil
		}
		out = append(out, items...)
	}

	return out
}

func getDirContents(dir string) ([]list.Item, error) {
	var entries []list.Item

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		entries = append(entries, Note{
			title: d.Name(),
			desc:  fmt.Sprintf("Mode: %s | Size: %d bytes\n", info.Mode(), info.Size()),
			path:  relPath,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return entries, nil
}
