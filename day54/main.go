package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{45, 12, 89, 33, 7, 100}

	fmt.Println("Before sorting:", numbers)

	// Sort in place
	sort.Ints(numbers)

	fmt.Println("After sorting :", numbers)
}
