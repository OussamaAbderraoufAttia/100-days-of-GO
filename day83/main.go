package main

import (
	"fmt"
	"log"
	"net/http"
)

// LoggingMiddleware logs information about every request
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r) // Call the original handler
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Home!")
}

func main() {
	// Wrap our handler with the middleware
	http.HandleFunc("/", LoggingMiddleware(homeHandler))

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
