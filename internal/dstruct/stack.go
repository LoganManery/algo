package dstruct

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item at the top of the stack
// Returns an error if the stack is empty
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the item at the top of the stakc without removing it
// Returns an error is the stack is empty
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack has no items
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear removes all items from the stack
func (s *Stack) Clear() {
	s.items = make([]interface{}, 0)
}

// String returns a string representaion of the stack
func (s *Stack) String() string {
	if s.IsEmpty() {
		return "Stack: []"
	}

	result := "Stack ["
	for i, item := range s.items {
		result += fmt.Sprintf("%v", item)
		if i < len(s.items)-1 {
			result += ", "
		}
	}
	result += "] <- Top"
	return result
}
