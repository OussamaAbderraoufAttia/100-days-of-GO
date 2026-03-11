package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// String implements the fmt.Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	// This automatically calls p.String()
	fmt.Println(p)
	fmt.Printf("Displaying person: %s\n", p)
}
