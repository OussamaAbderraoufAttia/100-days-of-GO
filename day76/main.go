package main

import (
	"fmt"
	"sync"
)

func main() {
	// WARNING: THIS CODE CONTAINS A RACE CONDITION
	// Two goroutines try to increment 'count' at the same time.
	count := 0
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			count++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			count++
		}
	}()

	wg.Wait()
	fmt.Println("Final count (might not be 2000):", count)
	fmt.Println("Tip: Run this with 'go run -race main.go'")
}
