package main

import (
	"fmt"
	"time"
)

// Worker function that sends a message to the channel
func worker(id int, ch chan int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Worker %d: sending %d\n", id, i)
		ch <- i // Send data to channel
		time.Sleep(time.Second)
	}
}

// Function to receive data from the channel
func receiver(ch chan int) {
	for i := 0; i < 6; i++ {
		val := <-ch // Receive data from channel
		fmt.Printf("Receiver: received %d\n", val)
	}
}

func main() {
	// Create a channel for communicating between goroutines
	ch := make(chan int)

	// Launch two worker goroutines
	go worker(1, ch)
	go worker(2, ch)

	// Start receiving data in the main goroutine
	go receiver(ch)

	// Wait for a while to allow goroutines to finish
	time.Sleep(5 * time.Second)
}
