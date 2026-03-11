package main

import "fmt"

func main() {
	counter := 0

	for {
		counter++

		if counter == 3 {
			continue // skip printing 3
		}

		fmt.Println(counter)

		if counter >= 5 {
			break // exit the loop
		}
	}
}
