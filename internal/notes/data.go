package notes

import (
	"fmt"
	"path/filepath"
	"time"
)

// TODO: make template data more customizable

type TemplateData struct {
	Path      string
	Dir       string
	ParentDir string
	Date      string
}

func templateData(notePath string) TemplateData {
	out := TemplateData{}

	out.Path = notePath

	dir := filepath.Dir(notePath)
	out.Dir = filepath.Base(dir)
	out.ParentDir = filepath.Base(filepath.Dir(dir))

	t := time.Now()
	out.Date = fmt.Sprintf("%s %s, %s", t.Format("January"), t.Format("2"), t.Format("2006"))

	return out
}
