package lib

import (
	"fmt"
	"html/template"
	"io"
	"path"
	"path/filepath"
	"runtime"
)

// TemplateManager .
type TemplateManager struct {
}

func (tm TemplateManager) loadTemplate(templateName string) (*template.Template, error) {
	var err error
	t := template.New(templateName)
	_, filename, _, _ := runtime.Caller(1)
	cwd := path.Dir(filename)
	templateFileName := filepath.Join(cwd, fmt.Sprintf("./tmpl/%s", templateName))
	t, err = t.ParseFiles(templateFileName)
	return t, err
}

// Execute .
func (tm TemplateManager) Execute(templateName string, data interface{}, outputStream io.Writer) error {
	t, err := tm.loadTemplate(templateName)
	if err != nil {
		return err
	}
	return t.Execute(outputStream, data)
}
