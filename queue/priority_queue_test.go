package queue

import (
	"fmt"
	"github.com/devsquared/gods/queue/heap"
	"github.com/devsquared/gods/queue/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPriorityQueue_Length(t *testing.T) {
	type scenario struct {
		name           string
		queue          *PriorityQueue
		expectedLength int
	}

	pQueueWithSingleItem := NewPriorityQueue()
	pQueueWithSingleItem.Push(PQItem{value: "hi", priority: 0})

	pQueueWithMultipleItems := NewPriorityQueue()
	pQueueWithMultipleItems.Push(PQItem{value: "hello", priority: 1})
	pQueueWithMultipleItems.Push(PQItem{value: "hiya", priority: 2})

	testScenarios := []scenario{
		{
			name:           "length of empty queue",
			queue:          NewPriorityQueue(),
			expectedLength: 0,
		},
		{
			name:           "length of queue with one item",
			queue:          pQueueWithSingleItem,
			expectedLength: 1,
		},
		{
			name:           "length of queue with multiple items",
			queue:          pQueueWithMultipleItems,
			expectedLength: 2,
		},
	}

	for _, ts := range testScenarios {
		actualLength := ts.queue.Length()

		if actualLength != ts.expectedLength {
			test.ReportTestFailure(ts.name, actualLength, ts.expectedLength)
		}
	}
}

