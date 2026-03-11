package main

import "fmt"

func main() {
	// Initializing a slice
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Initial slice:", fruits)

	// Appending new items
	fruits = append(fruits, "Date", "Elderberry")
	fmt.Println("Updated slice:", fruits)
	fmt.Println("Number of fruits:", len(fruits))
}
