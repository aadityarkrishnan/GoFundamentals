package main

import (
	"fmt"
	"time"
)

// Producer function that generates tasks and sends them to the buffered channel
func buffer_producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: generated task %d\n", i)
		ch <- i                 // Send task to the buffered channel
		time.Sleep(time.Second) // Simulate some work
	}
	close(ch) // Close the channel after producing all tasks
}

// Consumer function that processes tasks from the buffered channel
func buffer_consumer(id int, ch chan int) {
	for task := range ch { // Receive tasks until the channel is closed
		fmt.Printf("Consumer %d: processing task %d\n", id, task)
		time.Sleep(2 * time.Second) // Simulate some work
	}
}

func main() {
	// Create a buffered channel with a capacity of 3
	taskChannel := make(chan int, 3)

	// Launch one producer goroutine
	go buffer_producer(taskChannel)

	// Launch two consumer goroutines
	go buffer_consumer(1, taskChannel)
	go buffer_consumer(2, taskChannel)

	// Wait for a while to allow consumers to finish processing tasks
	time.Sleep(10 * time.Second)
}

//?A buffered channel in Go allows sending and receiving data with a specified capacity,
//meaning the channel can hold multiple items without blocking until it reaches its maximum capacity.
//Buffered channels are useful when you want to allow the producer to send data without waiting for
//the consumer to be ready immediately, up to a certain limit.
