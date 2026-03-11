package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	setup := func() {
		fmt.Println("--- Only one initialization should happen ---")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// This will only execute setup() the very first time it's called
			once.Do(setup)
			
			fmt.Printf("Worker %d is done initializing.\n", id)
		}(i)
	}

	wg.Wait()
}
