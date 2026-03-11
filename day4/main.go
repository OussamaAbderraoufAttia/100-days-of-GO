package main

import "fmt"

func main() {
	a := 10.5
	b := 2.5

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
	fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
}
