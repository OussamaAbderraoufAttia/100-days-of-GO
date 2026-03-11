package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go Web Server!")
}

func main() {
	// Register the handler for the root path
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server starting on :8080...")
	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
