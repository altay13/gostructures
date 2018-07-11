package rbtree

import "testing"

func TestRBTreeAdd(t *testing.T) {
	tree := NewRBTree()

	tree.Add(5, "Test 5")
	tree.Add(4, "Test 4")
	tree.Add(3, "Test 3")
	tree.Add(2, "Test 2")

	if tree.key != 4 {
		t.Errorf("Failed to add node to tree. Expected root key = 4, instead key = %d", tree.key)
	}

	if tree.left.key != 3 {
		t.Errorf("Failed to add node to tree. Expected root left node key = 4, instead key = %d", tree.left.key)
	}

	if tree.right.key != 5 {
		t.Errorf("Failed to add node to tree. Expected root right node key = 4, instead key = %d", tree.right.key)
	}

	if tree.color != black {
		t.Errorf("Failed to add node to tree. Expected root color = black(false), instead key = %t", tree.color)
	}

	if tree.left.color != black {
		t.Errorf("Failed to add node to tree. Expected root left node color = black(false), instead key = %t", tree.left.color)
	}

	if tree.right.color != black {
		t.Errorf("Failed to add node to tree. Expected root right node color = black(false), instead key = %t", tree.right.color)
	}

	if tree.left.left.color != red {
		t.Errorf("Failed to add node to tree. Expected root left left node color = red(true), instead key = %t", tree.left.left.color)
	}
}

func TestRBTreeGet(t *testing.T) {
	tree := NewRBTree()

	tree.Add(5, "Test 5")
	tree.Add(4, "Test 4")
	tree.Add(3, "Test 3")
	tree.Add(2, "Test 2")

	v := tree.Get(3)
	if v != "Test 3" {
		t.Errorf("Failed to get node from tree. Expected 'Test 3', instead %s", v)
	}

	v = tree.Get(2)
	if v != "Test 2" {
		t.Errorf("Failed to get node from tree. Expected 'Test 2', instead %s", v)
	}
}

func TestRBTreeRemove(t *testing.T) {
	tree := NewRBTree()

	tree.Add(5, "Item 5")
	tree.Add(7, "Item 7")
	tree.Add(2, "Item 2")
	tree.Add(1, "Item 1")
	tree.Add(14, "Item 14")
	tree.Add(10, "Item 10")
	// tree.Add(11, "Item 11")
	// tree.Add(9, "Item 9")
	// tree.Add(6, "Item 6")

	v := tree.right.value
	if v != "Item 10" {
		t.Errorf("Failed to remove the node. Expected '10', instead %s", v)
	}

	tree.Remove(10)

	v = tree.Get(10)
	if v != nil {
		t.Errorf("Failed to remove the node. Expected 'nil', instead %s", v)
	}

	v = tree.right.value
	if v != "Item 14" {
		t.Errorf("Failed to remove the node. Expected 'Item 14', instead %s", v)
	}
}
