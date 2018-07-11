package cartesiantree

import (
	"fmt"
	"math/rand"
)

// CTree ...
type CTree struct {
	x uint32
	y float32

	left  *CTree
	right *CTree

	init bool
}

// NewCartesianTree ...
func NewCartesianTree() *CTree {
	t := &CTree{}

	t.init = false

	return t
}

func newCartesianTree(x uint32, y float32, left *CTree, right *CTree) *CTree {
	t := &CTree{}

	t.x = x
	t.y = y

	t.left = nil
	t.right = nil

	if left != nil && left.init {
		t.left = left
	}
	if right != nil && right.init {
		t.right = right
	}

	t.init = true

	return t
}

func (t CTree) merge(l *CTree, r *CTree) *CTree {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	if r.y > l.y {
		newTree := t.merge(l, r.left)
		return newCartesianTree(r.x, r.y, newTree, r.right)
	}

	newTree := t.merge(l.right, r)
	return newCartesianTree(l.x, l.y, l.left, newTree)
}

func (t *CTree) split(x uint32, l *CTree, r *CTree) {
	var newTree CTree

	if t.x <= x {
		if t.right == nil {
			r = nil
		} else {
			t.right.split(x, &newTree, r)
		}

		*l = *newCartesianTree(t.x, t.y, t.left, &newTree)
	} else {
		if t.left == nil {
			l = nil
		} else {
			t.left.split(x, l, &newTree)
		}

		*r = *newCartesianTree(t.x, t.y, &newTree, t.right)
	}
}

// Add ...
func (t *CTree) Add(x uint32) {
	if !t.init {
		t.init = true

		t.x = x
		t.y = rand.Float32()

		return
	}

	var l, r CTree

	t.split(x, &l, &r)

	m := newCartesianTree(x, rand.Float32(), nil, nil)

	ttmp := t.merge(&l, m)
	tmp := t.merge(ttmp, &r)

	t.x = tmp.x
	t.y = tmp.y
	t.left = tmp.left
	t.right = tmp.right
}

// Remove ...
func (t *CTree) Remove(x uint32) {
	var l, m, r CTree

	t.split(x-1, &l, &r)
	r.split(x, &m, &r)

	tmp := t.merge(&l, &r)

	t.x = tmp.x
	t.y = tmp.y
	t.left = tmp.left
	t.right = tmp.right
}

// GetSortedArray ...
func (t *CTree) GetSortedArray(res *[]uint32) {
	if t == nil {
		return
	}

	t.left.GetSortedArray(res)

	*res = append(*res, t.x)

	t.right.GetSortedArray(res)
}

// just for fun...
func (t *CTree) print() {
	if t == nil {
		return
	}
	t.left.print()
	fmt.Printf("%d:%f\n", t.x, t.y)
	t.right.print()
}
