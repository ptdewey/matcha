package notes

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type CreateFinishedMsg struct {
	Err error
}

func CreateNote(notePath string, templatePath string) error {
	info, err := os.Stat(notePath)
	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(notePath)
		if err != nil {
			return err
		}
		defer f.Close()

		// if template path is specifed, populate generated note with template contents
		if templatePath != "" {
			if err := populateTemplate(f, templatePath); err != nil {
				return err
			}

		}

		return nil

	} else if err != nil {
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("Error: %s is a directory.", notePath)
	}

	return nil
}

func populateTemplate(f *os.File, templatePath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	if err = tmpl.Execute(f, templateData(f.Name())); err != nil {
		return err
	}

	return nil
}

func GenFileNamePath() (string, string) {
	// TODO: make this more customizable with config
	t := time.Now()
	date := t.Format("01-02-06")

	wd, err := os.Getwd()
	if err != nil {
		return "", ""
	}

	cwd := filepath.Base(wd)
	name := fmt.Sprintf("%s-%s", cwd, date)

	splitPath := strings.Split(wd, string(os.PathSeparator))
	if len(splitPath) < 3 {
		if len(splitPath) < 2 {
			return name, ""
		}
		return name, splitPath[0]
	}
	path := strings.Join(splitPath[len(splitPath)-3:len(splitPath)-1], string(os.PathSeparator))

	return name, path
}
