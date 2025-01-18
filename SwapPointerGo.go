package main

import "fmt"

func main() {
	a := 8
	b := 4

	// Use pointers to swap the values
	c := &b // Pointer to b
	*c = a  // Assign value of a to what c points to (b)
	a = b   // Assign value of b to a

	fmt.Println("a:", a, "b:", b)
}
