package main

import (
	"fmt"
	"time"
)

// Function to simulate sending log messages to a channel
func logSender(ch chan string) {
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Log message #%d", i)
		ch <- message
		time.Sleep(2 * time.Second) // Simulate a delay in sending log messages
	}
	close(ch) // Close the channel when done
}

// Function to simulate sending email notifications to a channel
func emailSender(ch chan string) {
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
	go logSender(logChannel)
	go emailSender(emailChannel)

	// Use select and range to read from both channels concurrently
	for {
		select {
		case logMsg, ok := <-logChannel:
			if ok {
				fmt.Println("Received log:", logMsg)
			} else {
				// Channel closed, stop receiving from this channel
				logChannel = nil
			}
		case emailMsg, ok := <-emailChannel:
			if ok {
				fmt.Println("Received email:", emailMsg)
			} else {
				// Channel closed, stop receiving from this channel
				emailChannel = nil
			}
		}

		// Exit when both channels are closed
		if logChannel == nil && emailChannel == nil {
			break
		}
	}
}
