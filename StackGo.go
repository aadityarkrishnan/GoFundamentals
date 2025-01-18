package main

import "fmt"

type Stack struct {
	Value int
	Next  *Stack
	Size  int
}

// Push method to add a new element to the top of the stack
func (s *Stack) Push(newValue int) {
	newStack := &Stack{Value: newValue, Next: s} // Create a new node and point to the current top
	s.Size++                                     // Increase the stack size
	*s = *newStack                               // Set the top of the stack to the new node
}

// Pop method to remove the top element from the stack
func (s *Stack) Pop() int {
	if s.Size == 0 {
		fmt.Println("Stack is Empty")
		return -1
	}

	value := s.Value // Get the value at the top of the stack
	*s = *s.Next     // Update the top of the stack to the next element
	s.Size--         // Decrease the stack size
	return value
}

// Peek method to get the top element without removing it
func (s *Stack) Peek() int {
	if s.Size == 0 {
		fmt.Println("Stack is Empty")
		return -1
	}

	return s.Value
}

// IsEmpty method to check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func main() {
	stack := &Stack{}
	stack.Push(10)
	fmt.Println("Pushed 10")
	stack.Push(20)
	fmt.Println("Pushed 20")
	stack.Push(30)
	fmt.Println("Pushed 30")

	// Peek the top element
	fmt.Printf("Top element: %d\n", stack.Peek())

	// Perform Pop operations
	fmt.Printf("Popped: %d\n", stack.Pop())
	fmt.Printf("Popped: %d\n", stack.Pop())

	// Peek after popping
	fmt.Printf("Top element: %d\n", stack.Peek())

	// Check if the stack is empty
	fmt.Printf("Is stack empty: %v\n", stack.IsEmpty())
}
