package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Create a Reader from a string
	content := "This is a stream of data being copied to stdout.\n"
	reader := strings.NewReader(content)

	// io.Copy(writer, reader)
	// os.Stdout is an io.Writer
	fmt.Print("Result of io.Copy: ")
	_, err := io.Copy(os.Stdout, reader)

	if err != nil {
		fmt.Println("Error occurred during copy:", err)
	}
}
