package main

import "fmt"

func main() {
	number := 17

	if number%2 == 0 {
		fmt.Printf("%d is Even\n", number)
	} else {
		fmt.Printf("%d is Odd\n", number)
	}
}
