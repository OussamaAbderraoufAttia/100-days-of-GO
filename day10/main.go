package main

import "fmt"

// multiply takes two integers and returns one integer
func multiply(a int, b int) int {
	return a * b
}

func main() {
	result := multiply(4, 5)
	fmt.Println("4 * 5 =", result)
}
