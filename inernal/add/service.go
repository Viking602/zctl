package add

import (
	"embed"
	"fmt"
	"io"
	"os"
	"text/template"
)

//go:embed handler.tmpl
//go:embed logic.tmpl

var templateData embed.FS

type service struct {
	ProjectName string
	LogicName   string
	ServiceName string
}

func (s *service) createFile(fileName string, tmpl string) error {
	f, err := os.Create(fmt.Sprintf(fileName))
	if err != nil {
		return err
	}
	if err := executeTemplate(tmpl, f, service{
		LogicName:   s.LogicName,
		ServiceName: s.ServiceName,
		ProjectName: s.ProjectName,
	}); err != nil {
		return err
	}
	return nil
}

func executeTemplate(tmplFile string, wr io.Writer, data any) error {
	tp, err := template.New(tmplFile).ParseFS(templateData, tmplFile)
	if err != nil {
		return err
	}
	if err := tp.ExecuteTemplate(wr, tmplFile, data); err != nil {
		return err
	}
	return nil
}
