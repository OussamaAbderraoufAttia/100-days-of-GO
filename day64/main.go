package main

import "fmt"

func main() {
	// A buffered channel with capacity for 2 strings
	ch := make(chan string, 2)

	// We can send these without needing a goroutine to wait for them
	ch <- "First"
	ch <- "Second"

	// If we sent a third, it would block here!
	// ch <- "Third"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
