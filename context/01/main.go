package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/foo", Index)

	http.Handle("/", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func Index(w http.ResponseWriter, req *http.Request) {

	context := req.Context()

	fmt.Fprintln(w, context)

}
