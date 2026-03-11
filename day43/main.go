package main

import "fmt"

func main() {
	var i any = "Go Language"

	// Type assertion: extracting the string
	s, ok := i.(string)

	if ok {
		fmt.Println("Inside the interface is a string:", s)
	} else {
		fmt.Println("The interface does not contain a string")
	}

	// Example of a failing assertion
	val, ok := i.(int)
	if !ok {
		fmt.Println("Correct! i is not an int.")
	} else {
		fmt.Println("Value is:", val)
	}
}
