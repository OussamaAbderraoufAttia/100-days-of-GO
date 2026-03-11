package main

import "fmt"

var config string

func init() {
	fmt.Println("Initializing configuration...")
	config = "Running in Development Mode"
}

func main() {
	fmt.Println("Main function started.")
	fmt.Println("Current Config:", config)
}
