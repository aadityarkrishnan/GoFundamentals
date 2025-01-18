package main

import "fmt"

// Define the Shape interface
type Shape interface {
	Area() float64
	Circumference() float64
}

// Define Rectangle struct
type Rectangle struct {
	length float64
	width  float64
}

// Define Circle struct
type Circle struct {
	radius float64
}

// Implement Area for Rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Implement Circumference for Rectangle
func (r Rectangle) Circumference() float64 {
	return 2 * (r.length + r.width)
}

// Implement Area for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Implement Circumference for Circle
func (c Circle) Circumference() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	// Create a Rectangle
	r := Rectangle{length: 10, width: 20}
	// Create a Circle
	c := Circle{radius: 10}

	// Use Shape interface
	var s Shape

	// Assign Rectangle to Shape and call methods
	s = r
	fmt.Printf("Rectangle Area: %.2f\n", s.Area())
	fmt.Printf("Rectangle Circumference: %.2f\n", s.Circumference())

	// Assign Circle to Shape and call methods
	s = c
	fmt.Printf("Circle Area: %.2f\n", s.Area())
	fmt.Printf("Circle Circumference: %.2f\n", s.Circumference())
}
