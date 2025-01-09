package main

import "fmt"

func main() {
	// Initial slice
	numbers := []int{1, 2, 3}

	// Appending a single element
	numbers = append(numbers, 4)

	// Print the updated slice
	fmt.Println("Updated slice:", numbers) // Output: [1 2 3 4]
}
