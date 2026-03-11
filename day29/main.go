package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// IsAdult is a method on the Person type
// (p Person) is the receiver
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	p1 := Person{Name: "Jack", Age: 17}
	p2 := Person{Name: "Jill", Age: 20}

	fmt.Printf("Is %s an adult? %v\n", p1.Name, p1.IsAdult())
	fmt.Printf("Is %s an adult? %v\n", p2.Name, p2.IsAdult())
}
