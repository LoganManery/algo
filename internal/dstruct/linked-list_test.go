package dstruct

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()

	if list.Head != nil {
		t.Errorf("Expected Head to be nil in new list")
	}

	if list.Tail != nil {
		t.Errorf("Expected Tail to be nil in new list")
	}

	if list.Size != 0 {
		t.Errorf("Expected Size to be 0 in new list, got %d", list.Size)
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkedList()

	// Test appending to empty list
	list.Append(10)

	if list.Head == nil {
		t.Errorf("Expected Head to be non-nil after append")
	}

	if list.Head.Value != 10 {
		t.Errorf("Expected 10, got %d")
	}

	if list.Tail != list.Head {
		t.Errorf("Expected Size to be 1, got %d", list.Size)
	}

	if list.Size != 1 {
		t.Errorf("Expected Size to be 1, got %d", list.Size)
	}

	list.Append(20)

	if list.Head.Value != 10 {
		t.Errorf("Expected 10, got %d")
	}

	if list.Tail.Value != 20 {
		t.Errorf("Expected 20, got %d", list.Tail.Value)
	}

	if list.Head.Next != list.Tail {
		t.Errorf("Expected Head.Next to equal Tail after second append")
	}

	if list.Size != 2 {
		t.Errorf("Expected Size to be 2, got %d", list.Size)
	}
}

func TestPrepend(t *testing.T) {
	list := NewLinkedList()

	// Test prepending to empty list
	list.Prepend(10)

	if list.Head == nil {
		t.Errorf("Expected Head to ne non-nil after prepend")
	}

	if list.Head.Value != 10 {
		t.Errorf("Expected Head.Value to be 10, got %d", list.Head.Value)
	}

	if list.Tail != list.Head {
		t.Errorf("Expected Tail to equal Head after single prepend")
	}

	// Try prepending to non-empty list
	list.Prepend(5)

	if list.Head.Value != 5 {
		t.Errorf("Expected Head.Value to be 5, got %d", list.Head.Value)
	}

	if list.Tail.Value != 10 {
		t.Errorf("Expected Tail.Value to remain 10, got %d", list.Tail.Value)
	}

	if list.Head.Next != list.Tail {
		t.Errorf("Expected Head.Next to equal Tail after second prepend")
	}

	if list.Size != 2 {
		t.Errorf("Expected Size to be 2, got %d", list.Size)
	}
}

func TestDelete(t *testing.T) {
	list := NewLinkedList()

	// Test delete on empty list
	result := list.Delete(10)
	if result != false {
		t.Errorf("Expected Delete to return when value exists")
	}

	// Setup list
	list.Append(10)
	list.Append(20)
	list.Append(30)

	// Test delete head
	if !result {
		t.Errorf("Expected Delete to return true when value exists")
	}

	if list.Head.Value != 20 {
		t.Errorf("Expected Head.Value to be 20 after deleteing 10, got %d", list.Head.Value)
	}

	if list.Size != 2 {
		t.Errorf("Expected Size to be 2 after delete, got %d", list.Size)
	}

	list = NewLinkedList()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	result = list.Delete(20)
	if !result {
		t.Errorf("Expected Delete to return when value exists")
	}

	if list.Head.Next.Value != 30 {
		t.Errorf("Expected Head.Next.Value to be 30 after deleting 20, got %d", list.Head.Next.Value)
	}

	if list.Size != 2 {
		t.Errorf("Expected Delete to return true when value exists")
	}

	result = list.Delete(100)
	if result {
		t.Errorf("Expected Delete to return false when value doesn't exist")
	}

	list = NewLinkedList()
	list.Append(10)

	result = list.Delete(10)
	if !result {
		t.Errorf("Expected Delete to return true when value exists")
	}

	if list.Head != nil {
		t.Errorf("Expected Head to be nil after deleting only element")
	}

	if list.Tail != nil {
		t.Errorf("Expected Tail to be nil after deleting only element")
	}
}

func TestIntegration(t *testing.T) {
	list := NewLinkedList()

	// Test various operations in sequence
	list.Append(20)
	list.Prepend(10)
	list.Append(30)

	if list.Size != 3 {
		t.Errorf("Expected Size to be 3, got %d", list.Size)
	}

	// Check values by traversing
	current := list.Head
	expected := []int{10, 20, 30}

	for i := 0; i < len(expected); i++ {
		if current == nil {
			t.Errorf("List ended prematurely at index %d", i)
			break
		}

		if current.Value != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], current.Value)
		}

		current = current.Next
	}

	// Test delete in the middle
	list.Delete(20)

	// Verify resulting list
	current = list.Head
	expected = []int{10, 30}

	for i := 0; i < len(expected); i++ {
		if current == nil {
			t.Errorf("List ended prematurely at index %d", i)
			break
		}

		if current.Value != expected[i] {
			t.Errorf("After delete, at index %d, expected %d, got %d", i, expected[i], current.Value)
		}

		current = current.Next
	}

	if list.Size != 2 {
		t.Errorf("Expected Size to be 2 after delete, got %d", list.Size)
	}
}
