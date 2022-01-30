package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Username  string
	FirstName string
	LastName  string
}

var tpl *template.Template

var DBUsers = map[string]User{}
var DBSessions = map[string]string{}

func init() {

	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/bar", bar)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")

	if err != nil {

		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)

	}

	var u User

	if un, ok := DBSessions[c.Value]; ok {
		u = DBUsers[un]
	}

	if r.Method == http.MethodPost {

		un := r.FormValue("username")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")

		u = User{un, fname, lname}

		DBSessions[c.Value] = un
		DBUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")

	if err != nil {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	un, ok := DBSessions[c.Value]

	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := DBUsers[un]

	tpl.ExecuteTemplate(w, "bar.html", u)

}
