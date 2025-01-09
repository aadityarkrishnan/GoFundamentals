package main

import (
	"fmt"
)

// Generic function to find an element in a slice
// T is a type parameter that allows the function to work with any type
func Find[T comparable](slice []T, element T) int {
	// Iterate through the slice
	for i, v := range slice {
		if v == element {
			return i // Return the index if the element is found
		}
	}
	return -1 // Return -1 if the element is not found
}

func main() {
	// Example 1: Using Find with a slice of integers
	ints := []int{10, 20, 30, 40, 50}
	fmt.Println("Index of 30 in ints:", Find(ints, 30))   // Output: 2
	fmt.Println("Index of 100 in ints:", Find(ints, 100)) // Output: -1

	// Example 2: Using Find with a slice of strings
	strings := []string{"apple", "banana", "cherry"}
	fmt.Println("Index of 'banana' in strings:", Find(strings, "banana")) // Output: 1
	fmt.Println("Index of 'grape' in strings:", Find(strings, "grape"))   // Output: -1

	// Example 3: Using Find with a slice of custom types
	type Point struct {
		X, Y int
	}
	points := []Point{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("Index of {3, 4} in points:", Find(points, Point{3, 4})) // Output: 1
	fmt.Println("Index of {7, 8} in points:", Find(points, Point{7, 8})) // Output: -1
}
