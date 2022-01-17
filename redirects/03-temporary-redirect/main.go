package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at foo:", r.Method, "\n\n")
}

func bar(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Your request method at bar:", r.Method, "\n\n")

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}

func barred(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Your request method at barrred:", r.Method, "\n\n")

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
