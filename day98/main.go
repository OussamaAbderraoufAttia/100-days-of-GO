package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var books []Book

func listBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/books", listBooks)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
