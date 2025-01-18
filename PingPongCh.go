package main

import (
	"fmt"
	"sync"
)

func addPingToChannel(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "PING"
}

func addPongToChannel(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "PONG"
}

func main() {
	ch := make(chan string, 1) // Buffered channel

	var wg sync.WaitGroup

	// Add an initial "PING" to the channel
	ch <- "PING"

	for i := 0; i < 5; i++ { // Loop 5 times
		wg.Add(1)
		select {
		case msg := <-ch: // Receive from the channel
			fmt.Println("Received:", msg)
			if msg == "PING" {
				go addPongToChannel(ch, &wg) // Add "PONG" to the channel
			} else {
				go addPingToChannel(ch, &wg) // Add "PING" to the channel
			}
		}
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Close the channel
	close(ch)
}
