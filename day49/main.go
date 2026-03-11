package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	// flag.String("name", "default value", "description")
	name := flag.String("name", "Guest", "your name")
	age := flag.Int("age", 0, "your age")

	// You MUST call Parse!
	flag.Parse()

	// Flags are pointers, so we use * to get the value
	fmt.Printf("Hello %s! According to your flags, you are %d years old.\n", *name, *age)
}
