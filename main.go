package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const (
	PORT string = ":4000"
)

func handleError(err error) {
	if err != nil {
		result := fmt.Errorf(err.Error())
		fmt.Print(result)
	}
}

func RenderTemplate(response http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(response, nil)

	handleError(err)
}

func Home(response http.ResponseWriter, request *http.Request) {
	RenderTemplate(response, "home.html")

	// n, error := fmt.Fprintf(response, "this is the home page")
	// handleError(error)
	// fmt.Println(n)
}

func About(response http.ResponseWriter, request *http.Request) {
	RenderTemplate(response, "about.html")

	// n, error := fmt.Fprintf(response, "Welcome to the about page.")
	// handleError(error)
	// fmt.Println(n)

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Print("Listening on http://localhost:4000... ")
	http.ListenAndServe(PORT, nil)
}
