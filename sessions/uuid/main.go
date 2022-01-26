package main

import (
	"fmt"
	"net/http"

	exPackage "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session")

	if err != nil {

		id, _ := exPackage.NewV4()

		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)

	}

	fmt.Println(cookie)

}
