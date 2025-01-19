package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter the limit")
	fmt.Scan(&n)

	mySlice := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&mySlice[i]) // Read each element into the slice
	}

	var sum int = 0
	for j := 0; j < n; j++ {
		value := &mySlice[j]
		sum = sum + *value
	}

	// Calculate the average
	avg := float64(sum) / float64(n)

	// Print the average
	fmt.Printf("The average of %d numbers is %.2f\n", n, avg)

}
