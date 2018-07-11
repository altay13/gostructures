package avltree

import "testing"

func TestAVLTreeAdd(t *testing.T) {
	tree := NewAVLTree()

	tree.Add(5, "Test 5")
	tree.Add(4, "Test 4")
	tree.Add(3, "Test 3")
	tree.Add(2, "Test 2")

	if getHeight(tree) != 3 {
		t.Errorf("Failed to add. Expected height = 3, instead got %d", getHeight(tree))
	}

	v := tree.value
	if v != "Test 4" {
		t.Errorf("Failed to add item to AVL Tree. Expected 'Test 4', instead %s", v)
	}

	v = tree.left.value
	if v != "Test 3" {
		t.Errorf("Failed to add item to AVL Tree. Expected 'Test 3', instead %s", v)
	}

	v = tree.left.left.value
	if v != "Test 2" {
		t.Errorf("Failed to add item to AVL Tree. Expected 'Test 2', instead %s", v)
	}
}

func TestAVLTreeGetHeight(t *testing.T) {
	tree := NewAVLTree()
	if getHeight(tree) != 0 {
		t.Errorf("Failed to test getHeight. Expected 0, instead %d", getHeight(tree))
	}

	tree.Add(100, "Test 100")
	if getHeight(tree) != 1 {
		t.Errorf("Failed to test getHeight. Expected 0, instead %d", getHeight(tree))
	}
}

func TestAVLTreeTurnRight(t *testing.T) {
	tree := NewAVLTree()

	tree.Add(5, "Test 5")
	tree.Add(4, "Test 4")
	tree.Add(6, "Test 6")

	tree.turnRight()

	v := tree.value
	if v != "Test 4" {
		t.Errorf("Failed to turn right. Expected root to be 'Test 4', instead %s", v)
	}

	v = tree.right.value
	if v != "Test 5" {
		t.Errorf("Failed to turn right. Expected root to be 'Test 5', instead %s", v)
	}

	v = tree.right.right.value
	if v != "Test 6" {
		t.Errorf("Failed to turn right. Expected root to be 'Test 6', instead %s", v)
	}

	if tree.left != nil {
		t.Errorf("Failed to turn right. Expected root left child to be nil, instead %d", tree.left)
	}
}

func TestAVLTreeTurnLeft(t *testing.T) {
	tree := NewAVLTree()

	tree.Add(5, "Test 5")
	tree.Add(4, "Test 4")
	tree.Add(6, "Test 6")

	tree.turnLeft()

	v := tree.value
	if v != "Test 6" {
		t.Errorf("Failed to turn left. Expected root to be 'Test 6', instead %s", v)
	}

	v = tree.left.value
	if v != "Test 5" {
		t.Errorf("Failed to turn left. Expected root to be 'Test 5', instead %s", v)
	}

	v = tree.left.left.value
	if v != "Test 4" {
		t.Errorf("Failed to turn left. Expected root to be 'Test 4', instead %s", v)
	}

	if tree.right != nil {
		t.Errorf("Failed to turn left. Expected root left child to be nil, instead %d", tree.left)
	}
}

func TestAVLTreeBalanceFactor(t *testing.T) {
	tree := NewAVLTree()

	tree.Add(5, "Test 5")
	v := tree.balanceFactor()
	if v != 0 {
		t.Errorf("Failed to test BalanceFactor. Expected 0, instead %d", v)
	}

	tree.Add(4, "Test 4")
	v = tree.balanceFactor()
	if v != -1 {
		t.Errorf("Failed to test BalanceFactor. Expected 0, instead %d", v)
	}

	tree.Add(3, "Test 3")
	v = tree.balanceFactor()
	if v != 0 {
		t.Errorf("Failed to test BalanceFactor. Expected 0, instead %d", v)
	}

	tree.Add(1, "Test 1")
	v = tree.balanceFactor()
	if v != -1 {
		t.Errorf("Failed to test BalanceFactor. Expected 0, instead %d", v)
	}

	tree.Add(10, "Test 10")
	v = tree.balanceFactor()
	if v != 0 {
		t.Errorf("Failed to test BalanceFactor. Expected 0, instead %d", v)
	}
}

func TestAVLTreeGet(t *testing.T) {
	tree := NewAVLTree()

	tree.Add(50, "Test 50")
	tree.Add(40, "Test 40")
	tree.Add(30, "Test 30")
	tree.Add(20, "Test 20")
	tree.Add(5, "Test 5")
	tree.Add(4, "Test 4")
	tree.Add(3, "Test 3")
	tree.Add(2, "Test 2")

	v := tree.Get(5)
	if v != "Test 5" {
		t.Errorf("Failed to get an element. Expected 'Test 5', instead %s", v)
	}

	v = tree.Get(20)
	if v != "Test 20" {
		t.Errorf("Failed to get an element. Expected 'Test 20', instead %s", v)
	}

	v = tree.Get(3)
	if v != "Test 3" {
		t.Errorf("Failed to get an element. Expected 'Test 3', instead %s", v)
	}

	v = tree.Get(100)
	if v != nil {
		t.Errorf("Failed to get an element. Expected 'nil', instead %s", v)
	}
}
