package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	h := http.NewServeMux()
	h.Handle("/", templ.Handler(home()))
	http.ListenAndServe("localhost:8000", h)
}
