package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go is an amazing language!"

	// Get length
	fmt.Println("Length:", len(text))

	// Convert to uppercase
	fmt.Println("Uppercase:", strings.ToUpper(text))

	// Check for substring
	contains := strings.Contains(text, "amazing")
	fmt.Println("Contains 'amazing'?", contains)
}
