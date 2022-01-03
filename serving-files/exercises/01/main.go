package main

import (
	"io"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/love/", love)
	http.HandleFunc("/love.jpg", loveImg)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, "foo ran")

}

func love(w http.ResponseWriter, r *http.Request) {

	var tpl *template.Template

	tpl = template.Must(template.ParseFiles("love.gohtml"))

	tpl.ExecuteTemplate(w, "love.gohtml", nil)

}

func loveImg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "love.jpg")
}
