package dstruct

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	if q == nil {
		t.Error("NewQueue() returned nil")
	}

	if !q.IsEmpty() {
		t.Error("New queue should ne empty")
	}

	if q.Size() != 0 {
		t.Errorf("Expected size of new queue to be 0, got %d", q.Size())
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	// Test enqueuing different types
	q.Enqueue(10)
	q.Enqueue("hello")
	q.Enqueue(3.14)

	if q.IsEmpty() {
		t.Error("Queue should not be empty after enqueuing")
	}

	if q.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", q.Size())
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()

	// Test dequeuing from empty queue
	_, err := q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}

	// Setup queue with items
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Test dequeuing - should follow FIFO order
	item, err := q.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error during dequeue: %v", err)
	}

	if item != 10 {
		t.Errorf("Expected 10, got %v", item)
	}

	if q.Size() != 2 {
		t.Errorf("Expected size to be 2 after dequeue")
	}

	// Dequeue again
	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error during dequeue: %v", err)
	}

	if item != 20 {
		t.Error("Expected 20, got %v", item)
	}

	// Dequeue until empty
	q.Dequeue() // Removes 30

	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all items")
	}

	// Try dequeuing from empty queue again
	_, err = q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue()

	// Test peeking at empty queue
	_, err := q.Peek()
	if err == nil {
		t.Error("Expected error when peeking at empty queue")
	}

	// Setup queue with items
	q.Enqueue(10)
	q.Enqueue(20)

	// Test peeking - should not remove item
	item, err := q.Peek()
	if err != nil {
		t.Errorf("Unexpected error during peek: %v", err)
	}

	if item != 10 {
		t.Errorf("Expected peek to return 10, got %v", item)
	}

	if q.Size() != 2 {
		t.Errorf("Expected size to remain 2 after peek, got %d", q.Size())
	}

	// Peek again - should still be the same item
	item, err = q.Peek()
	if item != 10 {
		t.Errorf("Expected peek to return 10 again, got %v", item)
	}
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.Enqueue(10)

	if q.IsEmpty() {
		t.Errorf("Queue with items should not be empty")
	}

	q.Dequeue()

	if !q.IsEmpty() {
		t.Error("Queue should be empty after removing all items")
	}
}

func TestSize(t *testing.T) {
	q := NewQueue()

	if q.Size() != 0 {
		t.Errorf("Expected new queue size to be 0, got %d", q.Size())
	}

	// Add items and check size
	q.Enqueue(10)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}

	q.Enqueue(20)
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}

	// Remove items and check size
	q.Dequeue()
	if q.Size() != 1 {
		t.Errorf("Expected size 0 after dequeue, got %d", q.Size())
	}
}

func TestClear(t *testing.T) {
	q := NewQueue()

	// Add some items
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Clear the queue
	q.Clear()

	if !q.IsEmpty() {
		t.Error("Queue should be empty after clear")
	}

	if q.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", q.Size())
	}
}

func TestString(t *testing.T) {
	q := NewQueue()

	// Empty queue
	expected := "Queue: []"
	if q.String() != expected {
		t.Errorf("Expected %q, got %q", expected, q.String())
	}

	// Queue with items
	q.Enqueue(10)
	q.Enqueue(20)

	expected = "Queue: [10, 20]"
	if q.String() != expected {
		t.Errorf("Expected %q, got %q", expected, q.String())
	}
}

func TestIntergation(t *testing.T) {
	q := NewQueue()

	// Test a sequence of operations
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Check peek abd dequeue order
	item, _ := q.Peek()
	if item != 10 {
		t.Errorf("Expected peek to return 10, got %v", item)
	}

	// Dequeue first item
	item, _ = q.Dequeue()
	if item != 10 {
		t.Errorf("Expected 10, got %v", item)
	}

	// Add more items
	q.Enqueue(40)
	q.Enqueue(50)

	if q.Size() != 4 {
		t.Errorf("Expected size 4, got %d", q.Size())
	}

	expected := []interface{}{20, 30, 40, 50}
	for i := 0; i < len(expected); i++ {
		item, err := q.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			break
		}

		if item != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], item)
		}
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after all dequeue operations")
	}
}
