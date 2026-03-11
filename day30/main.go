package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// HaveBirthday uses a pointer receiver (*Person) to modify the original struct
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	p := Person{Name: "Bob", Age: 25}
	fmt.Printf("Before birthday: %s is %d years old\n", p.Name, p.Age)

	p.HaveBirthday()

	fmt.Printf("After birthday: %s is %d years old\n", p.Name, p.Age)
}
