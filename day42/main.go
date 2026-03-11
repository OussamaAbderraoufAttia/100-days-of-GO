package main

import "fmt"

// display can take any type of argument
func display(i any) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	display(42)
	display("Hello Go")
	display(3.14)
	display(true)
}
