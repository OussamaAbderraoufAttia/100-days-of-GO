package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var movies = []Movie{
	{ID: 1, Title: "The Matrix", Year: 1999},
	{ID: 2, Title: "Inception", Year: 2010},
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Encode the slice into the response writer
	json.NewEncoder(w).Encode(movies)
}

func main() {
	http.HandleFunc("/movies", getMovies)

	fmt.Println("API starting on :8080/movies")
	http.ListenAndServe(":8080", nil)
}
