package main

import "fmt"

func main() {
	inventory := map[string]int{
		"Apples":  10,
		"Bananas": 20,
		"Oranges": 15,
	}

	fmt.Println("Before delete:", inventory)

	// Remove "Bananas" from the map
	delete(inventory, "Bananas")

	fmt.Println("After delete:", inventory)
}
