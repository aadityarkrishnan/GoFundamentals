package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct for Rectangle
type Rectangle struct {
	width, height float64
}

// Implement Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Define a struct for Circle
type Circle struct {
	radius float64
}

// Implement Shape interface for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// A function that works with any Shape
func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// Create instances of Rectangle and Circle
	rect := Rectangle{width: 10, height: 5}
	circle := Circle{radius: 7}

	fmt.Println("Rectangle Details:")
	printShapeDetails(rect)

	fmt.Println("\nCircle Details:")
	printShapeDetails(circle)
}
