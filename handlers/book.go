package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shwxta/gobooks/db"
	"github.com/shwxta/gobooks/models"
)

// CreateBook - Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO books (id, title, author_name, author_birth_year, publisher_name, publisher_year_founded, genres, published_date, pages)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = db.Conn.Exec(context.Background(), query, book.ID, book.Title, book.Author.Name, book.Author.BirthYear, book.Publisher.Name, book.Publisher.YearFounded, book.Genres, book.PublishedDate, book.Pages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Book added successfully")
}

// ReadBook - Get a book by ID
func ReadBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := `SELECT id, title, author_name, author_birth_year, publisher_name, publisher_year_founded, genres, published_date, pages FROM books WHERE id = $1`
	row := db.Conn.QueryRow(context.Background(), query, id)

	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.Author.Name, &book.Author.BirthYear, &book.Publisher.Name, &book.Publisher.YearFounded, &book.Genres, &book.PublishedDate, &book.Pages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// UpdateBook - Update a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE books SET title = $1, author_name = $2, author_birth_year = $3, publisher_name = $4, publisher_year_founded = $5, genres = $6, published_date = $7, pages = $8 WHERE id = $9`
	_, err = db.Conn.Exec(context.Background(), query, book.Title, book.Author.Name, book.Author.BirthYear, book.Publisher.Name, book.Publisher.YearFounded, book.Genres, book.PublishedDate, book.Pages, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Book updated successfully")
}

// DeleteBook - Delete a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := `DELETE FROM books WHERE id = $1`
	_, err := db.Conn.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Book deleted successfully")
}
