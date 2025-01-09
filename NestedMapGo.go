package main

import (
	"fmt"
	"strconv"
)

func incrementPopulation(countries map[string]map[string]string, country string, increment int) error {
	// Check if the country exists in the map
	if _, exists := countries[country]; !exists {
		return fmt.Errorf("country %s not found", country)
	}

	// Get the current population as a string
	populationStr, ok := countries[country]["Population"]
	if !ok {
		return fmt.Errorf("population data not found for country %s", country)
	}

	// Convert the population to an integer
	population, err := strconv.Atoi(populationStr)
	if err != nil {
		return fmt.Errorf("invalid population format for country %s", country)
	}

	// Increment the population
	population += increment

	// Update the map with the new population
	countries[country]["Population"] = strconv.Itoa(population)
	return nil
}

func main() {
	// Define a nested map
	countries := map[string]map[string]string{
		"USA": {
			"Capital":    "Washington, D.C.",
			"Currency":   "USD",
			"Population": "331000000",
		},
		"India": {
			"Capital":    "New Delhi",
			"Currency":   "INR",
			"Population": "1400000000",
		},
		"Germany": {
			"Capital":    "Berlin",
			"Currency":   "EUR",
			"Population": "1400000",
		},
	}

	// Display initial population
	fmt.Println("Initial Population of USA:", countries["USA"]["Population"])

	// Increment the population of the USA
	err := incrementPopulation(countries, "USA", 1000000)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Updated Population of USA:", countries["USA"]["Population"])
	}

	// Attempt to increment the population of a non-existent country
	err = incrementPopulation(countries, "Canada", 500000)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
