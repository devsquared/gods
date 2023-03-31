package queue

import "fmt"

// TODO: let's test this and figure out how to return a nil T

// Stack is a Last In, Last Out (LIFO) data structure. It is similar to a queue but with the difference being
// how elements are popped off.
type Stack[T any] struct {
	coreSlice []T
}

// Length returns the number of elements currently in the stack.
func (s *Stack[T]) Length() int {
	return len(s.coreSlice)
}

// Peek returns the top most element or last added element. This does not remove the element from the stack.
func (s *Stack[T]) Peek() (T, error) {
	/if len(s.coreSlice) == 0 {
		return T{}, fmt.Errorf("stack is empty")
	} else {
		index := s.Length() - 1 // index of top most element
		peeked := s.coreSlice[index]
		return peeked, nil
	}
}

// Pop removes and returns the top most element or last added from the stack.
func (s *Stack[T]) Pop() (T, error) {
	if len(s.coreSlice) == 0 {
		return T{}, fmt.Errorf("stack is empty")
	} else {
		index := s.Length() - 1 // index of top most element
		popped := s.coreSlice[index]
		s.coreSlice = s.coreSlice[:index] // slice the underlying slice off at top
		return popped, nil
	}
}

// Push adds a new element to the top of the stack.
func (s *Stack[T]) Push(element T) {
	s.coreSlice = append(s.coreSlice, element)
}
