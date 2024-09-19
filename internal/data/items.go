package data

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/charmbracelet/bubbles/list"
)

// TODO: get note extensions from config file? (or use these)
var noteExts = []string{".md", ".mdx", ".tex", ".txt", ".qmd", ".rmd", ".Rmd"}

type Note struct {
	title   string
	desc    string
	path    string
	relPath string
	ext     string
}

func (i Note) Title() string {
	return i.title
	// return i.relPath
}

// FIX: if filtering by relpath and using title for title, fix underlining (probably a delegate setting?)
func (i Note) Description() string {
	return i.relPath
	// return i.title
	// return i.desc
}

func (i Note) FilterValue() string {
	// return i.title
	return i.relPath
}

func (i Note) Path() string {
	return i.path
}

func GetItems(noteSources []string) []list.Item {
	var out []list.Item
	for _, src := range noteSources {
		items, err := getDirContents(src)
		if err != nil {
			fmt.Println(err)
			continue
		}
		out = append(out, items...)
	}

	return out
}

func getDirContents(dir string) ([]list.Item, error) {
	var entries []list.Item

	err := filepath.WalkDir(dir, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		info, err := entry.Info()
		if err != nil {
			return err
		}

		ext := filepath.Ext(entry.Name())
		if !slices.Contains(noteExts, ext) {
			return nil
		}

		entries = append(entries, Note{
			title:   entry.Name(),
			desc:    fmt.Sprintf("Mode: %s | Size: %d bytes\n", info.Mode(), info.Size()),
			relPath: relPath,
			path:    path,
			ext:     ext,
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	return entries, nil
}
