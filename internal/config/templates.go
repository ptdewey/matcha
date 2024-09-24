package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type NoteTemplate struct {
	Name string
	Ext  string
}

func ReadTemplates(cfg Config) []NoteTemplate {
	out := []NoteTemplate{}

	if err := filepath.WalkDir(cfg.TemplateDir, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		out = append(out, NoteTemplate{
			Name: path,
			Ext:  filepath.Ext(entry.Name()),
		})

		return nil
	}); err != nil {
		fmt.Printf("Error walking TemplateDir %s: %v\n", cfg.TemplateDir, err)
	}

	return out
}
