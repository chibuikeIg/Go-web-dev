package main

import (
	"io"
	"net/http"
)

type HotDog int

func (m HotDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "This is the dog path")
	case "/cat":
		io.WriteString(w, "This is the cat path")
	default:
		io.WriteString(w, "404 not found page")
	}
}

func main() {

	var d HotDog

	http.ListenAndServe(":8080", d)
}
