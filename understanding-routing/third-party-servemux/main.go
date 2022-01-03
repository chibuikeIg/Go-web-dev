package main

import (
	"log"
	"net/http"
	"text/template"

	expkg "github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := expkg.New()

	mux.GET("/", Index)
	// mux.GET("/about", about)
	// mux.GET("/contact", contact)
	// mux.GET("/apply", apply)
	// mux.POST("/apply", applyProcess)
	// mux.GET("/user/:name", user)
	// mux.GET("/blog/:category/:article", blogRead)
	// mux.POST("/blog/:category/:article", blogWrite)

	http.ListenAndServe(":8080", mux)
}

func Index(w http.ResponseWriter, r *http.Request, _ expkg.Params) {

	err := tpl.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Panic(err)
	}

}
