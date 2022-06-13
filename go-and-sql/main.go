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
	http.HandleFunc("/create", CreateDBTable)
	http.HandleFunc("/insert", Insert)

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

func CreateDBTable(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`CREATE TABLE customers (name VARCHAR(200))`)
	check(err)

	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Tables Created: ", n)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customers (name) VALUES ("John")`)
	check(err)

	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Records Created: ", n)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
