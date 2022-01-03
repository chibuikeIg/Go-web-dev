package main

import (
	"log"
	"net/http"
	"text/template"
)

type Hotdog int

func (m Hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.html", r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var d Hotdog

	http.ListenAndServe(":8080", d)
}
