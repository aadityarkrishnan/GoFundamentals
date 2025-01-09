package main

import "fmt"

// Variadic function that sums up all integers passed to it
func sumSpreadIntegers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// Creating a slice
	nums := []int{1, 2, 3, 4, 5}

	// Spread the slice elements when calling the variadic function
	result := sumSpreadIntegers(nums...) // Spread slice elements as individual arguments
	fmt.Println("Sum of numbers:", result)
}
