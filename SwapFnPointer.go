package main

import "fmt"

// Function to swap two values using pointers
func swapPointer(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	a := 10
	b := 20

	// Print before swapping
	fmt.Printf("Before Swapping A=%d, B=%d\n", a, b)

	// Call swapPointer with pointers to a and b
	swapPointer(&a, &b)

	// Print after swapping
	fmt.Printf("After Swapping A=%d, B=%d\n", a, b)
}
