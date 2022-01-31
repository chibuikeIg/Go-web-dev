package main

import (
	"net/http"
	"text/template"
)

type User struct {
	Username  string
	Firstname string
	Password  string
	Lastname  string
}

var DBUsers = map[string]User{}
var DBSessions = map[string]string{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	u := getUser(w, r)

	tpl.ExecuteTemplate(w, "index.html", u)
}
