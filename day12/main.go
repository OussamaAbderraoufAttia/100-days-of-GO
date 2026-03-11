package main

import "fmt"

// divide uses named return values for clarity
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return // naked return: automatically returns quotient and remainder
}

func main() {
	q, r := divide(20, 6)
	fmt.Printf("20 / 6 = %d, remainder = %d\n", q, r)
}
