package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "mycookie",
		Value: "my cookie value",
	})

	fmt.Fprintln(w, "Your Cookie has been Set, Check Your Browser")

	fmt.Fprintln(w, "in chrome go to dev tools/ application / cookies")
}

func read(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("mycookie")

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	fmt.Fprintln(w, "Your Cookie:", c)
}
