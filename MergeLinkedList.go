package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func (n *Node) Insert(newValue int) {
	newNode := &Node{Data: newValue}
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

}

func (n *Node) Display() {
	current := n
	for current.Next != nil {
		fmt.Print(current.Data)
		current = current.Next
		if current.Next != nil {
			fmt.Print("->")

		}

	}
}

func MergeLinkedList(node1 *Node, node2 *Node) *Node {
	// Dummy node to serve as the starting point of the merged list
	dummy := &Node{}
	tail := dummy

	// Merge the two lists
	for node1 != nil && node2 != nil {
		if node1.Data < node2.Data {
			tail.Next = node1
			node1 = node1.Next
		} else {
			tail.Next = node2
			node2 = node2.Next
		}
		tail = tail.Next
	}

	// Append the remaining elements from either list
	if node1 != nil {
		tail.Next = node1
	} else {
		tail.Next = node2
	}

	return dummy.Next
}

func main() {
	headList1 := &Node{Data: 10}
	headList1.Insert(20)
	headList1.Insert(30)
	headList1.Insert(40)
	headList1.Insert(50)
	//headList1.Display()

	headList2 := &Node{Data: 60}
	headList2.Insert(70)
	headList2.Insert(80)
	headList2.Insert(90)
	headList2.Insert(100)

	mergedList := MergeLinkedList(headList1, headList2)
	mergedList.Display()

}
