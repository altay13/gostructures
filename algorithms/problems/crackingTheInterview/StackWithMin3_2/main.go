package main

import (
	"fmt"

	"github.com/altay13/gostructures/stack"
)

// StackWithMin is a stack which also returns minimum element
// Complexity O(1)
type StackWithMin struct {
	*stack.Stack
	minStack *stack.Stack
}

// NewStackWithMin returns empty StackWithMin
func NewStackWithMin() *StackWithMin {
	s := stack.NewStack()
	m := stack.NewStack()
	sm := &StackWithMin{
		Stack:    s,
		minStack: m,
	}

	return sm
}

// Push overwrites the Push method for adding
// elements to linked list
func (s *StackWithMin) Push(item interface{}) {
	s.Stack.Push(item)

	m, err := s.Min()
	if err != nil {
		s.minStack.Push(item)
		return
	}

	if item.(int) <= m.(int) {
		s.minStack.Push(item)
	}
}

// Pop overwrites the Pop method for handling
// minimum Stack
func (s *StackWithMin) Pop() (interface{}, error) {
	v, err := s.Stack.Pop()
	if err != nil {
		return nil, err
	}
	if min, err := s.Min(); err == nil {
		if v == min {
			s.minStack.Pop()
		}
	}

	return v, nil
}

// Min ...
func (s *StackWithMin) Min() (interface{}, error) {
	return s.minStack.Peek()
}

func main() {
	s := NewStackWithMin()

	s.Push(1)
	s.Push(2)
	s.Push(5)
	s.Push(7)
	s.Push(3)
	s.Push(4)
	s.Push(0)
	s.Push(1)
	s.Push(10)

	r, _ := s.Min()
	fmt.Println("Min:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Min()
	fmt.Println("Min:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Min()
	fmt.Println("Min:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Min()
	fmt.Println("Min:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)
	r, _ = s.Pop()
	fmt.Println("Pop:", r)

	r, _ = s.Min()
	fmt.Println("Min:", r)
}
