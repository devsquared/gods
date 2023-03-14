package queue

import (
	"fmt"
	"github.com/devsquared/gods/queue/heap"
	"reflect"
)

// PQItem represents a priority queue item with a value and a priority.
// A higher priority item gets popped sooner than a lower priority item.
type PQItem struct {
	value    any
	priority int
}

// PriorityQueue represents a queue in which the elements of the queue are sorted to be popped based on priority.
// The higher the queue, the sooner it pops from the queue. Due to utilizing a slice-based max heap for implementation,
// resizing and sorting is done as items are added or popped from the queue.
type PriorityQueue struct {
	Heap  *heap.MaxHeap
	Count int
}

// NewPriorityQueue is a simple constructor that creates an empty priority queue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		Heap:  heap.NewMaxHeap(),
		Count: 0,
	}
}

// Pop removes the item with the highest priority from the queue.
func (q *PriorityQueue) Pop() (any, error) {
	// return error if the queue is empty
	if q.Count <= 0 {
		return nil, fmt.Errorf("priority queue: pop called on empty queue")
	}

	poppedValue, err := q.Heap.Pop()
	if err != nil {
		return nil, fmt.Errorf("priority queue: error in pop: %e", err)
	}

	q.Count--
	return poppedValue, nil
}

// Push enqueues an element onto the PriorityQueue. If an element is given that is not a PQItem, a priority of 0 is given.
func (q *PriorityQueue) Push(element any) {
	// in the case that an empty struct was used, let's initialize the underlying heap
	if q.Heap == nil {
		q.Heap = heap.NewMaxHeap()
	}

	var item PQItem
	if reflect.ValueOf(element).Kind() != reflect.ValueOf(PQItem{}).Kind() {
		item = PQItem{value: element, priority: 0}
	} else {
		item = element.(PQItem)
	}

	newHeapNode := heap.NewNode(item.priority, item.value)
	q.Heap.Add(newHeapNode)
	q.Count++
}

// Peek returns the value of the item with the highest priority in the queue. This, however, does not remove the item.
func (q *PriorityQueue) Peek() (any, error) {
	// return error if the queue is empty
	if q.Count <= 0 {
		return nil, fmt.Errorf("priority queue: peek called on empty queue")
	}

	peekValue, err := q.Heap.GetFirstValue()
	if err != nil {
		return nil, fmt.Errorf("priority queue: error in peek: %e", err)
	}

	return peekValue, nil
}

// Length gives the length of the priority queue.
func (q *PriorityQueue) Length() int {
	return q.Count
}
