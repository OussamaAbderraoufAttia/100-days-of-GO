package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Scan() reads the next line, returns false when done
	for scanner.Scan() {
		lineCount++
		// fmt.Println(scanner.Text()) // To print the line content
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning:", err)
	}

	fmt.Printf("Total lines in file: %d\n", lineCount)
}
