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

	min := mySlice[0]
	max := mySlice[0]

	for j := 0; j < n; j++ {
		value := &mySlice[j]
		if *value > max {
			max = *value
		} else if *value < min {
			min = *value
		}
	}

	fmt.Println("Max value: %d", max)
	fmt.Println("Min value: %d", min)
}
