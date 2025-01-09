package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeMap is a struct that wraps a map with RWMutex for thread-safe operations
type SafeMap struct {
	mu   sync.RWMutex // RWMutex to manage read/write access
	data map[string]string
}

// Set writes a key-value pair to the map (write operation)
func (sm *SafeMap) Set(key, value string) {
	sm.mu.Lock()         // Acquire write lock
	defer sm.mu.Unlock() // Release write lock
	sm.data[key] = value
	fmt.Printf("Set key: %s, value: %s\n", key, value)
}

// Get reads a value by key from the map (read operation)
func (sm *SafeMap) Get(key string) string {
	sm.mu.RLock()         // Acquire read lock
	defer sm.mu.RUnlock() // Release read lock
	return sm.data[key]
}

func main() {
	// Initialize the SafeMap
	safeMap := SafeMap{data: make(map[string]string)}

	var wg sync.WaitGroup // WaitGroup to synchronize goroutines

	// Start a writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		safeMap.Set("name", "Aaditya")
		time.Sleep(1 * time.Second) // Simulate some delay
		safeMap.Set("age", "29")
	}()

	// Start multiple reader goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			time.Sleep(500 * time.Millisecond) // Simulate delay to ensure writes start first
			fmt.Printf("Reader %d: name = %s\n", readerID, safeMap.Get("name"))
			fmt.Printf("Reader %d: age = %s\n", readerID, safeMap.Get("age"))
		}(i + 1)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All operations completed.")
}
