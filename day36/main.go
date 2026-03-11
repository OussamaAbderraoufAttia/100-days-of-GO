package main

import (
	"fmt"
	"day36/calculator" // This path depends on your module name
)

func main() {
	sum := calculator.Add(5, 7)
	fmt.Println("Sum:", sum)
}
