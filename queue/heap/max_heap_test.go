package heap

import (
	"fmt"
	"github.com/devsquared/gods/queue/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

// When testing the heap, it is important to understand how the slice implementation works. To test that the
// heap is in a good state at any given moment, we can directly check the underlying slice.

func TestBinaryHeap_Add(t *testing.T) {
	type scenario struct {
		name         string
		startingHeap *MaxHeap
		input        Node
		expected     []Node
	}

	maxHeapWithLow := NewMaxHeap()
	maxHeapWithLow.Add(NewNode(0, "testing!"))

	maxHeapWithHigh := NewMaxHeap()
	maxHeapWithHigh.Add(NewNode(99, "testing it all!"))

	maxHeapWithSpread := NewMaxHeap()
	maxHeapWithSpread.Add(NewNode(0, "yee"))
	maxHeapWithSpread.Add(NewNode(99, "haw"))

	testScenarios := []scenario{
		{
			name:         "add a node to an empty heap",
			startingHeap: NewMaxHeap(),
			input:        NewNode(1, "hello!"),
			expected:     []Node{NewNode(1, "hello!")},
		},
		{
			name:         "add a node to a heap with Key lower than rest",
			startingHeap: maxHeapWithHigh,
			input:        NewNode(0, "konichiwa!"),
			expected:     []Node{NewNode(99, "testing it all!"), NewNode(0, "konichiwa!")},
		},
		{
			name:         "add a node to a heap with Key higher than rest",
			startingHeap: maxHeapWithLow,
			input:        NewNode(999999, "woah"),
			expected:     []Node{NewNode(999999, "woah"), NewNode(0, "testing!")},
		},
		{
			name:         "add a node to a heap with a Key in between the rest",
			startingHeap: maxHeapWithSpread,
			input:        NewNode(50, "middle"),
			expected:     []Node{NewNode(99, "haw"), NewNode(0, "yee"), NewNode(50, "middle")},
		},
	}

	for _, ts := range testScenarios {
		// add the new input node
		ts.startingHeap.Add(ts.input)

		actualHeap := ts.startingHeap.Heap
		if !cmp.Equal(actualHeap, ts.expected) {
			t.Fatalf(test.ReportTestFailure(ts.name, actualHeap, ts.expected))
		}
	}
}

func TestMaxHeap_Pop(t *testing.T) {
	type scenario struct {
		name                  string
		startingHeap          *MaxHeap
		expectedValue         any
		expectedErr           error
		expectedRemainingHeap []Node
	}

	maxHeapWithLow := NewMaxHeap()
	maxHeapWithLow.Add(NewNode(0, "testing!"))

	maxHeapWithHigh := NewMaxHeap()
	maxHeapWithHigh.Add(NewNode(99, "testing it all!"))

	maxHeapWithSpread := NewMaxHeap()
	maxHeapWithSpread.Add(NewNode(0, "yee"))
	maxHeapWithSpread.Add(NewNode(99, "haw"))

	testScenarios := []scenario{
		{
			name:                  "pop on an empty heap",
			startingHeap:          NewMaxHeap(),
			expectedErr:           fmt.Errorf("max heap: pop called on empty heap"),
			expectedRemainingHeap: []Node{},
		},
		{
			name:                  "pop on a single node queue",
			startingHeap:          maxHeapWithLow,
			expectedValue:         "testing!",
			expectedRemainingHeap: []Node{},
		},
		{
			name:                  "pop with multiple possible nodes",
			startingHeap:          maxHeapWithSpread,
			expectedValue:         "haw",
			expectedRemainingHeap: []Node{NewNode(0, "yee")},
		},
	}

	for _, ts := range testScenarios {
		actualValue, actualErr := ts.startingHeap.Pop()

		test.CheckErrorsAreSame(actualErr, ts.expectedErr)

		if actualValue != ts.expectedValue {
			t.Fatalf(test.ReportTestFailure(ts.name, actualValue, ts.expectedValue))
		}

		actualHeap := ts.startingHeap.Heap
		if !cmp.Equal(actualHeap, ts.expectedRemainingHeap) {
			t.Fatalf(test.ReportTestFailure(ts.name, actualHeap, ts.expectedRemainingHeap))
		}
	}
}

func TestMaxHeap_GetFirstValue(t *testing.T) {
	type scenario struct {
		name          string
		heap          *MaxHeap
		expectedValue any
		expectedErr   error
		expectedHeap  []Node
	}

	maxHeapWithHigh := NewMaxHeap()
	maxHeapWithHigh.Add(NewNode(99, "testing it all!"))

	maxHeapWithSpread := NewMaxHeap()
	maxHeapWithSpread.Add(NewNode(0, "yee"))
	maxHeapWithSpread.Add(NewNode(99, "haw"))

	maxHeapWithFiveNodes := NewMaxHeap()
	maxHeapWithFiveNodes.Add(NewNode(0, 0))
	maxHeapWithFiveNodes.Add(NewNode(50, 50))
	maxHeapWithFiveNodes.Add(NewNode(200, 200))
	maxHeapWithFiveNodes.Add(NewNode(100, 100))
	maxHeapWithFiveNodes.Add(NewNode(999999, 999999))

	testScenarios := []scenario{
		{
			name:         "try to get first Value from empty heap",
			heap:         NewMaxHeap(),
			expectedErr:  fmt.Errorf("max heap: get first Value called on empty heap"),
			expectedHeap: []Node{},
		},
		{
			name:          "get from only single node heap",
			heap:          maxHeapWithHigh,
			expectedValue: "testing it all!",
			expectedHeap:  []Node{NewNode(99, "testing it all!")},
		},
		{
			name:          "get max Value from heap with 2 nodes",
			heap:          maxHeapWithSpread,
			expectedValue: "haw",
			expectedHeap:  []Node{NewNode(99, "haw"), NewNode(0, "yee")},
		},
		{
			name:          "get max Value from heap with multiple nodes",
			heap:          maxHeapWithFiveNodes,
			expectedValue: 999999,
			expectedHeap:  []Node{NewNode(999999, 999999), NewNode(200, 200), NewNode(50, 50), NewNode(0, 0), NewNode(100, 100)},
		},
	}

	for _, ts := range testScenarios {
		actualValue, actualErr := ts.heap.GetFirstValue()

		test.CheckErrorsAreSame(actualErr, ts.expectedErr)

		if actualValue != ts.expectedValue {
			t.Fatalf(test.ReportTestFailure(ts.name, actualValue, ts.expectedValue))
		}

		actualHeap := ts.heap.Heap
		if !cmp.Equal(actualHeap, ts.expectedHeap) {
			t.Fatalf(test.ReportTestFailure(ts.name, actualHeap, ts.expectedHeap))
		}
	}
}
