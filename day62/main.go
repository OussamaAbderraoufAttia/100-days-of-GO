package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Signal that this goroutine is done when the function exits
	defer wg.Done()

	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished!\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	fmt.Println("Main: Waiting for all workers...")
	wg.Wait() // Block until the counter is 0
	fmt.Println("Main: All workers finished.")
}
