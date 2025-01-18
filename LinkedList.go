package main

import "fmt"

// Define the Node struct
type Node struct {
	Data int
	Next *Node
}

// Method to insert a new node at the end of the linked list
func (n *Node) Insert(newElem int) {
	newNode := &Node{Data: newElem} // Create a new node
	current := n

	// Traverse to the last node
	for current.Next != nil {
		current = current.Next
	}

	// Attach the new node to the last node's "Next" pointer
	current.Next = newNode
}

// Method to display all nodes in the linked list
func (n *Node) Display() {
	current := n
	// Traverse and print all nodes
	for current != nil {
		fmt.Print(current.Data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Initialize an empty linked list with the first node
	head := &Node{Data: 10}

	// Insert new elements into the linked list
	head.Insert(20)
	head.Insert(30)
	head.Insert(40)
	head.Insert(50)

	// Display the linked list after inserting some elements
	fmt.Println("Linked list values:")
	head.Display()

	// Insert more elements
	head.Insert(60)
	head.Insert(70)

	// Display the linked list after inserting more elements
	fmt.Println("Linked list values after inserting more elements:")
	head.Display()
}
