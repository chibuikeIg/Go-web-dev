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

	CookieKeys := []string{"General", "MyCookie"}

	readAllCookie(w, r, CookieKeys)

}

func abundant(w http.ResponseWriter, r *http.Request) {

	cookie1 := Cookie{Name: "General", Value: "General Value"}
	cookie2 := Cookie{Name: "MyCookie", Value: "My Cookie Value"}

	setMultiCookie(w, []Cookie{cookie1, cookie2})

	fmt.Fprintln(w, "Your Cookies has been set")
}

func setMultiCookie(w http.ResponseWriter, cookies []Cookie) {

	for _, element := range cookies {

		http.SetCookie(w, &http.Cookie{
			Name:  element.Name,
			Value: element.Value,
		})
	}

	return
}

type Cookie struct {
	Name  string
	Value string
}

func readAllCookie(w http.ResponseWriter, r *http.Request, CookieKeys []string) {

	for _, element := range CookieKeys {

		c, err := r.Cookie(element)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Fprintln(w, "YOUR COOKIE:", c)
		}
	}
}
