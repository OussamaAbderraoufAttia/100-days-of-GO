package main

import (
	"fmt"
	"math"
)

func main() {
	// Square root
	val := 16.0
	fmt.Printf("Square root of %.1f is %.1f\n", val, math.Sqrt(val))

	// Absolute value
	neg := -25.5
	fmt.Printf("Absolute value of %.1f is %.1f\n", neg, math.Abs(neg))

	// Rounding
	dec := 4.7
	fmt.Printf("%.1f rounded is %.0f\n", dec, math.Round(dec))
}
