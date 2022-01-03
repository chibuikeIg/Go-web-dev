package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog", d)

	http.ListenAndServe(":8080", nil)
}

func d(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="https://t3.ftcdn.net/jpg/04/21/59/34/360_F_421593489_b0VOTkrKSBfVNf30WMkvCZDmvI6xmRyD.jpg">
	`)
}
