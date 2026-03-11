package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Create public folder and file for the demo
	os.Mkdir("public", 0755)
	os.WriteFile("public/index.html", []byte("<h1>Welcome to my static site!</h1>"), 0644)

	// Create a FileServer for the "public" directory
	fs := http.FileServer(http.Dir("./public"))

	// Handle all requests by checking the public folder
	http.Handle("/", fs)

	fmt.Println("Serving files on :8080...")
	http.ListenAndServe(":8080", nil)
}
