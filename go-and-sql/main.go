package main

import (
	"database/sql"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "admin:passwordaws@tcp(database-2.cajbdc97fqeg.us-east-1.rds.amazonaws.com:3306)/test01?charset=utf8")

	check(err)

	defer db.Close()

	err = db.Ping()

	check(err)

	http.HandleFunc("/", Index)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	_, err = io.WriteString(w, "Successfully connected")

	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}