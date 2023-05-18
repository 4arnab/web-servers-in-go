package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleError(err error) {
	if err != nil {
		result := fmt.Errorf(err.Error())
		fmt.Print(result)
	}
}

func RenderTemplate(response http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(response, nil)

	HandleError(err)
}
