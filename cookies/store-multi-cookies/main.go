package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundant", abundant)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "mycookie",
		Value: "My Cookie Value",
	})

	fmt.Fprintln(w, "Your Cookie has been set")
}

func read(w http.ResponseWriter, r *http.Request) {

	c1, err := r.Cookie("mycookie")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}

	c2, err := r.Cookie("general")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", c2)
	}

	c3, err := r.Cookie("secret")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
	}
}

func abundant(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "general's cookie",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "secret",
		Value: "secret cookie",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "mycookie",
		Value: "my cookie",
	})

	fmt.Fprintln(w, "Your Cookies has been set")
}
