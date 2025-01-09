package main

import "fmt"

func main() {
	// Define a slice
	numbers := []int{10, 20, 30, 40, 50}

	// Iterate over the slice using for range
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
