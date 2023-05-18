package handlers

import (
	"net/http"

	"github.com/4arnab/web-servers-in-go/pkg/render"
)

func Home(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, "home.page.tmpl")

	// n, error := fmt.Fprintf(response, "this is the home page")
	// handleError(error)
	// fmt.Println(n)
}

func About(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, "about.page.tmpl")

	// n, error := fmt.Fprintf(response, "Welcome to the about page.")
	// handleError(error)
	// fmt.Println(n)

}
