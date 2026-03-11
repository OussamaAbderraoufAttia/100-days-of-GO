package main

import (
	"fmt"
	"time"
)

func main() {
	// A channel that receives a "pulse" every 200 milliseconds
	limiter := time.Tick(200 * time.Millisecond)

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		// This blocks until the limiter channel sends a pulse
		<-limiter
		fmt.Println("Processed request", req, time.Now().Format("15:04:05.000"))
	}
}
