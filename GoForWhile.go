package main

import "fmt"

func checkForResult() {
	i := 0
	for i <= 10 {
		if i%2 == 0 {
			fmt.Printf("The %d is even\n", i)
		} else {
			fmt.Printf("The %d is odd\n", i)
		}
		i++
	}
}

func main() {
	checkForResult()
}
