package main

import "fmt"

func calculate_cost(message int) float64 {
	total_cost := 0.0
	for i := 0; i < message; i++ {
		total_cost += float64(1) + (0.01 * float64(i))
	}
	return float64(total_cost)

}

func main() {
	cost := calculate_cost(20)
	fmt.Println("cost:", cost)
}
