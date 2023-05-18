package main

import (
	"html/template"
	"net/http"
)

func RenderTemplate(response http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(response, nil)

	handleError(err)
}
