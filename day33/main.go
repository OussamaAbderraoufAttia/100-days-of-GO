package main

import "fmt"

type DivisionError struct {
	Num   float64
	Denom float64
}

// Implement the error interface
func (e *DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by zero denominator %.2f", e.Num, e.Denom)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivisionError{Num: a, Denom: b}
	}
	return a / b, nil
}

func main() {
	_, err := Divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
}
