package rbtree

import (
	"log"
)

type color bool

const (
	red   color = true
	black color = false
)

var (
	nilNode = &RBTree{0, 0, black, nil, nil, nil}
)

// RBTree ...
type RBTree struct {
	key   uint32
	value interface{}

	color color // true - red, false - black

	left  *RBTree
	right *RBTree

	parent *RBTree
}

// NewRBTree ...
func NewRBTree() *RBTree {
	tree := &RBTree{
		color:  black,
		parent: nil,
		left:   nil,
		right:  nil,
	}
	return tree
}

// Add ...
func (tree *RBTree) Add(key uint32, value interface{}) {
	if tree.left == nil {
		tree.key = key
		tree.value = value

		tree.color = black

		tree.left = nilNode
		tree.right = nilNode

		tree.parent = nil

		return
	}

	tree.addNode(key, value)

	tree.color = black
}

func (tree *RBTree) addNode(key uint32, value interface{}) {
	var cTree *RBTree

	if tree.key > key {
		// move left
		cTree = tree.left
	} else if tree.key < key {
		// move right
		cTree = tree.right
	} else {
		// there is already this key. Later implement error
		return
	}

	if cTree == nilNode {
		newNode := &RBTree{
			key:    key,
			value:  value,
			color:  red,
			parent: tree,
			left:   nilNode,
			right:  nilNode,
		}

		if tree.key > key {
			tree.left = newNode
		} else {
			tree.right = newNode
		}

		if newNode.parent.color == red {
			log.Println("balancing")
			newNode.rebalance()
		}

		return
	}

	cTree.addNode(key, value)
}

// Remove ...
func (tree *RBTree) Remove(key uint32) {
	dNode := tree.get(key)
	log.Println("Removing", dNode)
	var ttmp *RBTree

	if dNode.left == nilNode && dNode.right == nilNode {
		*dNode = RBTree{}
	} else if dNode.left != nilNode && dNode.right == nilNode {
		ttmp = dNode.left
		*dNode = *ttmp
	} else if dNode.left == nilNode && dNode.right != nilNode {
		ttmp = dNode.right
		*dNode = *ttmp
	} else {
		var parent *RBTree
		lNode := dNode.right

		for {
			if lNode.left != nilNode {
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

// Get ...
func (tree *RBTree) Get(key uint32) interface{} {
	t := tree.get(key)
	if t != nil {
		return t.value
	}

	return nil
}

func (tree *RBTree) get(key uint32) *RBTree {
	var t *RBTree
	if tree.key > key {
		t = tree.left
	} else if tree.key < key {
		t = tree.right
	} else {
		return tree
	}

	if t == nilNode {
		return nil
	}

	return t.get(key)
}

func (tree *RBTree) rebalance() {
	var uncle *RBTree
	var grand *RBTree
	parent := tree.parent

	if tree.parent.parent != nil {
		grand = tree.parent.parent
	} else {
		return
	}

	currentPosition := "left"

	if tree == parent.right {
		currentPosition = "right"
	}

	// case 1: unclePosition == right, uncleColor == red, currentPosition = left | right
	// case 2: unclePosition == right, uncleColor == black, currentPosition = right
	// case 3: unclePosition == right, uncleColor == black, currentPosition = left
	// case 4: unclePosition == left, uncleColor == red, currentPosition = right | left
	// case 5: unclePosition == left, uncleColor == black, currentPosition = left
	// case 6: unclePosition == left, uncleColor == black, currentPosition = right

	if parent == grand.right {
		// uncle position is left
		uncle = grand.left
		if uncle.color == black {
			if currentPosition == "left" {
				// turn parent tree to right
				parent.turnRight()
				parent.rebalance()
			} else {
				// change colors
				parent.color = black
				grand.color = red

				// turn grand tree to left
				grand.turnLeft()
			}
		} else {
			// change colors
			parent.color = black
			uncle.color = black
			grand.color = red
		}
	} else {
		// uncle position is right
		uncle = grand.right
		if uncle.color == black {
			if currentPosition == "right" {
				// turn parent tree to left
				parent.turnLeft()
				parent.rebalance()
			} else {
				// change colors
				parent.color = black
				grand.color = red

				// turn grand tree to right
				grand.turnRight()
			}
		} else {
			// change colors
			parent.color = black
			uncle.color = black
			grand.color = red
		}
	}
}

func (tree *RBTree) turnLeft() {
	var temp RBTree
	temp = *tree

	rTree := temp.right
	clTree := rTree.left

	rTree.left = &temp
	rTree.parent = temp.parent

	temp.left = clTree
	temp.parent = rTree
	temp.right = nilNode

	*tree = *rTree
}

func (tree *RBTree) turnRight() {
	var temp RBTree
	temp = *tree

	lTree := temp.left
	crTree := lTree.right

	lTree.right = &temp
	lTree.parent = temp.parent

	temp.right = crTree
	temp.parent = lTree
	temp.left = nilNode

	*tree = *lTree
}

// GetSortedFromLargest ...
func (tree *RBTree) TraverseTreeFromLargest(res *[]interface{}) {
	if tree == nilNode {
		return
	}

	tree.right.TraverseTreeFromLargest(res)

	*res = append(*res, tree.value)

	tree.left.TraverseTreeFromLargest(res)
}