package main

import (
	"fmt"
	"os"
)

func main() {
	content := "This text will be saved into a file."
	fileName := "output.txt"

	// Convert string to bytes and write to file
	// 0644 means "owner can read/write, others can only read"
	err := os.WriteFile(fileName, []byte(content), 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("Successfully wrote to %s\n", fileName)
}
