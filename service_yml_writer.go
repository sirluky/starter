package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/cloud66/starter/common"
	"github.com/cloud66/starter/packs"
)

type ServiceYAMLWriter struct {
	TemplateDir     string
	ShouldOverwrite bool
}

func (w *ServiceYAMLWriter) write(analyzer packs.Analyzer) error {
	destFullPath := filepath.Join(analyzer.GetRootDir(), "service.yml")

	tmpl, err := template.ParseFiles(filepath.Join(w.TemplateDir, "service.yml.template"))
	if err != nil {
		return err
	}

	if _, err := os.Stat(destFullPath); !os.IsNotExist(err) && !w.ShouldOverwrite {
		return fmt.Errorf("service.yml exists and will not be overwritten unless the overwrite flag (-o) is set")
	}

	destFile, err := os.Create(destFullPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := destFile.Close(); err != nil {
			fmt.Printf("%s Cannot close file service.yml due to %s\n", common.MsgError, err.Error())
		}
	}()

	fmt.Println(common.MsgL1, "Writing service.yml...")
	err = tmpl.Execute(destFile, analyzer.GetContext())
	if err != nil {
		return err
	}

	return nil
}
