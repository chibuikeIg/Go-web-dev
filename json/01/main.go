package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Fname       string
	Lname       string
	Personality []string
}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/marshal", Marshal)
	http.HandleFunc("/encode", Encde)

	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	s := `<!DOCTYPE html>
		 	<html>
				<head>
					<title>Html Document</title>
				</head>
				<body>
					<h5>Hello World</h5>
				</body>
			</html>`

	w.Write([]byte(s))

}

func Marshal(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	person := Person{
		"James",
		"Parker",
		[]string{"Suit", "Gun", "Werry sense of humor"},
	}

	json, err := json.Marshal(person)

	if err != nil {
		log.Fatalln(err)
	}

	w.Write([]byte(json))

}

func Encde(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	person := Person{
		"James",
		"Parker",
		[]string{"Suit", "Gun", "Werry sense of humor"},
	}

	err := json.NewEncoder(w).Encode(person)

	if err != nil {
		log.Fatalln(err)
	}

}
