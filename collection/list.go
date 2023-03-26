package collection

type List[T any] struct {
	coreSlice []T
	size      int
}

// NewList constructs a new list with the given type T.
func NewList[T any]() *List[T] {
	return &List[T]{
		coreSlice: make([]T, 0),
		size:      0,
	}
}

// NewListFromSlice constructs a new list from a given slice with the given type T.
func NewListFromSlice[T any](slice []T) *List[T] {
	return &List[T]{
		coreSlice: slice,
		size:      len(slice),
	}
}

// Empty removes all elements from the List and reduces its size to 0.
func (l *List[T]) Empty() {
	l.size = 0
	l.coreSlice = make([]T, 0)
}

//TODO: add these via TDD. we will want the following functions for a list:
// - add
// - remove
// - set
// - get
// - empty
// - length

//TODO: things to consider
// - basic iterator
// - find element
// - sort?
