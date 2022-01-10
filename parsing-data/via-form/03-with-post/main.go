package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type Person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	f := r.FormValue("fname")
	l := r.FormValue("lname")
	s := r.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", Person{f, l, s})

	if err != nil {

		http.Error(w, err.Error(), 500)

		log.Fatalln(err)
	}
}
