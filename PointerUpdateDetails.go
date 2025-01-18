package main

import "fmt"

type PersonPointer struct {
	Name string
	Age  int32
}

// Method with a pointer receiver to modify the struct fields
func (p *PersonPointer) UpdateDetails(newName string, newAge int32) {
	p.Name = newName
	p.Age = newAge
}

func main() {
	// Initializing a Person struct
	p := PersonPointer{
		Name: "Aaditya",
		Age:  29,
	}

	// Calling the method to update the fields
	p.UpdateDetails("Aaditya R Krishnan", 30)

	// Printing the updated struct
	fmt.Println(p)
}
