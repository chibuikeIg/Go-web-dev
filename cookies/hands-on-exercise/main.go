package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("visits")

	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "visits",
			Value: "0",
		}
	}

	count, _ := strconv.Atoi(c.Value)
	count++
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)

	io.WriteString(w, c.Value)

}
