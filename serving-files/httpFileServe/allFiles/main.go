package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/lovebirds", love)

	http.ListenAndServe(":8080", nil)
}

func love(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="img/love.jpg">`)
}
