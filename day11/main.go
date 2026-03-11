package main

import "fmt"

// divide returns both the quotient and the remainder
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	q, r := divide(10, 3)
	fmt.Printf("10 divided by 3 is %d with a remainder of %d\n", q, r)
}
