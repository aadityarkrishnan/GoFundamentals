package main

import "fmt"

// Define the Car struct
type Car struct {
	Mileage   int
	No__Wheel int
	Color     string
	FuelType  string
}

// Define the Bike struct
type Bike struct {
	Mileage   int
	No__Wheel int
	Color     string
	FuelType  string
}

// Define the Vehicle interface
type Vehicle interface {
	startVehicle() string
	stopVehicle() string
	FuelEfficiency() string
}

// Implement the Vehicle interface for Bike
func (b Bike) startVehicle() string {
	return "Starting bike vehicle"
}

func (b Bike) stopVehicle() string {
	return "Stopping bike vehicle"
}

func (b Bike) FuelEfficiency() string {
	return fmt.Sprintf("Bike mileage is %d km/l, fuel type is %s", b.Mileage, b.FuelType)
}

// Implement the Vehicle interface for Car
func (c Car) startVehicle() string {
	return "Starting car vehicle"
}

func (c Car) stopVehicle() string {
	return "Stopping car vehicle"
}

func (c Car) FuelEfficiency() string {
	return fmt.Sprintf("Car mileage is %d km/l, fuel type is %s", c.Mileage, c.FuelType)
}

// Function to demonstrate the interface
func main() {
	// Create instances of Car and Bike
	car := Car{
		Mileage:   20,
		No__Wheel: 4,
		Color:     "red",
		FuelType:  "Diesel",
	}

	bike := Bike{
		Mileage:   80,
		No__Wheel: 2,
		Color:     "black",
		FuelType:  "Petrol",
	}

	// Use the Vehicle interface
	var vehicle Vehicle

	// Assign bike to the interface
	vehicle = bike
	fmt.Println(vehicle.startVehicle())
	fmt.Println(vehicle.stopVehicle())
	fmt.Println(vehicle.FuelEfficiency())

	// Assign car to the interface
	vehicle = car
	fmt.Println(vehicle.startVehicle())
	fmt.Println(vehicle.stopVehicle())
	fmt.Println(vehicle.FuelEfficiency())
}
