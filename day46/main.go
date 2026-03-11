package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "test.txt"

	// Create a dummy file first for the example
	os.WriteFile(filePath, []byte("Hello, this is a test file content!"), 0644)

	// Read the whole file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:")
	fmt.Println(string(data))
}
