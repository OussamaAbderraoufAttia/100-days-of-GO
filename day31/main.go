package main

import "fmt"

func main() {
	x := 42
	ptr := &x // ptr now holds the memory address of x

	fmt.Printf("x value: %d\n", x)
	fmt.Printf("x address: %p\n", &x)
	fmt.Printf("ptr value (address): %p\n", ptr)
	fmt.Printf("ptr points to: %d\n", *ptr)

	// Changing x via the pointer
	*ptr = 100
	fmt.Printf("New value of x: %d\n", x)
}
