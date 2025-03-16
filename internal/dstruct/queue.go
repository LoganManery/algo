package dstruct

import (
	"errors"
	"fmt"
)

// Queue represents a FIFO data structure
type Queue struct {
	items []interface{}
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue
// Returns an error if the queue is empty
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Peek returns the item at the front of queue without removing it
// Returns an error if the queue is emtpy
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// Clear removes all items from the queue
func (q *Queue) Clear() {
	q.items = make([]interface{}, 0)
}

// String returns a string representation of the queue
func (q *Queue) String() string {
	if q.IsEmpty() {
		return "Queue: []"
	}

	result := "Queue: ["
	for i, item := range q.items {
		result += fmt.Sprintf("%v", item)
		if i < len(q.items)-1 {
			result += ", "
		}
	}
	result += "]"
	return result
}
