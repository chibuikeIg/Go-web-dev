package main

import "net/http"

func main() {

	http.HandleFunc("/", Index)

	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte("This is an example server "))
}
