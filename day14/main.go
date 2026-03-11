package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// bufio.NewReader creates a way to read from standard input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Type a sentence: ")
	text, _ := reader.ReadString('\n')

	// Remove the newline character from the end
	text = strings.TrimSpace(text)

	fmt.Printf("You typed: %s\n", text)
}
