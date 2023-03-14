package queue

type Queue interface {
	Peek() (any, error)
	Pop() (any, error)
	Push(element any)
}
