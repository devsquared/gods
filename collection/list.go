package collection

//TODO: should this be of type comparable? or should comparabilty be handled in an iterator or similar?

// List defines an unordered collection of any data.
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

// Add appends a new value.
func (l *List[T]) Add(value T) {
	l.coreSlice = append(l.coreSlice, value)
	l.size++
}

// Remove will take out the element at the given index from the List.
func (l *List[T]) Remove(index int) {

}

// Set will add the value to the list at the given index.
func (l *List[T]) Set(value T, index int) {

}

// Get will get the element at the given index.
func (l *List[T]) Get(index int) any {
	return nil //change any to T
}

// Length returns the size of the List.
func (l *List[T]) Length() int {
	return l.size
}

//TODO: things to consider
// - basic iterator
// - find element
// - sort?
