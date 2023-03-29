package collection

// Collectioner defines the needed methods to implement a collection.
type Collectioner[T any] interface {
	Add(value T)
	Length() int
	Remove(index int)
}
