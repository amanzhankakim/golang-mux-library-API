package main

import(
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
)

type Book struct {

	ID    string     `json:"id"`
    Name  string  `json:"name"`
    Price int `json:"price"`
	Author *Author `json:"author"`	
}

type Author struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request ){
	json.NewEncoder(w).Encode(books)
}

func postBook(w http.ResponseWriter, r *http.Request ){
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func findBook(w http.ResponseWriter, r *http.Request ){

	params := mux.Vars(r)

	var book Book
	for _, item := range books {
		if item.ID == params["id"]{
			book = item
		}
	}
	json.NewEncoder(w).Encode(book)
}

func main() {


	books = append(books, Book{ID: "1", Name: "The Cuckoo's Calling", Price: 20, Author: &Author{Name: "Robert", Surname: "Galbraith"}})
	router := mux.NewRouter()

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books", postBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", findBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBooks).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8000", router))
}