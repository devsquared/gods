package queue

import (
	"fmt"
	"github.com/devsquared/gods/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRingQueue_Length(t *testing.T) {
	type scenario struct {
		name           string
		queue          *RingQueue
		expectedLength int
	}

	rQueueWithSingleItem := NewRingQueue()
	rQueueWithSingleItem.Push("item")

	rQueueWithMultipleItem := NewRingQueue()
	rQueueWithMultipleItem.Push("item1")
	rQueueWithMultipleItem.Push("item2")

	testScenarios := []scenario{
		{
			name:           "length of empty queue",
			queue:          NewRingQueue(),
			expectedLength: 0,
		},
		{
			name:           "length of queue with single item",
			queue:          rQueueWithSingleItem,
			expectedLength: 1,
		},
		{
			name:           "length of queue with multiple items",
			queue:          rQueueWithMultipleItem,
			expectedLength: 2,
		},
	}

	for _, ts := range testScenarios {
		actualLength := ts.queue.Length()

		if actualLength != ts.expectedLength {
			test.ReportTestFailure(t, actualLength, ts.expectedLength)
		}
	}
}

func TestRingQueue_Peek(t *testing.T) {
	type scenario struct {
		name           string
		queue          *RingQueue
		copiedQueue    *RingQueue
		expectedValue  any
		expectedErr    error
		expectedLength int
	}

	rQueueWithSingleItem := NewRingQueue()
	rQueueWithSingleItem.Push("item")

	rQueueWithMultipleItem := NewRingQueue()
	rQueueWithMultipleItem.Push("item1")
	rQueueWithMultipleItem.Push("item2")

	testScenarios := []scenario{
		{
			name:           "attempted peak on empty queue",
			queue:          NewRingQueue(),
			copiedQueue:    NewRingQueue(),
			expectedValue:  nil,
			expectedErr:    fmt.Errorf("peek attempted on empty queue"),
			expectedLength: 0,
		},
		{
			name:           "peek on queue with single item",
			queue:          rQueueWithSingleItem,
			copiedQueue:    rQueueWithSingleItem,
			expectedValue:  "item",
			expectedErr:    nil,
			expectedLength: 1,
		},
		{
			name:           "peek on queue with multiple items",
			queue:          rQueueWithMultipleItem,
			copiedQueue:    rQueueWithMultipleItem,
			expectedValue:  "item1",
			expectedErr:    nil,
			expectedLength: 2,
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			actualValue, actualErr := ts.queue.Peek()

			if actualValue != ts.expectedValue {
				test.ReportTestFailure(t, actualValue, ts.expectedValue)
			}

			if !test.IsErrSame(actualErr, ts.expectedErr) {
				test.ReportTestFailure(t, actualErr, ts.expectedErr)
			}

			if ts.queue.Length() != ts.expectedLength {
				test.ReportTestFailure(t, ts.queue.Length(), ts.expectedLength)
			}

			// make sure that the queue is unchanged by a peek
			if !cmp.Equal(ts.queue, ts.copiedQueue) {
				test.ReportTestFailure(t, ts.queue, ts.copiedQueue)
			}
		})
	}
}

func TestRingQueue_Pop(t *testing.T) {
	type scenario struct {
		name           string
		queue          *RingQueue
		expectedQueue  *RingQueue
		expectedValue  any
		expectedErr    error
		expectedLength int
	}

	rQueueWithSingleItem := NewRingQueue()
	rQueueWithSingleItem.Push("item")

	rQueueWithMultipleItem := NewRingQueue()
	rQueueWithMultipleItem.Push("item1")
	rQueueWithMultipleItem.Push("item2")

	rQueueAfterPopOnrQueueWithMultipleItem := NewRingQueue()
	rQueueAfterPopOnrQueueWithMultipleItem.Push("item2")

	testScenarios := []scenario{
		{
			name:           "attempted pop on empty queue",
			queue:          NewRingQueue(),
			expectedQueue:  NewRingQueue(),
			expectedValue:  nil,
			expectedErr:    fmt.Errorf("pop attempted on empty queue"),
			expectedLength: 0,
		},
		{
			name:           "pop on queue with single item",
			queue:          rQueueWithSingleItem,
			expectedQueue:  NewRingQueue(),
			expectedValue:  "item",
			expectedErr:    nil,
			expectedLength: 0,
		},
		{
			name:           "pop on queue with multiple items",
			queue:          rQueueWithMultipleItem,
			expectedQueue:  rQueueAfterPopOnrQueueWithMultipleItem,
			expectedErr:    nil,
			expectedLength: 1,
			expectedValue:  "item1",
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			actualValue, actualErr := ts.queue.Pop()

			if actualValue != ts.expectedValue {
				test.ReportTestFailure(t, actualValue, ts.expectedValue)
			}

			if !test.IsErrSame(actualErr, ts.expectedErr) {
				test.ReportTestFailure(t, actualErr, ts.expectedErr)
			}

			if ts.queue.Length() != ts.expectedLength {
				test.ReportTestFailure(t, ts.queue.Length(), ts.expectedLength)
			}

			// worth noting that the underlying queue is difficult to test as the head and tail changes after adding
			// and removing items. This is fine and expected.
		})
	}
}

