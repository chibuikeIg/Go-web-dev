package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "username:password@tcp(endpoint:3306)/database?charset=utf8")

	check(err)

	defer db.Close()

	err = db.Ping()

	check(err)

	http.HandleFunc("/", Index)
	http.HandleFunc("/names", Names)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	_, err = io.WriteString(w, "Successfully connected")

	check(err)
}

func Names(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(`SELECT name FROM users`)

	check(err)

	var s, name string

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&name)
		check(err)

		s += name + "\n"

	}

	fmt.Fprintln(w, s)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
