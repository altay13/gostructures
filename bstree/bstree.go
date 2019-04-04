package bstree

import (
	"fmt"
	"reflect"
)

var (
	uniqueness = -1
)

// BSTree Binary Search Tree
type BSTree struct {
	key   interface{}
	value interface{}

	left  *BSTree
	right *BSTree
}

// NewBSTree Initiates Binary Search Tree
func NewBSTree() *BSTree {
	b := &BSTree{
		left:       &BSTree{},
		right:      &BSTree{},
	}
	return b
}

func (b *BSTree) SetUniqueness(v bool) {
	if v {
		uniqueness = -1
	} else {
		uniqueness = 1
	}
}

// Put ...
func (b *BSTree) Put(key interface{}, value interface{}) error {
	if reflect.TypeOf(key) != reflect.TypeOf(b.key) && b.key != nil {
		return fmt.Errorf("Failed to insert value. Wrong format. Expecting %s, instead %s", reflect.TypeOf(b.key), reflect.TypeOf(key))
	}

	b.putRecursive(key, value)

	return nil
}

// Remove ...
func (b *BSTree) Remove(key interface{}) {
	dNode := b.get(key)
	var ttmp *BSTree

	if dNode.left.key == nil && dNode.right.key == nil {
		*dNode = BSTree{}
	} else if dNode.left.key != nil && dNode.right.key == nil {
		ttmp = dNode.left
		*dNode = *ttmp
	} else if dNode.left.key == nil && dNode.right.key != nil {
		ttmp = dNode.right
		*dNode = *ttmp
	} else {
		var parent *BSTree
		lNode := dNode.right

		for {
			if lNode.left.key != nil {
				parent = lNode
				lNode = lNode.left
			} else {
				break
			}
		}

		ttmp := dNode.right
		*dNode = *lNode
		dNode.right = ttmp
		if parent != nil {
			parent.left = nil
		}

	}

}

func (b *BSTree) putRecursive(key interface{}, value interface{}) {
	if b.key == nil {

		b.key = key
		b.value = value
		b.left = &BSTree{}
		b.right = &BSTree{}

		return
	}

	n := b.getCorrectNode(key)

	if n != nil {
		n.putRecursive(key, value)
	}

	return
}

func (b *BSTree) getRecursive(key interface{}) *BSTree {
	if b.key == key {
		return b
	}

	if b.key == nil {
		return nil
	}

	n := b.getCorrectNode(key)

	return n.getRecursive(key)
}

func (b *BSTree) get(key interface{}) *BSTree {
	if reflect.TypeOf(key) != reflect.TypeOf(b.key) && b.key != nil {
		return nil
	}

	node := b.getRecursive(key)

	return node
}

// Get ...
func (b *BSTree) Get(key interface{}) interface{} {
	node := b.get(key)

	if node != nil {
		return node.value
	}

	return nil
}

func (b *BSTree) getCorrectNode(key interface{}) *BSTree {
	res := b.compare(key, b.key)
	if res == 1 {
		return b.right
	} else if res == 0 {
		return b.left
	}

	return nil
}

// GetSortedArray ...
func (b *BSTree) GetSortedArray(res *[]interface{}) {
	if b.key == nil {
		return
	}

	b.left.GetSortedArray(res)

	*res = append(*res, b.value)

	b.right.GetSortedArray(res)
}

// GetSortedArrayOfKeys ...
func (b *BSTree) GetSortedArrayOfKeys(res *[]interface{}) {
	if b.key == nil {
		return
	}

	b.left.GetSortedArrayOfKeys(res)

	*res = append(*res, b.key)

	b.right.GetSortedArrayOfKeys(res)
}

// GetSortedFromLargest ...
func (b *BSTree) GetSortedFromLargest(res *[]interface{}) {
	if b.key == nil {
		return
	}

	b.right.GetSortedFromLargest(res)

	*res = append(*res, b.value)

	b.left.GetSortedFromLargest(res)
}

// Contains ...
func (b *BSTree) Contains(key interface{}) bool {
	if reflect.TypeOf(key) != reflect.TypeOf(b.key) && b.key != nil {
		return false
	}

	r := b.getRecursive(key)
	if r != nil {
		return true
	}

	return false
}

func (b BSTree) compare(it1 interface{}, it2 interface{}) int {
	switch it1.(type) {
	case string:
		if it1.(string) > it2.(string) {
			return 1
		} else if it1.(string) == it2.(string) {
			return uniqueness
		}
	case int:
		if it1.(int) > it2.(int) {
			return 1
		} else if it1.(int) == it2.(int) {
			return uniqueness
		}
	case float32:
		if it1.(float32) > it2.(float32) {
			return 1
		} else if it1.(float32) == it2.(float32) {
			return uniqueness
		}
	case float64:
		if it1.(float64) > it2.(float64) {
			return 1
		} else if it1.(float64) == it2.(float64) {
			return uniqueness
		}
	}
	return 0
}
