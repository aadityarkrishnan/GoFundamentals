package main

import (
	"fmt"
	"sync"
)

func messageFromA(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Message By A"
}

func processMessageByB(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	ch <- msg + " -> Processed by B"
}

func processMessageByC(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	ch <- msg + " -> Processed by C"
}

func main() {
	ch := make(chan string, 1) // Buffered channel to avoid deadlock
	var wg sync.WaitGroup

	// Start the relay cycle 3 times
	for i := 0; i < 3; i++ {
		wg.Add(3) // 3 goroutines per cycle

		// Goroutines to process messages in the correct sequence
		go messageFromA(ch, &wg)
		go processMessageByB(ch, &wg)
		go processMessageByC(ch, &wg)

		// Wait for all goroutines in the current cycle to complete
		wg.Wait()

		// Receive the final message from the channel
		finalMsg := <-ch
		fmt.Printf("Cycle %d: %s\n", i+1, finalMsg)
	}

	// Close the channel after processing
	close(ch)
}
