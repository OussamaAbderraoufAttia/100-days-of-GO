package main

import "fmt"

func main() {
	fmt.Println("Starting...")

	// This will run at the very end of main()
	defer fmt.Println("Ending (deferred)")

	fmt.Println("Middle...")
	fmt.Println("Still in the middle...")
}
