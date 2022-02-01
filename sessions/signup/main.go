package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
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
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	u := GetUser(w, r)

	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := GetUser(w, r)

	if !AlreadyLoggedIn(r) {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signup(w http.ResponseWriter, r *http.Request) {

	if AlreadyLoggedIn(r) {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		un := r.FormValue("username")
		pwd := r.FormValue("password")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")

		if _, ok := DBUsers[un]; ok {

			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()

		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)

		DBSessions[c.Value] = un

		// store users

		DBUsers[un] = User{un, fname, pwd, lname}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)

}
