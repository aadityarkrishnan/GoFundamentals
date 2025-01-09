package main

import (
	"fmt"
)

// Define a custom error type
type DivisionByZeroError struct {
	Divisor int
}

// Implement the Error method for DivisionByZeroError with a value receiver
func (e DivisionByZeroError) Error() string {
	return fmt.Sprintf("division by zero is not allowed, provided divisor: %d", e.Divisor)
}

// A function that performs division and returns a custom error if divisor is zero
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		// Return a custom error as a value
		return 0, DivisionByZeroError{
			Divisor: divisor,
		}
	}
	return dividend / divisor, nil
}

func main() {
	// Example 1: Valid division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	// Example 2: Division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("__________________________________________")
		// Use a type assertion to check for the custom error
		if divisionErr, ok := err.(DivisionByZeroError); ok {
			fmt.Printf("Custom error: %s\n", divisionErr.Error())
		} else {
			fmt.Printf("General error: %s\n", err.Error())
		}
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}
