package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

// Insert a new node at the end of the list
func (n *Node) Insert(newValue int) {
	newNode := &Node{Data: newValue}
	current := n

	// Traverse until the end of the list
	for current.Next != nil {
		current = current.Next
	}

	// Insert the new node
	current.Next = newNode
}

// Display the linked list
func (n *Node) Display() {
	current := n

	// Traverse and print the list
	for current != nil {
		fmt.Print(current.Data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// Reverse the linked list
func (n *Node) Reverse() *Node {
	var prev *Node
	current := n

	// Reverse the linked list in-place
	for current != nil {
		nextNode := current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}

	// Return the new head of the reversed list
	return prev
}

func main() {
	head := &Node{Data: 10}
	head.Insert(20)
	head.Insert(30)
	head.Insert(40)
	head.Insert(50)
	head.Insert(60)
	head.Insert(70)
	head.Insert(80)

	// Display the original list
	fmt.Println("Original List:")
	head.Display()

	// Reverse the list and update the head
	head = head.Reverse()

	// Display the reversed list
	fmt.Println("Reversed List:")
	head.Display()
}
