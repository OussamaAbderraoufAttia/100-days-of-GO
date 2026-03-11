package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Goroutine: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Start the function as a goroutine
	go printNumbers()

	fmt.Println("Main: Started goroutine, now waiting...")
	
	// We need to wait here, or the program will exit before the goroutine finishes!
	time.Sleep(600 * time.Millisecond)
	
	fmt.Println("Main: Exiting.")
}
