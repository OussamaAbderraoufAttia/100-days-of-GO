package main

import "fmt"

func main() {
	// Initialize an array with 5 elements
	numbers := [5]int{10, 20, 30, 40, 50}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	fmt.Println("Array:", numbers)
	fmt.Println("Sum of elements:", sum)
}
