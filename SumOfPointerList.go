package main

import "fmt"

func main() {
	sum := 0
	z := []int{1, 2, 3, 4, 5}

	// Create pointers to the elements in the slice
	for i := 0; i < len(z); i++ {
		value := &z[i]     // Pointer to the element at index i
		sum = sum + *value // Dereference the pointer to get the value
	}

	fmt.Println(sum)
}
