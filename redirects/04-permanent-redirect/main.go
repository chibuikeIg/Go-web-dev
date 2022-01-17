package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at foo:", r.Method, "\n\n")
}

func bar(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Your request method at bar:", r.Method, "\n\n")

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)

}
