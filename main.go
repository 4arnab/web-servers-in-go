package main

import (
	"fmt"
	"net/http"
)

const (
	PORT string = ":4000"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Print("Listening on http://localhost:4000... ")
	http.ListenAndServe(PORT, nil)
}
