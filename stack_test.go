package stack_test

import (
	"testing"

	"github.com/josuedeavila/stack"
	"github.com/matryer/is"
)

func TestStack_New(t *testing.T) {
	h := is.New(t)
	t.Run("Creates new empty stack", func(t *testing.T) {
		s := stack.New[int]()
		if s == nil {
			h.Fail()
		}
		h.Equal(0, s.Len())
	})

	t.Run("Creates new filled stack", func(t *testing.T) {
		i := []int{1, 2, 3}
		s := stack.New(i...)
		if s == nil {
			h.Fail()
		}
		h.Equal(3, s.Len())
	})

	t.Run("Creates new stack with complex type", func(t *testing.T) {
		type complex struct {
			Name string
			Age  int
		}
		i := []complex{{"John", 30}, {"Doe", 40}}
		s := stack.New(i...)
		if s == nil {
			h.Fail()
		}
		h.Equal(2, s.Len())
	})
}

func TestStack_Push(t *testing.T) {
	h := is.New(t)
	t.Run("Pushes values onto the stack", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		h.Equal(3, s.Len())
		for i := 3; i > 0; i-- {
			v, _ := s.Pop()
			h.Equal(i, v)
		}
	})

	t.Run("Pushes multiple values onto the stack", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(1, 2, 3)
		h.Equal(3, s.Len())
		for i := 3; i > 0; i-- {
			v, _ := s.Pop()
			h.Equal(i, v)
		}
	})

	t.Run("Pushes complex type onto the stack", func(t *testing.T) {
		type complex struct {
			Name string
			Age  int
		}
		s := stack.New[complex]()
		s.Push(complex{"John", 30}, complex{"Doe", 40})
		h.Equal(2, s.Len())
	})
}

func TestStack_Pop(t *testing.T) {
	h := is.New(t)
	t.Run("Pops values off the stack", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		h.Equal(3, s.Len())
		for i := 3; i > 0; i-- {
			v, _ := s.Pop()
			h.Equal(i, v)
		}
	})

	t.Run("Returns zero value when stack is empty", func(t *testing.T) {
		s := stack.New[int]()
		v, ok := s.Pop()
		h.Equal(0, v)
		h.Equal(false, ok)
	})

	t.Run("Pops complex type off the stack", func(t *testing.T) {
		type complex struct {
			Name string
			Age  int
		}
		s := stack.New[complex]()
		s.Push(complex{"John", 30}, complex{"Doe", 40})
		h.Equal(2, s.Len())
		v, _ := s.Pop()
		h.Equal("Doe", v.Name)
		h.Equal(40, v.Age)
	})
}
