package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, error := fmt.Fprintf(w, "hello, from the server side")

		if error != nil {
			fmt.Println(error.Error())
			panic(error)
		}

		fmt.Println(n)

	})

	fmt.Print("Listening on http://localhost:4000... ")
	http.ListenAndServe(":4000", nil)
}
