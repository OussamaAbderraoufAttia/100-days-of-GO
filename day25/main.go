package main

import "fmt"

func main() {
	// Creating and initializing a map
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	// Adding a new entry
	ages["Charlie"] = 35

	fmt.Println("Full map:", ages)

	// Accessing a value
	fmt.Println("Alice's age is:", ages["Alice"])

	// The "comma ok" idiom to check if a key exists
	age, exists := ages["David"]
	if exists {
		fmt.Println("David's age is:", age)
	} else {
		fmt.Println("David is not in the map.")
	}
}
