package dstruct

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (ll *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}

	if ll.Head == nil {
		// List is empty
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.Size++
}

// Prepend adds a new node with the given value to the beginning of the list
func (ll *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: ll.Head}
	ll.Head = newNode

	if ll.Tail == nil {
		// If the lsit was empty
		ll.Tail = newNode
	}
	ll.Size++
}

// Delete removes the first node with the given value
func (ll *LinkedList) Delete(value int) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.Size--

		// If the list is now empty
		if ll.Head == nil {
			ll.Tail = nil
		}
		return true
	}

	// Search for the value
	current := ll.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	// If value was found
	if current.Next != nil {
		// Remove the node
		if current.Next == ll.Tail {
			ll.Tail = current
		}
		current.Next = current.Next.Next
		ll.Size--
		return true
	}
	return false
}

func (ll *LinkedList) Print() {
	current := ll.Head
	fmt.Print("List: ")
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
