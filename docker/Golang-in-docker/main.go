package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", Index)

	http.ListenAndServe(":80", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello world from docker container")
}
