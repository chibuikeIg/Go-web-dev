package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", love)

	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}

func love(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/resources/img/love.jpg">`)
}
