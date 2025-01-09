package main

import (
	"fmt"
	"strings"
)

// Function to replace profane words using pointers
func replaceProfanity(sentences *[]string, profanities []string) {
	// Dereference the pointer to access the slice
	for i := range *sentences {
		for _, word := range profanities {
			// Replace profanity with asterisks
			censored := strings.Repeat("*", len(word))
			(*sentences)[i] = strings.ReplaceAll((*sentences)[i], word, censored)
		}
	}
}

func main() {
	// List of sentences
	sentences := []string{
		"This is a damn good example.",
		"That idiot messed it up!",
		"What a stupid thing to do.",
	}

	// List of profane words to censor
	profanities := []string{"damn", "idiot", "stupid"}

	// Call the function with a pointer to the sentences
	replaceProfanity(&sentences, profanities)

	// Print the cleaned sentences
	fmt.Println("Censored Sentences:")
	for _, sentence := range sentences {
		fmt.Println(sentence)
	}
}
