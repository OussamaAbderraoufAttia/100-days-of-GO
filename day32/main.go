package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("cannot take square root of a negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	val := -10.0
	result, err := Sqrt(val)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Sqrt of %.2f is %.2f\n", val, result)
}
