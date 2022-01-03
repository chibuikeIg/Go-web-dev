package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/lovebirds", ld)
	http.HandleFunc("/download.jfif", dwn)
	http.HandleFunc("/img/love.jpg", l)
	http.ListenAndServe(":8080", nil)

}

func ld(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<img src="download.jfif">
		<img src="/img/love.jpg">
	`)

}

func dwn(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("download.jfif")

	if err != nil {
		http.Error(w, "File not Found", 404)
	}

	defer f.Close()

	io.Copy(w, f)
}

func l(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("img/love.jpg")

	if err != nil {
		http.Error(w, "File not found", 404)
	}

	defer f.Close()

	io.Copy(w, f)
}
