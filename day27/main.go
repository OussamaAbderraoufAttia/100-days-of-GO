package main

import "fmt"

// Define the struct
type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Create an instance using a struct literal
	p := Person{
		Name:  "John Doe",
		Age:   28,
		Email: "john@example.com",
	}

	fmt.Println("Person Details:")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Email:", p.Email)
}
