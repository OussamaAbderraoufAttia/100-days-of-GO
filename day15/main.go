package main

import "fmt"

func main() {
	var f float64 = 45.99
	var i int = int(f) // Converts 45.99 to 45 (removes decimal)

	fmt.Printf("Float: %f -> Int: %d\n", f, i)

	x := 10
	y := float64(x) // Converts 10 to 10.0

	fmt.Printf("Int: %d -> Float: %f\n", x, y)
}
