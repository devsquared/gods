package queue

// TODO: need to refactor and genercize all implmentations?

type Queue[T any] interface {
	Length() int
	Peek() (T, error)
	Pop() (T, error)
	Push(element T)
}
