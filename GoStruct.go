package main

import "fmt"

// Basic Struct in Go
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// Embedded Struct in Go
type Employee struct {
	Person         // Embedding the Person struct
	company string // Additional field specific to Employee
}

// Method for Person // this is a struct method
func (p Person) displayPersonDetails() {
	fmt.Printf("Name: %s, Age: %d, Job: %s, Salary: %d\n", p.name, p.age, p.job, p.salary)
}

// Method for Employee
func (e Employee) displayEmployeeDetails() {
	e.displayPersonDetails() // Call the method from the embedded Person struct
	fmt.Printf("Company: %s\n", e.company)
}

func main() {
	// Create a Person instance
	person := Person{
		name:   "Alice",
		age:    30,
		job:    "Engineer",
		salary: 80000,
	}

	// Create an Employee instance
	employee := Employee{
		Person:  person, // Embed the Person instance
		company: "TechCorp",
	}

	// Anonymous struct for a temporary representation
	car := struct {
		name  string
		brand string
		year  int
	}{
		name:  "Model S",
		brand: "Tesla",
		year:  2022,
	}

	// Display details of the anonymous struct
	fmt.Println("Anonymous Struct Details:")
	fmt.Printf("Car Name: %s, Brand: %s, Year: %d\n", car.name, car.brand, car.year)

	// Accessing fields directly through embedding
	fmt.Println("\nEmployee Details (Access Fields Directly):")
	fmt.Printf("Name: %s, Age: %d, Job: %s, Salary: %d, Company: %s\n",
		employee.name, employee.age, employee.job, employee.salary, employee.company)

	// Using methods to display details
	fmt.Println("\nEmployee Details (Using Methods):")
	employee.displayEmployeeDetails()
}
