package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := page()

	http.Handle("/", templ.Handler(component))

	http.ListenAndServe(":8080", nil)
}
