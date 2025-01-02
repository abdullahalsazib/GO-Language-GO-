package main

import (
	"fmt"
)

// Define the Shap interface
type Shap interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Implement Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle struct
type Circle struct {
	Radius float64
}

// Implement Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implement Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	fmt.Println("Welcome to interfaces in Go!")

	// Create instances of Rectangle and Circle
	r := Rectangle{Width: 5, Height: 7}
	c := Circle{Radius: 10}

	// Declare a Shap interface variable
	var s Shap

	// Assign Rectangle to the interface
	s = r
	fmt.Println("Rectangle Area: ", s.Area())
	fmt.Println("Rectangle Perimeter: ", s.Perimeter())

	// Assign Circle to the interface
	s = c
	fmt.Println("Circle Area: ", s.Area())
	fmt.Println("Circle Perimeter: ", s.Perimeter())
}
