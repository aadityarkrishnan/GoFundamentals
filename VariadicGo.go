package main

import "fmt"

// Variadic function that sums up all integers passed to it
func sumIntegers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// Calling the variadic function with different numbers of arguments
	fmt.Println(sumIntegers(1, 2, 3))       // Output: 6
	fmt.Println(sumIntegers(5, 10, 15, 20)) // Output: 50
	fmt.Println(sumIntegers(100))           // Output: 100
	fmt.Println(sumIntegers())              // Output: 0 (no numbers passed)
}
