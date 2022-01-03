package main

import (
	"io"
	"net/http"
)

type HotDog int

func (m HotDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "this is a hot dog path")

}

type HotCat int

func (m HotCat) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "This is a hot cat path")

}

func main() {

	var d HotDog
	var c HotCat
	mux := http.NewServeMux()

	mux.Handle("/hot-dog", d)
	mux.Handle("/hot-cat", c)

	http.ListenAndServe(":8080", mux)
}
