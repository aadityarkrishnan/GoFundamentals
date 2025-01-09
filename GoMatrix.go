package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	// Create a 2D matrix (slice of slices)
	matrix := make([][]int, rows)

	// Initialize each row with a slice of integers
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols) // Create a slice for each row
		// Optionally initialize values in the matrix, e.g., filling with row numbers
		for j := 0; j < cols; j++ {
			matrix[i][j] = i*cols + j // Assign values (example)
		}
	}

	return matrix
}

func main() {
	rows := 3
	cols := 4

	// Create a 3x4 matrix
	matrix := createMatrix(rows, cols)

	// Print the matrix
	for _, row := range matrix {
		fmt.Println(row)
	}
}
