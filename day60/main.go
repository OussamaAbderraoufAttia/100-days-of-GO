package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: day60 <file_path>")
		return
	}

	filePath := os.Args[1]

	// Get file information
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File: %s\n", info.Name())
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("Modified: %v\n", info.ModTime())
	fmt.Printf("Is Directory: %v\n", info.IsDir())
	fmt.Printf("Permissions: %v\n", info.Mode())
}
