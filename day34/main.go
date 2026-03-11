package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func riskyOperation() {
	defer handlePanic()
	fmt.Println("Starting risky operation...")
	
	// Triggering a panic
	panic("Something went horribly wrong!")
	
	fmt.Println("This line will never run")
}

func main() {
	riskyOperation()
	fmt.Println("Program continues normally after recovery.")
}
