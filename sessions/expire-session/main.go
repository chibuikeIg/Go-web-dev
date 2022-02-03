package main

import (
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string
	Firstname string
	Password  []byte
	Lastname  string
	Type      string
}

type Session struct {
	un           string
	lastActivity time.Time
}

var DBUsers = map[string]User{}
var DBSessions = map[string]Session{}
var SessionsStart time.Time
var sessionLength int = 30

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
	SessionsStart = time.Now()
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	u := GetUser(w, r)
	ShowSessions()

	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := GetUser(w, r)

	if !AlreadyLoggedIn(r) {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if u.Type != "admin" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ShowSessions()

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
		t := r.FormValue("type")

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

		c.MaxAge = sessionLength

		http.SetCookie(w, c)

		DBSessions[c.Value] = Session{un, time.Now()}

		// store users

		bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		DBUsers[un] = User{un, fname, bs, lname, t}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)

}

func login(w http.ResponseWriter, r *http.Request) {

	if AlreadyLoggedIn(r) {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		un := r.FormValue("username")
		pwd := r.FormValue("password")

		u, ok := DBUsers[un]

		if !ok {

			http.Error(w, "Username and/or password is incorrect", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pwd))

		if err != nil {

			http.Error(w, "Username and/or password is incorrect", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()

		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)

		DBSessions[c.Value] = Session{un, time.Now()}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)

}

func logout(w http.ResponseWriter, r *http.Request) {

	if !AlreadyLoggedIn(r) {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// get cookie
	c, _ := r.Cookie("session")

	// delete session

	delete(DBSessions, c.Value)

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, c)

	if time.Now().Sub(SessionsStart) > (time.Second * 30) {
		go CleanSessions()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

	return
}
