package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	var s string

	if r.Method == http.MethodPost {

		// catch file from request body
		f, h, err := r.FormFile("q")

		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// close file
		defer f.Close()

		// output file information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerror:", err)

		// read the file

		bs, err := ioutil.ReadAll(f)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)

	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `<form method="post" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form><br>`+s)
}
