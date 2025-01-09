package main

import "fmt"

// Struct representing a counter
type Counter struct {
	Value int
}

// Method with a pointer receiver to increment the counter
func (c *Counter) Increment() {
	c.Value++
}

// Method with a pointer receiver to reset the counter
func (c *Counter) Reset() {
	c.Value = 0
}

// Method with a value receiver (doesn't modify the original Counter)
func (c Counter) Display() {
	fmt.Println("Counter value:", c.Value)
}

func main() {
	// Create a Counter instance
	counter := Counter{Value: 10}

	// Call methods
	counter.Display() // Output: Counter value: 10

	counter.Increment()
	counter.Display() // Output: Counter value: 11

	counter.Reset()
	counter.Display() // Output: Counter value: 0
}
