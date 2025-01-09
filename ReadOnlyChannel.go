package main

import (
	"fmt"
	"time"
)

// Function to simulate reading messages from a read-only channel
func printMessages(ch <-chan string) {
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}

func main() {
	// Create a channel for reading only (the channel is read-only in this main function)
	readOnlyChannel := make(chan string)

	// Start a goroutine to send messages to the channel
	go func() {
		for i := 1; i <= 5; i++ {
			message := fmt.Sprintf("Message #%d", i)
			readOnlyChannel <- message
			time.Sleep(1 * time.Second)
		}
		close(readOnlyChannel) // Close the channel once done sending
	}()

	// Call the function that can only read from the channel
	printMessages(readOnlyChannel)
}
