package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := time.After(2100 * time.Millisecond) // This returns a channel that receives once

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t.Format("15:04:05.000"))
		case <-done:
			ticker.Stop() // Important: stop the ticker!
			fmt.Println("Timer expired, stopping ticker.")
			return
		}
	}
}
