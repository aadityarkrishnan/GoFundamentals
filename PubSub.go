package main

import (
	"fmt"
	"time"
)

// Producer function that generates tasks and sends them to the channel
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: generated task %d\n", i)
		ch <- i                 // Send task to the channel
		time.Sleep(time.Second) // Simulate some work
	}
	close(ch) // Close the channel when done sending tasks
}

// Consumer function that processes tasks from the channel
func consumer(id int, ch chan int) {
	for task := range ch { // Receive tasks until the channel is closed
		fmt.Printf("Consumer %d: processing task %d\n", id, task)
		time.Sleep(time.Second) // Simulate some work
	}
}

func main() {
	// Create a channel to pass tasks between producer and consumers
	taskChannel := make(chan int)

	// Launch one producer goroutine
	go producer(taskChannel)

	// Launch two consumer goroutines
	go consumer(1, taskChannel)
	go consumer(2, taskChannel)

	// Wait for all consumers to finish processing tasks
	// Here, we'll sleep for a while to give enough time for the work to complete.
	time.Sleep(8 * time.Second)
}

//aadityaradhakrishnan@Aadityas-MacBook-Air GO % go run PubSub.go
//Producer: generated task 1
//Consumer 1: processing task 1
//Producer: generated task 2
//Consumer 2: processing task 2
//Producer: generated task 3
//Consumer 1: processing task 3
//Producer: generated task 4
//Consumer 2: processing task 4
//Producer: generated task 5
//Consumer 1: processing task 5

// The produced task is being taken by consumer 1, consumer 2
