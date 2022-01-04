package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	// serve files in the public folder when request route is /public

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template/index.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)

	if err != nil {
		log.Fatal("Could not pass template")
	}
}
