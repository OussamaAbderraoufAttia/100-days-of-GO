package main

import (
	"fmt"
	"os"
)

func main() {
	// Try to get "APP_PORT" from the system
	port := os.Getenv("APP_PORT")

	// Set a default if it's empty
	if port == "" {
		port = "8080"
		fmt.Println("APP_PORT not set, using default.")
	}

	fmt.Printf("Server will start on port: %s\n", port)
}
