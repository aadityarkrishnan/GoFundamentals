package main

import (
	"fmt"
)

// Employee interface with two methods
type Employee interface {
	WorkHours() int
	Salary() float64
}

// FullTimeEmployee struct
type FullTimeEmployee struct {
	monthlySalary float64
}

// ContractorEmployee struct
type ContractorEmployee struct {
	hourlyRate  float64
	hoursWorked int
}

// WorkHours method for FullTimeEmployee
func (fte FullTimeEmployee) WorkHours() int {
	return 40 // Fixed 40 hours per week
}

// Salary method for FullTimeEmployee
func (fte FullTimeEmployee) Salary() float64 {
	return fte.monthlySalary
}

// WorkHours method for ContractorEmployee
func (ce ContractorEmployee) WorkHours() int {
	return ce.hoursWorked
}

// Salary method for ContractorEmployee
func (ce ContractorEmployee) Salary() float64 {
	return float64(ce.hoursWorked) * ce.hourlyRate
}

// AdjustSalary function to modify the salary of an employee
func AdjustSalary(emp Employee, adjustment float64) {
	switch e := emp.(type) {
	case *FullTimeEmployee:
		e.monthlySalary += adjustment
	case *ContractorEmployee:
		e.hourlyRate += adjustment
	default:
		fmt.Println("Unknown employee type")
	}
}

func main() {
	// Create instances of employees
	employees := []Employee{
		&FullTimeEmployee{monthlySalary: 4000},
		&ContractorEmployee{hourlyRate: 25, hoursWorked: 30},
	}

	// Iterate through employees and display details
	for _, emp := range employees {
		// Work Hours and Salary
		fmt.Printf("Work Hours: %d\n", emp.WorkHours())
		fmt.Printf("Salary: $%.2f\n", emp.Salary())

		// Type assertion to identify employee type
		switch emp.(type) {
		case *FullTimeEmployee:
			fmt.Println("This is a full-time employee.")
		case *ContractorEmployee:
			fmt.Println("This is a contractor employee.")
		default:
			fmt.Println("Unknown employee type.")
		}
		fmt.Println()
	}

	// Adjust salaries
	fmt.Println("Adjusting salaries...")
	AdjustSalary(employees[0], 500) // Increase full-time salary by $500
	AdjustSalary(employees[1], 2.5) // Increase contractor rate by $2.5

	// Display updated salaries
	for _, emp := range employees {
		fmt.Printf("Updated Salary: $%.2f\n", emp.Salary())
	}
}
