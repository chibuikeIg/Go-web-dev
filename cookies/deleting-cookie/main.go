package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<a href="/set">Set Cookie</a>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "mycookie",
		Value: "Cookie value",
	})

	fmt.Fprintln(w, `<a href="/read">Read Cookie</a>`)

}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mycookie")

	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
	}

	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><a href="/expire">Expire</a>`, c.Value)
}

func expire(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("mycookie")

	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
	}

	c.MaxAge = -1

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
