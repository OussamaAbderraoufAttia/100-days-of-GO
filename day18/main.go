package main

import "fmt"

// Package level variable
var message = "Hello from Package level"

func main() {
	// Function level variable (shadowing)
	message := "Hello from Function level"

	fmt.Println(message)
	printPackageLevel()
}

func printPackageLevel() {
	// Since there is no local 'message', it uses the package level one
	fmt.Println(message)
}
