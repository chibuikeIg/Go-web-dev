package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string
	Firstname string
	Password  []byte
	Lastname  string
}

var DBUsers = map[string]User{}
var DBSessions = map[string]string{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	DBUsers["test@user.com"] = User{"test@user.com", "test", bs, "user"}
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

		bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		DBUsers[un] = User{un, fname, bs, lname}

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

		DBSessions[c.Value] = un

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

	http.Redirect(w, r, "/", http.StatusSeeOther)

	return
}
