package main

import "fmt"

func main() {
	value := 1

	for value < 100 {
		fmt.Println(value)
		value *= 2
	}
}
