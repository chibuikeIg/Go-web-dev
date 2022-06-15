package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/instance", Instance)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":80", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello world from AWS")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func Instance(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")

	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, resp.ContentLength)

	resp.Body.Read(bs)

	resp.Body.Close()

	io.WriteString(w, string(bs))

}
