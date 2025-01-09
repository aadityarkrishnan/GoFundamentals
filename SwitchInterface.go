package main

import "fmt"

// Shape interface defines the method Area that different shapes will implement
type SwitchShape interface {
	Area() float64
}

// Circle struct represents a circle with a radius
type SwitchCircle struct {
	Radius float64
}

// Rectangle struct represents a rectangle with width and height
type SwitchRectangle struct {
	Width, Height float64
}

// Circle implements the Shape interface
func (c SwitchCircle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle implements the Shape interface
func (r SwitchRectangle) Area() float64 {
	return r.Width * r.Height
}

// Function to calculate the area of different shapes using a type switch
func calculateArea(s SwitchShape) {
	switch shape := s.(type) {
	case SwitchCircle:
		fmt.Printf("Circle Area: %.2f\n", shape.Area())
	case SwitchRectangle:
		fmt.Printf("Rectangle Area: %.2f\n", shape.Area())
	default:
		fmt.Println("Unknown Shape")
	}
}

func main() {
	// Create instances of Circle and Rectangle
	circle := SwitchCircle{Radius: 5}
	rectangle := SwitchRectangle{Width: 4, Height: 6}

	// Use the calculateArea function with different shapes
	calculateArea(circle)    // Calls the Circle case
	calculateArea(rectangle) // Calls the Rectangle case
}
