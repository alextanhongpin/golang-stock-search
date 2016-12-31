package client

import (
	"net/http"
	"text/template"
)

var templates map[string]*template.Template

const (
	baseTemplate  string = "client/templates/base.html"
	indexTemplate string = "client/templates/index.html"
)

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles(indexTemplate, baseTemplate))
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {

	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template do not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
