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

	db, err = sql.Open("mysql", "user:password@tcp(endpoint:3306)/database?charset=utf8")

	http.HandleFunc("/", Index)
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/instance", Instance)
	http.HandleFunc("/users", GetUsers)
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

	io.WriteString(w, GetInstance())
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(`SELECT name FROM users`)
	check(err)
	defer rows.Close()

	s := GetInstance()

	s += "\n Retrieved Records\n"
	var name string

	for rows.Next() {
		err := rows.Scan(&name)
		check(err)
		s += name
	}

	io.WriteString(w, s)

}

func GetInstance() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")

	check(err)

	bs := make([]byte, resp.ContentLength)

	resp.Body.Read(bs)

	resp.Body.Close()

	return string(bs)
}

func check(err error) {
	if err != nil {
		panic(err)
		return
	}
}
