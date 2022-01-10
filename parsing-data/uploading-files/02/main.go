package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	var s string

	if r.Method == http.MethodPost {

		f, h, err := r.FormFile("q")

		handleError(err, w)

		defer f.Close()

		bs, err := ioutil.ReadAll(f)

		handleError(err, w)

		s = string(bs)

		newFilePath, err := os.Create(filepath.Join("./files/", h.Filename))

		handleError(err, w)

		_, err = newFilePath.Write(bs)

		handleError(err, w)

	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	tpl.ExecuteTemplate(w, "index.gohtml", s)
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
