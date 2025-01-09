package main

import (
	"fmt"
	"sync" // Provides the Mutex type for managing concurrency
)

// Global variables
var (
	counter int        // Shared resource that multiple goroutines will access
	mu      sync.Mutex // Mutex to ensure only one goroutine can access the shared resource at a time
)

// increment safely increments the counter using a mutex
func increment(wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the WaitGroup counter when the goroutine finishes

	// Lock the mutex to ensure only one goroutine enters this critical section
	mu.Lock()
	defer mu.Unlock() // Unlock the mutex when the critical section is done

	// Critical section: Increment the shared counter
	counter++
	fmt.Println("Counter:", counter)
}

func main() {
	var wg sync.WaitGroup // Used to wait for all goroutines to finish

	// Spawn 5 goroutines, each incrementing the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)         // Increment the WaitGroup counter for each goroutine
		go increment(&wg) // Start the goroutine
	}

	wg.Wait() // Block until all goroutines finish
	fmt.Println("Final Counter Value:", counter)
}
