package main

import "fmt"

func main() {
	// Create a new channel of strings
	messages := make(chan string)

	// Send a value into the channel from a goroutine
	go func() {
		messages <- "Hello there!"
	}()

	// Receive the value in the main thread
	// This "blocks" (waits) until a value is sent
	msg := <-messages

	fmt.Println("Received:", msg)
}
