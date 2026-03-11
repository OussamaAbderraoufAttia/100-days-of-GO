package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Lock the resource
	c.mu.Lock()
	c.count++
	// Unlock when finished
	c.mu.Unlock()
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.Increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Count:", counter.count)
}
