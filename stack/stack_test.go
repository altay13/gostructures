package stack

import (
	"testing"
)

func TestStackPushAndPop(t *testing.T) {
	s := NewStack()
	s.Push("test_1")

	if v, err := s.Pop(); err == nil {
		if "test_1" != v && s.top == 0 {
			t.Errorf("Failed to Pop")
		}
	} else {
		t.Errorf("Failed to Pop")
	}

	if s.top != -1 {
		t.Errorf("The top should be -1. Instead: %d", s.top)
	}

	s.Push("test_2")
	s.Push("test_3")

	if s.top != 1 {
		t.Errorf("The top should be 1. Instead: %d", s.top)
	}

	if v, err := s.Pop(); err == nil {
		if "test_3" != v && s.top == 1 {
			t.Errorf("Failed to Pop")
		}
	} else {
		t.Errorf("Failed to Pop")
	}
}

func TestStackSize(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Errorf("The size of the stack should be 0. Instead %d", s.Size())
	}

	s.Push("Test")
	s.Push("Test_2")
	s.Push("Test_3")

	if s.Size() != 3 {
		t.Errorf("The size of the stack should be 3. Instead %d", s.Size())
	}
}