func TestPriorityQueue_Peek(t *testing.T) {
	type scenario struct {
		name           string
		queue          *PriorityQueue
		expectedValue  any
		expectedErr    error
		expectedLength int
		copiedQueue    *PriorityQueue
	}

	pQueueWithSingleItem := NewPriorityQueue()
	pQueueWithSingleItem.Push(PQItem{value: "hi", priority: 0})

	pQueueWithMultipleItems := NewPriorityQueue()
	pQueueWithMultipleItems.Push(PQItem{value: "hello", priority: 1})
	pQueueWithMultipleItems.Push(PQItem{value: "hiya", priority: 2})

	testScenarios := []scenario{
		{
			name:           "peek on empty queue",
			queue:          NewPriorityQueue(),
			expectedErr:    fmt.Errorf("priority queue: peek called on empty queue"),
			expectedValue:  nil,
			expectedLength: 0,
			copiedQueue:    NewPriorityQueue(),
		},
		{
			name:           "peek on queue with single item",
			queue:          pQueueWithSingleItem,
			expectedErr:    nil,
			expectedValue:  "hi",
			expectedLength: 1,
			copiedQueue:    pQueueWithSingleItem,
		},
		{
			name:           "peek on queue with multiple items",
			queue:          pQueueWithMultipleItems,
			expectedErr:    nil,
			expectedValue:  "hiya",
			expectedLength: 2,
			copiedQueue:    pQueueWithMultipleItems,
		},
	}

	for _, ts := range testScenarios {
		actualValue, actualErr := ts.queue.Peek()

		if actualValue != ts.expectedValue {
			test.ReportTestFailure(ts.name, actualValue, ts.expectedValue)
		}

		test.CheckErrorsAreSame(actualErr, ts.expectedErr)

		if ts.queue.Length() != ts.expectedLength {
			test.ReportTestFailure(ts.name, ts.queue.Length(), ts.expectedLength)
		}

		// make sure that the queue is unchanged by a peek
		if !cmp.Equal(ts.queue, ts.copiedQueue) {
			test.ReportTestFailure(ts.name, ts.queue, ts.copiedQueue)
		}
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	type scenario struct {
		name           string
		queue          *PriorityQueue
		expectedValue  any
		expectedErr    error
		expectedLength int
		underlyingHeap *heap.MaxHeap
	}

	pQueueWithSingleItem := NewPriorityQueue()
	pQueueWithSingleItem.Push(PQItem{value: "hi", priority: 0})

	pQueueWithMultipleItems := NewPriorityQueue()
	pQueueWithMultipleItems.Push(PQItem{value: "hello", priority: 1})
	pQueueWithMultipleItems.Push(PQItem{value: "hiya", priority: 2})

	heapFromPQueueWithSingleItemAfterPop := heap.NewMaxHeap() // becomes empty; left for read- and reason-ability

	heapFromPQueueWithMultipleItemsAfterPop := heap.NewMaxHeap()
	heapFromPQueueWithMultipleItemsAfterPop.Add(heap.NewNode(1, "hello")) // heap left with one node after pop

	testScenarios := []scenario{
		{
			name:           "pop on an empty queue",
			queue:          NewPriorityQueue(),
			expectedValue:  nil,
			expectedErr:    fmt.Errorf("priority queue: pop called on empty queue"),
			expectedLength: 0,
			underlyingHeap: heap.NewMaxHeap(), // empty heap
		},
		{
			name:           "pop on queue with single item",
			queue:          pQueueWithSingleItem,
			expectedValue:  "hi",
			expectedErr:    nil,
			expectedLength: 0,
			underlyingHeap: heapFromPQueueWithSingleItemAfterPop, // becomes empty after pop
		},
		{
			name:           "pop on queue with multiple items",
			queue:          pQueueWithMultipleItems,
			expectedValue:  "hiya",
			expectedErr:    nil,
			expectedLength: 1,
			underlyingHeap: heapFromPQueueWithMultipleItemsAfterPop,
		},
	}

	for _, ts := range testScenarios {
		actualValue, actualErr := ts.queue.Pop()

		if actualValue != ts.expectedValue {
			test.ReportTestFailure(ts.name, actualValue, ts.expectedValue)
		}

		test.CheckErrorsAreSame(actualErr, ts.expectedErr)

		if ts.queue.Length() != ts.expectedLength {
			test.ReportTestFailure(ts.name, ts.queue.Length(), ts.expectedLength)
		}

		if ts.queue.Length() != ts.expectedLength {
			test.ReportTestFailure(ts.name, ts.queue.Length(), ts.expectedLength)
		}

		// make sure that the underlying is as expected after pop
		if !cmp.Equal(ts.queue.Heap, ts.underlyingHeap) {
			test.ReportTestFailure(ts.name, ts.queue.Heap, ts.underlyingHeap)
		}
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	type scenario struct {
		name           string
		queue          *PriorityQueue
		input          any
		expectedQueue  *PriorityQueue
		expectedLength int
	}

	pQueueWithSingleItem := NewPriorityQueue()
	pQueueWithSingleItem.Push(PQItem{value: "hello", priority: 1})

	pQueueWithMultipleItems := NewPriorityQueue()
	pQueueWithMultipleItems.Push(PQItem{value: "hello", priority: 1})
	pQueueWithMultipleItems.Push(PQItem{value: "hiya", priority: 2})

	pQueueWithMultipleItemsWithPushedZeroPrio := pQueueWithMultipleItems
	pQueueWithMultipleItemsWithPushedZeroPrio.Push("item")

	testScenarios := []scenario{
		{
			name:           "push on empty struct",
			queue:          &PriorityQueue{}, // create empty struct without constructor
			input:          nil,
			expectedQueue:  NewPriorityQueue(),
			expectedLength: 0,
		},
		{
			// For this test, it is important to also test the casting to a PQItem in the case that wasn't given.
			name:           "push non-PQItem to queue for casting and queueing at 0 priority",
			queue:          pQueueWithMultipleItems,
			input:          "item",
			expectedQueue:  pQueueWithMultipleItemsWithPushedZeroPrio,
			expectedLength: 3,
		},
		{
			name:           "push on empty queue",
			queue:          NewPriorityQueue(),
			input:          PQItem{value: "hi", priority: 0},
			expectedQueue:  pQueueWithSingleItem,
			expectedLength: 1,
		},
		{
			name:           "push on queue with current items",
			queue:          pQueueWithSingleItem,
			input:          PQItem{value: "hiya", priority: 2},
			expectedQueue:  pQueueWithMultipleItems,
			expectedLength: 2,
		},
	}

	for _, ts := range testScenarios {
		ts.queue.Push(ts.input)

		if ts.queue.Length() != ts.expectedLength {
			test.ReportTestFailure(ts.name, ts.queue.Length(), ts.expectedLength)
		}

		// make sure that the underlying is as expected after pop
		if !cmp.Equal(ts.queue, ts.expectedQueue) {
			test.ReportTestFailure(ts.name, ts.queue, ts.expectedQueue)
		}
	}
}
