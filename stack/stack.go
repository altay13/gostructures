package stack

import (
	"fmt"
)

// Stack implemented with the help of slices
type Stack struct {
	values []interface{}
	top    int

	size int
}

// NewStack method creates the stack
func NewStack() *Stack {
	s := &Stack{
		top:    -1,
		values: make([]interface{}, 0),
		size:   -1,
	}
	return s
}

func (s *Stack) SetSize(size int) error {
	if s.top != -1 {
		return fmt.Errorf("[ ERR ] Stack is not empty. To set the size, stack should be empty.")
	}

	s.size = size
	return nil
}

// Push method inserts the item into the stack
func (s *Stack) Push(item interface{}) {
	if s.top+1 == s.size {
		s.Pop()
	}
	s.values = append(s.values, item)
	s.top++
}

// Pop method returns the item from the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.top < 0 {
		return nil, fmt.Errorf("[ ERR ] Index out of range")
	}

	item := s.values[s.top]
	s.values = append(s.values[:s.top], s.values[s.top+1:]...)
	s.top--

	return item, nil
}

// Peek method returns the top item from stack
// without removing it
func (s *Stack) Peek() (interface{}, error) {

	if s.top < 0 {
		return nil, fmt.Errorf("[ ERR ] Index out of range")
	}

	item := s.values[s.top]

	return item, nil
}

// Size method returns the size of the stack
func (s *Stack) Size() int {
	return (s.top + 1)
}
