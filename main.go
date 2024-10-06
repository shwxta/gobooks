package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shwxta/gobooks/db"
	"github.com/shwxta/gobooks/handlers"
)

func main() {
	// Connect to database
	db.Connect()
	defer db.Disconnect()

	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.ReadBook).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
