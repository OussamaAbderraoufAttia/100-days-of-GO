package main

import (
	"fmt"
	"runtime"
)

func main() {
	// NumCPU returns the actual number of cores your computer has
	numCores := runtime.NumCPU()
	fmt.Printf("Your computer has %d CPU cores.\n", numCores)

	// GOMAXPROCS sets/gets the number of cores Go is allowed to use
	// Passing a value < 1 just returns the current setting
	currentProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("Go is currently using %d cores.\n", currentProcs)

	// You could limit it like this:
	// runtime.GOMAXPROCS(2) 
}
