package main

import (
	"embed"
	"fmt"
	"net/http"
	"text/template"
)

//go:embed views/*.html
var templateFS embed.FS

//go:embed static/*/*
var staticFS embed.FS

var templateCache map[string]*template.Template = map[string]*template.Template{}

func renderTemplate(w http.ResponseWriter, templateName string, data any) {
	var tmpl *template.Template
	if _, ok := templateCache[templateName]; !ok {
		tmpl = template.Must(template.New("layout").ParseFS(templateFS, "views/layout.html", fmt.Sprintf("views/%s.html", templateName)))
		templateCache[templateName] = tmpl
	} else {
		tmpl = templateCache[templateName]
	}
	tmpl.ExecuteTemplate(w, fmt.Sprintf("%s.html", templateName), data)
}
