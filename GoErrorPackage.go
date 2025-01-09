package main

import (
	"errors"
	"fmt"
)

// Division function
func spec_divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	a, b := 10, 0
	result, err := spec_divide(a, b)
	if err != nil {
		// Check for the specific error message
		if err.Error() == "division by zero is not allowed" {
			fmt.Printf("Custom error: %s\n", err)
		} else {
			fmt.Printf("General error: %s\n", err)
		}
		return
	}
	fmt.Printf("Result: %d\n", result)
}
