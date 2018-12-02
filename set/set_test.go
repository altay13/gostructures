package set

import "testing"

func TestSetAddRemove(t *testing.T) {
	s := NewSet(10)

	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(1)

	if s.Size() != 3 {
		t.Errorf("Failed to add item to Set. Expecting 3, instead %d", s.Size())
	}

	s.Remove(1)

	if s.Size() != 2 {
		t.Errorf("Failed to add item to Set. Expecting 2, instead %d", s.Size())
	}
}

func TestSetContains(t *testing.T) {
	s := NewSet(10)

	s.Add(10)
	s.Add(2)
	s.Add(3)
	s.Add(11)

	if s.Contains(1) {
		t.Errorf("Failed to check contains in Set. Expecting false, instead true")
	}

	if !s.Contains(2) {
		t.Errorf("Failed to check contains in Set. Expecting true, instead false")
	}
}
