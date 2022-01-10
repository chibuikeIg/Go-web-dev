package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	v := r.FormValue("q")

	io.WriteString(w, `<form method="get">
	<input type="text" name="q">
	<input type="submit">
	</form><br>`+v)
}
