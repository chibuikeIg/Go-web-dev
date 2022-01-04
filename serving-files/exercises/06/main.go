package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {

	// parse all template files in the template folder using ParseGlob

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	// create routes to render pages

	http.HandleFunc("/", index)

	http.HandleFunc("/about", about)

	http.HandleFunc("/contact", contact)

	// This route serves response for both GET and POST requests

	http.HandleFunc("/apply", apply)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {

	// check the request methods on this route

	if req.Method == http.MethodPost {
		exec(w, "applyProcess.gohtml")
	} else if req.Method == http.MethodGet {
		exec(w, "apply.gohtml")
	}

}

func exec(w http.ResponseWriter, template string) {
	err := tpl.ExecuteTemplate(w, template, nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
