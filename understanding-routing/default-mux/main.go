package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	io.WriteString(w, "The path entered is "+path)
}
func main() {

	http.HandleFunc("/hot-dog", d)

	http.ListenAndServe(":8080", nil)

}
