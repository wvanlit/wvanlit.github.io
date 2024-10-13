package templates

import (
	"html/template"
	"site-generator/filesystem"
	"strings"
)

func ParseTemplates(files []filesystem.File) ([]template.Template, error) {
	componentFiles := make([]string, 0)
	templateFiles := make([]string, 0)

	for _, f := range files {
		if strings.Contains(f.FilePath, "components") {
			componentFiles = append(componentFiles, f.FilePath)
		} else {
			templateFiles = append(templateFiles, f.FilePath)
		}
	}

	templates := make([]template.Template, len(templateFiles))

	for i, t := range templateFiles {
		template, err := template.ParseFiles(append([]string{t}, componentFiles...)...)
		if err != nil {
			return nil, err
		}

		templates[i] = *template
	}

	return templates, nil
}
