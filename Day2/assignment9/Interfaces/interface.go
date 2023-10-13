// Description: Create a program that demonstrates the use of interfaces.
// Instructions:
// Define an interface called Shape with a method Area() that returns a float64.
// Create two struct types, Circle and Rectangle, both implementing the Shape interface with their respective Area() methods.
// Create instances of both Circle and Rectangle and calculate and print their areas using the Area() method.

package main

import (
	"fmt"
	"math"
)

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
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	circle := Circle{Radius: 6.0}

	rectangle := Rectangle{Width: 2.0, Height: 3.0}

	fmt.Printf("Area and radius of circle %.2f: %.2f\n", circle.Radius, circle.Area())

	fmt.Printf("Area and width of rect %.2f and height %.2f: %.2f\n", rectangle.Width, rectangle.Height, rectangle.Area())
}
