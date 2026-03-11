package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Atomically add 1 to the count
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Final Count (Atomic):", count)
}
