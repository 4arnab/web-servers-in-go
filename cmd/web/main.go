package main

import (
	"fmt"
	"net/http"

	"github.com/4arnab/web-servers-in-go/pkg/handlers"
)

const (
	PORT string = ":4000"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Print("Listening on http://localhost:4000... ")
	http.ListenAndServe(PORT, nil)
}
