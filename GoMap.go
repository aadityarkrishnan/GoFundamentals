package main

import (
	"errors"
	"fmt"
)

func addToMap(users []string, phones []int64) (map[string]int64, error) {
	// Validate that the input slices are of equal length
	if len(users) != len(phones) {
		return nil, errors.New("Invalid input: users and phones slices must have the same length")
	}

	// Create a map to store user-phone pairs
	userMap := make(map[string]int64)

	// Populate the map
	for i := 0; i < len(users); i++ {
		userMap[users[i]] = phones[i]
	}

	return userMap, nil
}

func main() {
	users := []string{"Alice", "Bob", "Charlie"}
	phones := []int64{1234567890, 9876543210, 1122334455}

	// Call the addToMap function
	userMap, err := addToMap(users, phones)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the resulting map
	for name, phone := range userMap {
		fmt.Printf("Name: %s, Phone: %d\n", name, phone)
	}
}
