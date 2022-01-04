package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	// route to handle resources

	http.Handle("/resources", http.StripPrefix("/resources", http.FileServer(http.Dir("./public"))))

	// route to access responses
	http.HandleFunc("/love", love)

	http.ListenAndServe(":8080", nil)

}

// initialize template library at a global scope

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

// love func of type responsewriter and request
func love(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)

	if err != nil {
		log.Fatal("Couldn't pass files")
	}
}
