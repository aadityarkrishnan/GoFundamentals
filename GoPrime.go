package main

import "fmt"

func printPrime(n int) {
	for num := 2; num <= n; num++ {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(num)
		}
	}
}

func main() {
	printPrime(600)
}
