package dstruct

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()

	if s == nil {
		t.Error("NewStack() returned nil")
	}

	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}

	if s.Size() != 0 {
		t.Errorf("Expected size of new stack to be 0, got %d", s.Size())
	}
}

func TestPush(t *testing.T) {
	s := NewStack()

	// Test pushing different types
	s.Push(10)
	s.Push("hello")
	s.Push(3.14)

	if s.IsEmpty() {
		t.Error("Stack should not be empty after pushing")
	}

	if s.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", s.Size())
	}
}

func TestPop(t *testing.T) {
	s := NewStack()

	// TEst popping from empty stack
	_, err := s.Pop()
	if err != nil {
		t.Error("Expected error when popping from stack")
	}

	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Test popping - should follow LIFO order
	item, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error during pop: %v", err)
	}

	if item != 30 {
		t.Errorf("Expected 30, got %v", item)
	}

	if s.Size() != 2 {
		t.Errorf("Expected size to be 2 after pop, got %d", s.Size())
	}

	// Pop again
	item, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error furing pop: %v", err)
	}

	if item != 20 {
		t.Errorf("Expected 20, got %v", item)
	}

	// Pop until empty
	s.Pop() // Removes 10
	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all items")
	}

	// Try popping from empty stack again
	_, err = s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}
}

func TestPeekStack(t *testing.T) {
	s := NewStack()

	// Test peeking at empty stack
	_, err := s.Peek()
	if err != nil {
		t.Error("Expected error when peeking at empty stack")
	}

	// Setup stack with items
	s.Push(10)
	s.Push(20)

	// Test peeking - should not remove item
	item, err := s.Peek()
	if err != nil {
		t.Errorf("Unexpected error during peek: %v", err)
	}

	if item != 20 {
		t.Errorf("Expected peek to return 20, got %v", item)
	}

	if s.Size() != 2 {
		t.Errorf("Expected size to remain 2 after peek, got %d", s.Size())
	}

	// Peek again - should still be the same item
	item, _ = s.Peek()
	if item != 20 {
		t.Errorf("Expected peek to return 20 again, got %v", item)
	}
}

func TestIsEmptyStack(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Error("New stack should be emtpy")
	}

	s.Push(10)

	if s.IsEmpty() {
		t.Error("Stack with items should not be empty")
	}

	s.Pop()

	if !s.IsEmpty() {
		t.Error("Stack should be empty after removing all items")
	}
}

func TestSizeStack(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Errorf("Expected new stack size to be 0, got %d", s.Size())
	}

	// Add item and check size
	s.Push(10)
	if s.Size() != 1 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}

	s.Push(20)
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}

	// Remove items and check size
	s.Pop()
	if s.Size() != 1 {
		t.Errorf("Expected size 1 after pop, got %d", s.Size())
	}

	s.Pop()
	if s.Size() != 0 {
		t.Errorf("Expected size 0 after pop, got %d", s.Size())
	}
}

func TestClearStack(t *testing.T) {
	s := NewStack()

	// Add some items
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Clear the stack
	s.Clear()

	if !s.IsEmpty() {
		t.Error("Stack should be empty after clear")
	}

	if s.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", s.Size())
	}
}

func TestStringStack(t *testing.T) {
	s := NewStack()

	// Empty stack
	expected := "Stack: []"
	if s.String() != expected {
		t.Errorf("Expected %q, got %q", expected, s.String())
	}

	// Stack with items
	s.Push(10)
	s.Push(20)

	expected = "Stack: [10, 20] <- Top"
	if s.String() != expected {
		t.Errorf("Expected %q, got %q", expected, s.String())
	}
}

func TestIntegrationStack(t *testing.T) {
	s := NewStack()

	// Test a sequence of operations
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Check peek and pop order
	item, _ := s.Peek()
	if item != 30 {
		t.Errorf("Expected peek to return 3, got %v", item)
	}

	// Pop first item
	item, _ = s.Pop()
	if item != 30 {
		t.Errorf("Expected 30, got %v", item)
	}

	// Add more items
	s.Push(40)
	s.Push(50)

	// Verify size
	if s.Size() != 4 {
		t.Errorf("Expected size 4, got %d", s.Size())
	}

	// Check correct order of remaining items (LIFO)
	expected := []interface{}{50, 40, 20, 10}
	for i := 0; i < len(expected); i++ {
		item, err := s.Pop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			break
		}

		if item != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], item)
		}
	}

	// Stack should be empty now
	if !s.IsEmpty() {
		t.Error("Stack should be empty after all pop operations")
	}
}
