package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "Go look at your terminal")
}
