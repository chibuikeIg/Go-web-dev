package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/lovebirds", love)
	http.HandleFunc("/love.jpg", lovePic)

	http.ListenAndServe(":8080", nil)
}

func love(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/love.jpg">`)
}

func lovePic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("love.jpg")

	if err != nil {
		http.Error(w, "File not found", 404)
	}

	defer f.Close()

	fi, err := f.Stat()

	if err != nil {
		http.Error(w, "File not found", 404)
	}

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}
