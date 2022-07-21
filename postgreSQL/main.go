package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

var db *sql.DB
var tpl *template.Template

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

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	defer db.Close()

	http.HandleFunc("/books", books)
	http.HandleFunc("/books/show", bookShow)
	http.HandleFunc("/books/create", createBook)
	http.HandleFunc("/books/store", storeBook)
	http.HandleFunc("/books/update", updateBook)
	http.HandleFunc("/books/update/process", updateBookProcess)
	http.HandleFunc("/books/delete", deleteBook)

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

		err = rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)

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
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", book.Isbn, book.Title, book.Author, book.Price)
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

	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", book.Isbn, book.Title, book.Author, book.Price)

}

func createBook(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "create.html", nil)

}

func storeBook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	book := Book{}

	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	if book.Isbn == "" || book.Title == "" || book.Author == "" || price == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(price, 32)

	if err != nil {
		http.Error(w, http.StatusText(406)+" Please go back and enter a valid amount", http.StatusNotAcceptable)
		return
	}

	book.Price = float32(f64)

	_, err = db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", book.Isbn, book.Title, book.Author, book.Price)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
	return

}

func updateBook(w http.ResponseWriter, r *http.Request) {

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

	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "update.html", book)

}

func updateBookProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	book := Book{}

	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	if book.Isbn == "" || book.Title == "" || book.Author == "" || price == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(price, 32)

	if err != nil {
		http.Error(w, http.StatusText(406)+" Please go back and enter a valid amount", http.StatusNotAcceptable)
		return
	}

	book.Price = float32(f64)

	_, err = db.Exec("UPDATE books SET title=$2, author=$3, price=$4 WHERE isbn=$1", book.Isbn, book.Title, book.Author, book.Price)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
	return

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	book := Book{}

	book.Isbn = r.FormValue("isbn")

	if book.Isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("DELETE FROM books WHERE isbn=$1", book.Isbn)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
	return

}
