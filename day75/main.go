package main

import (
	"fmt"
	"sync"
)

func main() {
	// Define a pool of objects
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("Creating a new object...")
			return make([]byte, 1024)
		},
	}

	// 1. Get an object from the pool
	buf := pool.Get().([]byte)
	fmt.Println("Buffer length:", len(buf))

	// 2. Do some work with it...
	
	// 3. Put it back when done
	pool.Put(buf)

	// 4. Get it again (it might be the same one!)
	buf2 := pool.Get().([]byte)
	fmt.Println("Buffer 2 length:", len(buf2))
}