func TestRingQueue_Push(t *testing.T) {
	type scenario struct {
		name           string
		queue          *RingQueue
		input          any
		expectedQueue  *RingQueue
		expectedLength int
	}

	rQueueWithSingleItem := NewRingQueue()
	rQueueWithSingleItem.Push("item1")

	rQueueWithMultipleItem := NewRingQueue()
	rQueueWithMultipleItem.Push("item1")
	rQueueWithMultipleItem.Push("item2")

	testScenarios := []scenario{
		{
			name:           "push nil",
			queue:          NewRingQueue(),
			input:          nil,
			expectedQueue:  NewRingQueue(),
			expectedLength: 0,
		},
		{
			name:           "push item on empty queue struct",
			queue:          &RingQueue{},
			input:          "item1",
			expectedQueue:  rQueueWithSingleItem,
			expectedLength: 1,
		},
		{
			name:           "push item on well constructed empty queue",
			queue:          NewRingQueue(),
			input:          "item1",
			expectedQueue:  rQueueWithSingleItem,
			expectedLength: 1,
		},
		{
			name:           "push item on queue with items",
			queue:          rQueueWithSingleItem,
			input:          "item2",
			expectedQueue:  rQueueWithMultipleItem,
			expectedLength: 2,
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			ts.queue.Push(ts.input)

			if ts.queue.Length() != ts.expectedLength {
				test.ReportTestFailure(t, ts.queue.Length(), ts.expectedLength)
			}

			// make sure that the underlying is as expected after pop
			if !cmp.Equal(ts.queue, ts.expectedQueue) {
				test.ReportTestFailure(t, ts.queue, ts.expectedQueue)
			}
		})
	}
}

func TestRingQueue_Push_Resize(t *testing.T) {
	// if the queue needs resizing, a push will do it if needed.
	type scenario struct {
		name               string
		queue              *RingQueue
		input              any
		expectedBufferSize int
	}

	rQueueWithMinSizeFilled := NewRingQueue()

	// fill the minQueueSize
	for i := 0; i < 17; i++ {
		rQueueWithMinSizeFilled.Push("thing")
	}

	testScenarios := []scenario{
		{
			name:               "initial size is the minimumQueueSize",
			queue:              NewRingQueue(),
			input:              nil,
			expectedBufferSize: minRingQueueSize,
		},
		{
			name:               "minQueueSize filled so double size",
			queue:              rQueueWithMinSizeFilled,
			input:              "thing",
			expectedBufferSize: minRingQueueSize * 2,
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			actualBufferSize := len(ts.queue.Buffer)
			if actualBufferSize != ts.expectedBufferSize {
				test.ReportTestFailure(t, actualBufferSize, ts.expectedBufferSize)
			}
		})
	}
}

func TestRingQueue_Pop_Resize(t *testing.T) {
	// if the queue is less than half full, resize buffer to save space
	type scenario struct {
		name               string
		queue              *RingQueue
		expectedBufferSize int
	}

	rQueueWithMinSizeFilledThenHalfPopped := NewRingQueue()

	// fill the minQueueSize and let it double to 32
	for i := 0; i < 17; i++ {
		rQueueWithMinSizeFilledThenHalfPopped.Push("thing")
	}

	// pop off 9 to go back to min size as the total in the queue will be at 8 at this point
	for i := 0; i < 9; i++ {
		_, _ = rQueueWithMinSizeFilledThenHalfPopped.Pop()
	}

	testScenarios := []scenario{
		{
			name:               "half buffer when pop on half empty queue",
			queue:              rQueueWithMinSizeFilledThenHalfPopped,
			expectedBufferSize: minRingQueueSize,
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			actualBufferSize := len(ts.queue.Buffer)
			if actualBufferSize != ts.expectedBufferSize {
				test.ReportTestFailure(t, actualBufferSize, ts.expectedBufferSize)
			}
		})
	}
}
