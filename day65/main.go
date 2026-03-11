package main

import "fmt"

func main() {
	queue := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			queue <- i
		}
		// Closing the channel signals that no more data is coming
		close(queue)
	}()

	// range will continue until the channel is closed
	for val := range queue {
		fmt.Println("Received:", val)
	}
	
	fmt.Println("Done.")
}
