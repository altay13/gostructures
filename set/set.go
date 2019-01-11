package set

import "github.com/altay13/gostructures/hashtable"

// Set ...
type Set struct {
	hash     *hashtable.Hash
	capacity int
}

func NewSet(n int) *Set {
	s := &Set{
		hash:     hashtable.NewHash(n),
		capacity: n,
	}

	return s
}

// Add ...
func (s *Set) Add(item interface{}) {
	s.hash.Add(item, item)
}

// Remove ...
func (s *Set) Remove(item interface{}) interface{} {
	e := s.hash.Remove(item)
	return e
}

// Contains ...
func (s *Set) Contains(item interface{}) bool {
	e := s.hash.Get(item)

	if e != nil {
		return true
	}

	return false
}

// Size ...
func (s *Set) Size() int {
	return s.hash.Size()
}

// RemoveAll ...
func (s *Set) RemoveAll() {
	s.hash.RemoveAll()
}
