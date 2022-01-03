package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/lovebirds", love)
	http.HandleFunc("/img/love.jpg", lovePic)

	http.ListenAndServe(":8080", nil)
}

func love(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/img/love.jpg">`)
}

func lovePic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "img/love.jpg")
}
