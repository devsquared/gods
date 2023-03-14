package queue

import (
	"fmt"
)

// Basic ring queue implementation with a basic slice of interface{}.
// With this implementation of a ring queue, we utilize bit-masking to speed up the processes.
// This means that it is important to keep the size of the buffer a power of 2.

// minRingQueueSize starts at 16 and must be a power of 2.
const minRingQueueSize = 16

// RingQueue represents the ring buffer queue.
type RingQueue struct {
	Buffer []any
	Head   int // marker of the head in the slice
	Tail   int // marker of the tail in the slice
	Count  int // length of the queues contents; NOT necessarily total length of queue's buffer
}

// NewRingQueue constructs a new RingQueue instance.
func NewRingQueue() *RingQueue {
	return &RingQueue{
		Buffer: make([]any, minRingQueueSize), //create a buffer the minimum size to begin
	}
}

// Length gets the length of the queue currently.
func (q *RingQueue) Length() int {
	return q.Count
}

// resize handles resizing the queue whenever it is needed. This will either double its length if space is needed
// or it will shrink the size if the queue is less than half full.
func (q *RingQueue) resize() {
	//start by doubling the size
	newBuffer := make([]any, q.Count<<1)

	//now appropriately copy the contents
	if q.Tail > q.Head {
		copy(newBuffer, q.Buffer[q.Head:q.Tail])
	} else {
		n := copy(newBuffer, q.Buffer[q.Head:])
		copy(newBuffer[n:], q.Buffer[:q.Tail])
	}

	//now reset the values in the newly resized queue instance
	q.Head = 0
	q.Tail = q.Count
	q.Buffer = newBuffer
}

// Push enqueues a new element on to the end of the queue.
func (q *RingQueue) Push(element any) {
	// if the element is nil, we don't need to add that
	if element == nil {
		return
	}

	// if the Buffer is uninitialized, let's initialize it
	if q.Buffer == nil {
		q.Buffer = make([]any, minRingQueueSize)
	}

	// if we have run out of room, let's resize
	if q.Count == len(q.Buffer) {
		q.resize()
	}

	q.Buffer[q.Tail] = element
	q.Tail = (q.Tail + 1) & (len(q.Buffer) - 1) //bitwise modulus using AND
	q.Count++
}

// Peek provides utility to see the front of the queue. Returns an error whenever the queue is empty.
func (q *RingQueue) Peek() (any, error) {
	// if the queue is empty, error
	if q.Count <= 0 {
		return nil, fmt.Errorf("peek attempted on empty queue")
	}
	return q.Buffer[q.Head], nil
}

// Pop dequeues the element from the front of the queue and returns it. If the queue is empty, an error is returned.
func (q *RingQueue) Pop() (any, error) {
	// if the queue is empty, error
	if q.Count <= 0 {
		return nil, fmt.Errorf("pop attempted on empty queue")
	}
	result := q.Buffer[q.Head]                  // get the result
	q.Buffer[q.Head] = nil                      // remove result from queue
	q.Head = (q.Head + 1) & (len(q.Buffer) - 1) // bitwise modulus using AND
	q.Count--

	// if buffer is bigger than minimum size and 1/4 full, resize
	if len(q.Buffer) > minRingQueueSize && (q.Count<<2) == len(q.Buffer) {
		q.resize()
	}

	return result, nil
}
