package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("postgres", "postgres://rootuser:password@localhost/bookstore?sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database.")
}

func main() {

	defer db.Close()

	http.HandleFunc("/books", books)
	http.HandleFunc("/books/show", bookShow)

	http.ListenAndServe(":8080", nil)

}

func books(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	defer rows.Close()

	books := make([]Book, 0)

	for rows.Next() {

		book := Book{}

		err = rows.Scan(&book.isbn, &book.title, &book.author, &book.price)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, book := range books {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", book.isbn, book.title, book.author, book.price)
	}

}

func bookShow(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM books WHERE isbn=$1", isbn)

	book := Book{}

	err := row.Scan(&book.isbn, &book.title, &book.author, &book.price)

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", book.isbn, book.title, book.author, book.price)

}
