package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Stopped.\n", id)
			return
		default:
			fmt.Printf("Worker %d: Still working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// Create a context we can manually stop
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Main: Stopping all workers...")
	
	// Trigger the Done signal
	cancel()

	// Give them a moment to print their exit message
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main: Done.")
}
