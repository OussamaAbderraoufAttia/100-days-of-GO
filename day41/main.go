package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("Area of the shape: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}

	PrintArea(c)
	PrintArea(r)
}
