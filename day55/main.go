package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Go uses a special date for formatting: 
	// 1 (month), 2 (day), 15 (hour), 04 (min), 05 (sec), 06 (year)
	fmt.Println("Default:", now.String())
	fmt.Println("Custom :", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Date   :", now.Format("Monday, Jan 2, 2006"))
}
