package main

import "fmt"

func main() {
	// make([]type, length, capacity)
	numbers := make([]int, 0, 5)

	fmt.Printf("Starting: len=%d cap=%d\n", len(numbers), cap(numbers))

	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("Added %2d: len=%2d cap=%2d\n", i, len(numbers), cap(numbers))
	}
}
