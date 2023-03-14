package heap

type heap interface {
	Add(node Node)
	Pop() (any, error)
	GetFirstValue() (any, error)
}

type Node struct {
	Key   int
	Value any
}

// NewNode creates a simple node structure for use in a heap.
func NewNode(key int, value any) Node {
	return Node{
		Key:   key,
		Value: value,
	}
}
