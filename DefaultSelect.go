package main

import (
	"fmt"
	"time"
)

func logDefaultSender(ch chan string) {
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Log message #%d", i)
		ch <- message
		time.Sleep(2 * time.Second) // Simulate a delay in sending log messages
	}
	close(ch) // Close the channel when done
}

func emailDefaultSender(ch chan string) {
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Email notification #%d", i)
		ch <- message
		time.Sleep(3 * time.Second) // Simulate a delay in sending email notifications
	}
	close(ch) // Close the channel when done
}

func main() {
	// Create two channels: one for logs and one for emails
	logChannel := make(chan string)
	emailChannel := make(chan string)

	// Start goroutines to simulate log and email sending
	go logDefaultSender(logChannel)
	go emailDefaultSender(emailChannel)

	// Use select with default to read from both channels concurrently
	for {
		select {
		case logMsg, ok := <-logChannel:
			if ok {
				fmt.Println("Received log:", logMsg)
			}
		case emailMsg, ok := <-emailChannel:
			if ok {
				fmt.Println("Received email:", emailMsg)
			}
		default:
			// Default case: If no channel is ready, perform this action
			fmt.Println("No messages available, doing other work...")
			time.Sleep(1 * time.Second) // Simulate doing some work
		}

		// Exit when both channels are closed
		if len(logChannel) == 0 && len(emailChannel) == 0 {
			break
		}
	}
}
