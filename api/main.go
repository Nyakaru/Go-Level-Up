package main

import(
				"encoding/json"
				"log"
				"fmt"
				"math/rand"
				"net/http"
				"strconv"
				"github.com/gorilla/mux"
			)

// Book struct (Model)
type Book struct {
	ID			string  `json:"id"`
	Isbn 		string 	`json:"isbn"`
	Title		string	`json:"title"`
	Author 	*Author	`json:"author"`
}

// Author struct (Author model)
type Author struct {
	Firstname			string  `json:"firstname"`
	Lastname 			string 	`json:"lastname"`
}

// Init var book as a slice
var books [] Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get any params
	
	//Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock Id
	books = append(books, book)
	json.NewEncoder(w).Encode(book) 
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get any params
	
	//Loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index + 1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return 
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init router
	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "422", Title: "Sample", Author: &Author{Firstname: "Kin", Lastname: "Nyaks"}})
	books = append(books, Book{ID: "2", Isbn: "423", Title: "Sample 1", Author: &Author{Firstname: "Mat", Lastname: "Roy"}})
	books = append(books, Book{ID: "3", Isbn: "424", Title: "Sample 2", Author: &Author{Firstname: "Mus", Lastname: "Lera"}})

	// Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("UPDATE")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("Server is up and running....")
	log.Fatal(http.ListenAndServe(":5000", r))
	
}
