package render

import (
	"embed"
	"html/template"
	"net/http"
)

func RenderTemplate(content embed.FS, w http.ResponseWriter, r *http.Request, html string) error {
	templateContent, err := content.ReadFile("dist/" + html)
	if err != nil {
		return err
	}

	parseTemplate, err := template.New("").Parse(string(templateContent))
	if err != nil {
		return err
	}

	err = parseTemplate.Execute(w, nil)
	if err != nil {
		return err
	}

	return nil
}

