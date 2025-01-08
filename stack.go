package stack

import (
	"sync"
)

// Stack is a basic LIFO stack that resizes as needed and is safe for concurrent access.
// It works with any type T, where T can be any type.
type Stack[T any] struct {
	mu    sync.Mutex
	items []T
}

// New creates and returns a new stack. It accepts zero or more arguments of type T.
func New[T any](i ...T) *Stack[T] {
	if i != nil {
		return &Stack[T]{items: i, mu: sync.Mutex{}}
	}
	return &Stack[T]{}
}

// Push adds one or more elements to the stack
func (s *Stack[T]) Push(c ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = append(s.items, c...)
}

// Pop removes and returns the top element of the stack
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Len returns the number of items in the stack
func (s *Stack[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}
