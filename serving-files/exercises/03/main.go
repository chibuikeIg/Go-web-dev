package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/pics/", fs)
	http.HandleFunc("/dog/", dog)

	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)

	if err != nil {
		log.Fatal("Couldn't pass templates")
	}
}
