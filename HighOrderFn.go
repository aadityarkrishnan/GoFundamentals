package main

import "fmt"

// Higher-order function: Takes a slice of integers and a filter function as input
func filter(nums []int, fn func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if fn(num) {
			result = append(result, num)
		}
	}
	return result
}

// Helper function to filter even numbers
func isEven(num int) bool {
	return num%2 == 0
}

// Helper function to filter numbers greater than 10
func greaterThanTen(num int) bool {
	return num > 10
}

func main() {
	nums := []int{1, 2, 3, 11, 14, 17, 20}

	// Use the higher-order function with different filter functions
	evens := filter(nums, isEven)
	fmt.Println("Even numbers:", evens) // Output: Even numbers: [2 14 20]

	greaterThanTenNums := filter(nums, greaterThanTen)
	fmt.Println("Numbers greater than 10:", greaterThanTenNums) // Output: Numbers greater than 10: [11 14 17 20]
}
