package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "name" parameter from the URL
	name := r.URL.Query().Get("name")
	
	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/hello", greetHandler)

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
