package avltree

// AVLTree ...
type AVLTree struct {
	key   uint32
	value interface{}

	height int8

	left  *AVLTree
	right *AVLTree

	parent *AVLTree
}

// NewAVLTree ...
func NewAVLTree() *AVLTree {
	tree := &AVLTree{
		height: 0,
		parent: nil,
		left:   nil,
		right:  nil,
	}
	return tree
}

// Add ...
func (tree *AVLTree) Add(key uint32, value interface{}) {
	if tree.height == 0 {
		tree.height = 1
		tree.key = key
		tree.value = value

		return
	}

	tree.add(key, value)
}

func (tree *AVLTree) add(key uint32, value interface{}) {
	if tree.key > key {
		if tree.left == nil {
			new := &AVLTree{
				key:    key,
				value:  value,
				height: 1,
				parent: tree,
				left:   nil,
				right:  nil,
			}
			tree.left = new
		} else {
			tree.left.add(key, value)
		}
	} else if tree.key < key {
		if tree.right == nil {
			new := &AVLTree{
				key:    key,
				value:  value,
				height: 1,
				parent: tree,
				left:   nil,
				right:  nil,
			}
			tree.right = new
		} else {
			tree.right.add(key, value)
		}
	} else {
		return
	}

	tree.rebalance()
}

// Get ...
func (tree *AVLTree) Get(key uint32) interface{} {
	t := tree.get(key)
	if t != nil {
		return t.value
	}

	return nil
}

func (tree *AVLTree) get(key uint32) *AVLTree {
	var t *AVLTree
	if tree.key > key {
		t = tree.left
	} else if tree.key < key {
		t = tree.right
	} else {
		return tree
	}

	if t == nil {
		return nil
	}

	return t.get(key)
}

func (tree *AVLTree) rebalance() {
	tree.fixHeight()

	if tree.balanceFactor() == 2 {
		// right is bigger
		if tree.right.balanceFactor() < 0 {
			tree.right.turnRight()
		}
		tree.turnLeft()
	} else if tree.balanceFactor() == -2 {
		// left is bigger
		if tree.left.balanceFactor() > 0 {
			tree.left.turnLeft()
		}
		tree.turnRight()
	}
}

func getHeight(tree *AVLTree) int8 {
	if tree != nil {
		return tree.height
	}
	return 0
}

func (tree *AVLTree) balanceFactor() int8 {
	return getHeight(tree.right) - getHeight(tree.left)
}

func (tree *AVLTree) fixHeight() {
	hl := getHeight(tree.left)
	hr := getHeight(tree.right)

	if hl > hr {
		tree.height = hl + 1
	} else {
		tree.height = hr + 1
	}
}

func (tree *AVLTree) turnLeft() {
	var temp AVLTree
	temp = *tree

	rTree := temp.right
	clTree := rTree.left

	rTree.left = &temp
	rTree.parent = temp.parent

	temp.parent = rTree
	temp.right = clTree

	*tree = *rTree

	tree.left.fixHeight()
	tree.fixHeight()
}

func (tree *AVLTree) turnRight() {
	var temp AVLTree
	temp = *tree

	lTree := temp.left
	lrTree := lTree.right

	lTree.right = &temp
	lTree.parent = temp.parent

	temp.parent = lTree
	temp.left = lrTree

	*tree = *lTree

	tree.right.fixHeight()
	tree.fixHeight()
}
