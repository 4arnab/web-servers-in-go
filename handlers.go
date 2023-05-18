package main

import (
	"fmt"
	"net/http"
)

func handleError(err error) {
	if err != nil {
		result := fmt.Errorf(err.Error())
		fmt.Print(result)
	}
}

func Home(response http.ResponseWriter, request *http.Request) {
	RenderTemplate(response, "home.page.tmpl")

	// n, error := fmt.Fprintf(response, "this is the home page")
	// handleError(error)
	// fmt.Println(n)
}

func About(response http.ResponseWriter, request *http.Request) {
	RenderTemplate(response, "about.page.tmpl")

	// n, error := fmt.Fprintf(response, "Welcome to the about page.")
	// handleError(error)
	// fmt.Println(n)

}
