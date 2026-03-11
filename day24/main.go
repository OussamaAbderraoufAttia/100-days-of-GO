package main

import "fmt"

func main() {
	scores := []int{85, 92, 78, 64, 99}

	fmt.Println("Iterating with range:")
	for index, value := range scores {
		fmt.Printf("Index: %d, Score: %d\n", index, value)
	}

	// Example ignoring the index
	sum := 0
	for _, value := range scores {
		sum += value
	}
	fmt.Println("Total sum:", sum)
}
