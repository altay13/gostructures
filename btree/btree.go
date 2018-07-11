package btree

// BTree ...
type BTree struct {
	t   int
	cnt int

	keys     []item
	children []*BTree

	parent *BTree

	leaf bool
}

type item struct {
	key   int32
	value interface{}
}

// NewBTree ...
func NewBTree(t int) *BTree {
	tree := &BTree{
		t:        t,
		keys:     make([]item, 2*t-1),
		children: make([]*BTree, 2*t),
		cnt:      0,
		leaf:     true,
	}

	return tree
}

// Add ...
func (tree *BTree) Add(key int32, value interface{}) {
	k := item {
		key:   key,
		value: value,
	}

	if tree.cnt == 0 {
		tree.insertNonFull(k)
		return
	}

	tree.add(k)
}

func (tree *BTree) add(k item) {
	if tree.isKeysFull() {
		newRoot := &BTree{
			t:        tree.t,
			keys:     make([]item, 2*tree.t-1),
			children: make([]*BTree, 2*tree.t),
			cnt:      0,
			leaf: 	  false,
		}
		tmp := *tree
		newRoot.children[0] = &tmp

		newRoot.splitChild(newRoot.children[0], 0)

		i := 0
		if newRoot.keys[i].key < k.key {
			i++
		}

		newRoot.children[i].insertNonFull(k)

		*tree = *newRoot
	} else {
		tree.insertNonFull(k)
	}
}

// Get ...
func (tree *BTree) Get(key int32) interface{} {
	return tree.get(key).value
}

func (tree *BTree) get(k int32) *item {
	i := 0

	for ; i < tree.cnt; i++ {
		if k <= tree.keys[i].key {
			break
		}
	}

	if i < tree.cnt && k == tree.keys[i].key {
		return &tree.keys[i]
	}

	if tree.leaf {
		return &item{}
	}

	if tree.children[i] != nil {
		return tree.children[i].get(k)
	} else {
		return &item{}
	}
}

func (tree *BTree) isKeysFull() bool {
	if tree.cnt == 2*tree.t-1 {
		return true
	}
	return false
}

func (tree *BTree) insertNonFull(k item) {
	i := tree.cnt - 1

	if tree.leaf {
		for ; i >= 0 && tree.keys[i].key > k.key; i-- {
			tree.keys[i+1] = tree.keys[i]
		}
		tree.keys[i+1] = k
		tree.cnt++
	} else {
		for ; i >= 0 && tree.keys[i].key > k.key; i-- {
		}

		i++
		if tree.children[i].isKeysFull() {
			// split tree.children[i+1] node
			tree.splitChild(tree.children[i], i)

			if tree.keys[i].key < k.key {
				i++
			}
		}

		tree.children[i].insertNonFull(k)
	}
}

func (tree *BTree) splitChild(child *BTree, n int) {
	nTree := &BTree{
		t:        tree.t,
		keys:     make([]item, 2*tree.t-1),
		children: make([]*BTree, 2*tree.t),
		cnt:      0,
		leaf:	  true,
	}
	j := child.getMedian()

	outK := child.keys[j]
	child.keys[j] = item{}
	child.cnt--
	for i := 0; i < j; i++ {
		nTree.keys[i] = child.keys[i+(j+1)]
		nTree.cnt++

		child.keys[i+(j+1)]=item{}
		child.cnt--
	}

	if !child.leaf {
		nTree.leaf = false
		for i := 0; i <= j; i++ {
			nTree.children[i] = child.children[i+(j+1)]
			child.children[i+(j+1)] = nil
		}
	}

	for i := tree.cnt-1; i >= n; i-- {
		tree.keys[i+1] = tree.keys[i]
		tree.children[(i+1)+1] = tree.children[i+1]
	}

	tree.keys[n] = outK
	tree.children[n+1] = nTree
	tree.cnt++
}

func (tree *BTree) getMedian() int {
	return tree.cnt / 2
}