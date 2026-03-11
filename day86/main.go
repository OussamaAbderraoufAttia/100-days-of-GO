package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var movies []Movie

func movieHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(movies)
	case "POST":
		var newMovie Movie
		// Read JSON from request body
		if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		movies = append(movies, newMovie)
		fmt.Fprintf(w, "Movie added successfully!")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/movies", movieHandler)

	fmt.Println("API starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
