package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context that times out after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Good practice to always cancel!

	select {
	case <-time.After(3 * time.Second):
		// This simulates a slow task
		fmt.Println("Task finished (should not happen)")
	case <-ctx.Done():
		// This happens because 2 seconds pass before the task finishes
		fmt.Println("Error:", ctx.Err())
	}
}
