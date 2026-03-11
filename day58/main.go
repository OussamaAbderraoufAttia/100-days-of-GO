package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14159

	// Use reflection to inspect x
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("Type :", t)
	fmt.Println("Value:", v)

	// We can also see the kind (float64)
	fmt.Println("Kind :", v.Kind())
}
