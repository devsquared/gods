package heap

import "fmt"

// For ease of implementation, we will use a simple slice or array implementation for a tree.
// This means that we follow these common rules for indices:
// - *Parent Index*: (i - 1) / 2
// - *Children Indices*
//		- Left Child: 2 * i + 1
// 		- Right Child: 2 * i + 2

// MaxHeap represents a heap with the max values towards the top
type MaxHeap struct {
	Heap []Node
}

// NewMaxHeap is a simple constructor to get a max heap.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		Heap: []Node{},
	}
}

// Add inserts a new node into the MaxHeap. After adding, the MaxHeap fixes the remaining nodes ordering.
func (b *MaxHeap) Add(node Node) {
	b.Heap = append(b.Heap, node)
	b.bubbleUp(len(b.Heap) - 1) // bubble up the new node
}

// Pop removes the max keyed node from the MaxHeap. After removing, the MaxHeap fixes the remaining nodes ordering.
func (b *MaxHeap) Pop() (any, error) {
	if len(b.Heap) <= 0 {
		return nil, fmt.Errorf("max heap: pop called on empty heap")
	}

	removed := b.Heap[0]
	heapSize := len(b.Heap)

	if heapSize > 1 {
		b.Heap[0] = b.Heap[len(b.Heap)-1]
	}

	b.Heap = b.Heap[:len(b.Heap)-1]
	b.bubbleDown(0)

	return removed.Value, nil
}

// GetFirstValue returns the Value from the max node of the heap. This does not remove this node from the heap.
// Similar to a peek in a queue.
func (b *MaxHeap) GetFirstValue() (any, error) {
	if len(b.Heap) <= 0 {
		return nil, fmt.Errorf("max heap: get first Value called on empty heap")
	}

	return b.Heap[0].Value, nil
}

func (b *MaxHeap) bubbleUp(index int) {
	for index > 0 {
		parentIndex := getParentIndex(index)

		if b.Heap[parentIndex].Key > b.Heap[index].Key {
			// the node is now in the correct place and is bubbled up; we are done
			return
		}

		b.Heap[parentIndex], b.Heap[index] = b.Heap[index], b.Heap[parentIndex] //swap the nodes
		index = parentIndex
	}
}

func (b *MaxHeap) bubbleDown(index int) {
	for 2*index+1 < len(b.Heap) {
		minChildIndex := b.maxChildIndex(index)

		if b.Heap[minChildIndex].Key < b.Heap[index].Key {
			// the node is now in the correct place and is bubbled down; we are done
			return
		}

		b.Heap[minChildIndex], b.Heap[index] = b.Heap[index], b.Heap[minChildIndex]
		index = minChildIndex
	}
}

func (b *MaxHeap) maxChildIndex(index int) int {
	if getRightIndex(index) >= len(b.Heap) {
		return getLeftIndex(index)
	}

	if b.Heap[getRightIndex(index)].Key > b.Heap[getLeftIndex(index)].Key {
		return getRightIndex(index)
	}

	return getLeftIndex(index)
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftIndex(index int) int {
	return 2*index + 1
}

func getRightIndex(index int) int {
	return 2*index + 2
}
