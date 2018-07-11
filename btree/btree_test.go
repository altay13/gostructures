package btree

import "testing"

func TestBTree_InsertNonFull_ToLeaf(t *testing.T) {
	tree := NewBTree(3)

	k := &item{
		key: 1,
		value: "Item 1",
	}
	tree.insertNonFull(*k)

	if tree.keys[0].value != "Item 1" {
		t.Errorf("Failed to insertNonFull to root leaf.")
	}
	if tree.cnt != 1 {
		t.Errorf("Failed to insertNonFull to root leaf. Expected cnt = 1, instead cnt = %d", tree.cnt)
	}

	k = &item{
		key: 2,
		value: "Item 2",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 3,
		value: "Item 3",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 4,
		value: "Item 4",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 5,
		value: "Item 5",
	}
	tree.insertNonFull(*k)

	if tree.keys[0].value != "Item 1" {
		t.Errorf("Failed to insertNonFull to root leaf.")
	}
	if tree.keys[4].value != "Item 5" {
		t.Errorf("Failed to insertNonFull to root leaf.")
	}
	if tree.cnt != 5 {
		t.Errorf("Failed to insertNonFull to root leaf. Expected cnt = 1, instead cnt = %d", tree.cnt)
	}
}

func TestBTree_IsKeysFull(t *testing.T) {
	tree := NewBTree(3)

	item_1 := &item{
		key: 1,
		value: "Item 1",
	}

	tree.keys[0] = *item_1
	tree.cnt++

	item_2 := &item{
		key: 2,
		value: "Item 2",
	}

	tree.keys[1] = *item_2
	tree.cnt++

	item_3 := &item{
		key: 3,
		value: "Item 3",
	}

	tree.keys[2] = *item_3
	tree.cnt++

	item_4 := &item{
		key: 4,
		value: "Item 4",
	}

	tree.keys[3] = *item_4
	tree.cnt++

	if tree.isKeysFull() {
		t.Errorf("isKeyFull method failed. Expecting true, instead %t ", tree.isKeysFull())
	}

	item_5 := &item{
		key: 5,
		value: "Item 5",
	}

	tree.keys[4] = *item_5
	tree.cnt++

	if !tree.isKeysFull() {
		t.Errorf("isKeyFull method failed. Expecting true, instead %t ", tree.isKeysFull())
	}
}

func TestBTree_SplitNode_DepthOne(t *testing.T) {
	tree := getTestTreeWithOneDepth(t)
	tree.splitChild(tree.children[2], 2)

	if tree.cnt != 5 {
		t.Errorf("Failed to SplitNode. Expecting parent tree count be 5, instead %d", tree.cnt)
	}
	if tree.children[2].cnt != 2 {
		t.Errorf("Failed to SplitNode. Expecting parent tree count be 2, instead %d", tree.children[2].cnt)
	}
	if tree.children[3].cnt != 2 {
		t.Errorf("Failed to SplitNode. Expecting parent tree count be 2, instead %d", tree.children[3].cnt)
	}

	if tree.keys[2].key != 23 {
		t.Errorf("Failed to SplitNode. Expecting element to be 23, instead %d", tree.keys[2].key)
	}

	if tree.children[2].keys[0].key != 21 {
		t.Errorf("Failed to SplitNode. Expecting element to be 21, instead %d", tree.children[2].keys[0].key)
	}
	if tree.children[2].keys[1].key != 22 {
		t.Errorf("Failed to SplitNode. Expecting element to be 22, instead %d", tree.children[2].keys[1].key)
	}

	if tree.children[3].keys[0].key != 24 {
		t.Errorf("Failed to SplitNode. Expecting element to be 24, instead %d", tree.children[3].keys[0].key)
	}
	if tree.children[3].keys[1].key != 25 {
		t.Errorf("Failed to SplitNode. Expecting element to be 25, instead %d", tree.children[2].keys[1].key)
	}
}

func TestBTree_SplitNode_DepthTwo(t *testing.T) {
	tree := getTestTreeWithTwoDepth(t)
	tree.splitChild(tree.children[2], 2)

	if tree.cnt != 5 {
		t.Errorf("Failed to SplitNode. Expecting parent tree count be 5, instead %d", tree.cnt)
	}
	if tree.children[2].cnt != 2 {
		t.Errorf("Failed to SplitNode. Expecting parent tree count be 2, instead %d", tree.children[2].cnt)
	}
	if tree.children[3].cnt != 2 {
		t.Errorf("Failed to SplitNode. Expecting parent tree count be 2, instead %d", tree.children[3].cnt)
	}

	if tree.keys[2].key != 23 {
		t.Errorf("Failed to SplitNode. Expecting element to be 23, instead %d", tree.keys[2].key)
	}

	if tree.children[2].keys[0].key != 21 {
		t.Errorf("Failed to SplitNode. Expecting element to be 21, instead %d", tree.children[2].keys[0].key)
	}
	if tree.children[2].keys[1].key != 22 {
		t.Errorf("Failed to SplitNode. Expecting element to be 22, instead %d", tree.children[2].keys[1].key)
	}

	if tree.children[3].keys[0].key != 28 {
		t.Errorf("Failed to SplitNode. Expecting element to be 28, instead %d", tree.children[3].keys[0].key)
	}
	if tree.children[3].keys[1].key != 29 {
		t.Errorf("Failed to SplitNode. Expecting element to be 29, instead %d", tree.children[2].keys[1].key)
	}

	if tree.children[3].children[0].keys[0].key != 24 {
		t.Errorf("Failed to SplitNode. Expecting element to be 24, instead %d", tree.children[3].children[0].keys[0].key)
	}
	if tree.children[3].children[0].keys[1].key != 27 {
		t.Errorf("Failed to SplitNode. Expecting element to be 27, instead %d", tree.children[3].children[0].keys[3].key)
	}
}

func TestBTree_InsertNonFull(t *testing.T) {
	tree := getTestTreeWithTwoDepth(t)
	k := item{
		key: 25,
		value: "Item 25",
	}
	tree.insertNonFull(k)

	k = item{
		key: 26,
		value: "Item 26",
	}
	tree.insertNonFull(k)

	if tree.children[2].cnt != 2 {
		t.Errorf("Failed to InsertNonFull. Expecting count to be 2, instead %d", tree.children[2].cnt)
	}
	if tree.children[3].cnt != 2 {
		t.Errorf("Failed to InsertNonFull. Expecting count to be 2, instead %d", tree.children[3].cnt)
	}
	if tree.children[2].keys[0].key != 21 {
		t.Errorf("Failed to InsertNonFull. Expecting key to be 21, instead %d", tree.children[2].keys[0].key)
	}
	if tree.children[3].keys[0].key != 28 {
		t.Errorf("Failed to InsertNonFull. Expecting key to be 28, instead %d", tree.children[3].keys[0].key)
	}

	if tree.children[3].children[0].cnt != 4 {
		t.Errorf("Failed to InsertNonFull. Expecting count to be 4, instead %d", tree.children[3].children[0].cnt)
	}
	if tree.children[3].children[0].keys[0].key != 24 {
		t.Errorf("Failed to InsertNonFull. Expecting key to be 24, instead %d", tree.children[3].children[0].keys[0].key)
	}
	if tree.children[3].children[0].keys[3].key != 27 {
		t.Errorf("Failed to InsertNonFull. Expecting key to be 27, instead %d", tree.children[3].children[0].keys[3].key)
	}
}

func TestBTree_Add_CustomCase(t *testing.T) {
	tree := getTestTreeWithTwoDepth(t)

	tree.Add(25, "New Item 25")

	if tree.children[3].children[0].keys[1].key != 25 {
		t.Errorf("Failed to Add(). Expecting key to be 25, instead %d", tree.children[3].children[0].keys[1].key)
	}
	if tree.children[2].cnt != 2 {
		t.Errorf("Failed to Add(). Expecting count to be 2, instead %d", tree.children[2].cnt)
	}
	if tree.children[3].cnt != 2 {
		t.Errorf("Failed to Add(). Expecting count to be 2, instead %d", tree.children[3].cnt)
	}

	tree.Add(26, "New Item 26")

	if tree.cnt != 1 {
		t.Errorf("Failed to Add(). Expecting root keys count to be 1, instead %d", tree.cnt)
	}
	if tree.keys[0].key != 23 {
		t.Errorf("Failed to Add(). Expecting root key to be 23, instead %d", tree.keys[0].key)
	}
	if tree.children[0].keys[0].key != 10 {
		t.Errorf("Failed to Add(). Expecting root key to be 10, instead %d", tree.children[0].keys[0].key)
	}
	if tree.children[1].keys[0].key != 30 {
		t.Errorf("Failed to Add(). Expecting root key to be 30, instead %d", tree.children[1].keys[0].key)
	}

	if tree.children[1].children[0].children[0].keys[2].key != 26 {
		t.Errorf("Failed to Add(). Expecting key to be 26, instead %d", tree.children[1].children[0].children[0].keys[2])
	}
}

func TestBTree_Add(t *testing.T) {
	tree := NewBTree(3)

	tree.Add(10, "Item 10")
	tree.Add(20, "Item 20")
	tree.Add(30, "Item 30")
	tree.Add(40, "Item 40")

	if tree.cnt != 4 {
		t.Errorf("Failed to Add(). Expecting count to be 4, instead %d", tree.cnt)
	}

	tree.Add(50, "Item 50")
	if tree.cnt != 5 {
		t.Errorf("Failed to Add(). Expecting count to be 5, instead %d", tree.cnt)
	}

	tree.Add(25, "Item 25")
	if tree.cnt != 1 {
		t.Errorf("Failed to Add(). Expecting count to be 1, instead %d", tree.cnt)
	}
	if tree.children[0].keys[2].key != 25 {
		t.Errorf("Failed to Add(). Expecting key to be 25, instead %d", tree.children[0].keys[2].key)
	}

	tree.Add(31, "Item 31")
	if tree.cnt != 1 {
		t.Errorf("Failed to Add(). Expecting count to be 1, instead %d", tree.cnt)
	}
	if tree.children[1].keys[0].key != 31 {
		t.Errorf("Failed to Add(). Expecting key to be 31, instead %d", tree.children[1].keys[0].key)
	}

	tree.Add(29, "Item 29")
	if tree.cnt != 1 {
		t.Errorf("Failed to Add(). Expecting count to be 1, instead %d", tree.cnt)
	}
	if tree.children[0].keys[3].key != 29 {
		t.Errorf("Failed to Add(). Expecting key to be 31, instead %d", tree.children[0].keys[3].key)
	}
}

func TestBTree_Get(t *testing.T) {
	tree := getTestTreeWithTwoDepth(t)

	v := tree.Get(23)
	if v != "child 23" {
		t.Errorf("Failed to Get an element. Expect value to be 'child 23', instead %s", v)
	}

	v = tree.Get(28)
	if v != "child 28" {
		t.Errorf("Failed to Get an element. Expect value to be 'child 28', instead %s", v)
	}

	v = tree.Get(40)
	if v != "Item 40" {
		t.Errorf("Failed to Get an element. Expect value to be 'child 40', instead %s", v)
	}

	v = tree.Get(27)
	if v != "child_child 27" {
		t.Errorf("Failed to Get an element. Expect value to be 'child 27', instead %s", v)
	}

	v = tree.Get(100)
	if v != nil {
		t.Errorf("Failed to Get an element. Expect value to be 'nil', instead %s", v)
	}
}

// building test tree with 1 child node.
//	 	10 20   30 40
//	 	      |
//	 	      |
//  21  22  23  24  25
func getTestTreeWithOneDepth(t *testing.T) *BTree {
	tree := NewBTree(3)

	k := &item{
		key: 10,
		value: "Item 10",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 20,
		value: "Item 20",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 30,
		value: "Item 30",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 40,
		value: "Item 40",
	}
	tree.insertNonFull(*k)

	tree.leaf = false

	child := NewBTree(3)
	k = &item{
		key: 21,
		value: "child 21",
	}
	child.insertNonFull(*k)

	k = &item{
		key: 22,
		value: "child 22",
	}
	child.insertNonFull(*k)

	k = &item{
		key: 23,
		value: "child 23",
	}
	child.insertNonFull(*k)

	k = &item{
		key: 24,
		value: "child 24",
	}
	child.insertNonFull(*k)

	k = &item{
		key: 25,
		value: "child 25",
	}
	child.insertNonFull(*k)

	tree.children[2] = child

	return tree
}

// building test tree with 1 child node.
//	 	10 20   30 40
//	 	      |
//	 	      |
//  21  22  23   28  29
//             |
//			   |
//			 24 27
func getTestTreeWithTwoDepth(t *testing.T) *BTree {
	tree := NewBTree(3)

	k := &item{
		key: 10,
		value: "Item 10",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 20,
		value: "Item 20",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 30,
		value: "Item 30",
	}
	tree.insertNonFull(*k)

	k = &item{
		key: 40,
		value: "Item 40",
	}
	tree.insertNonFull(*k)

	tree.leaf = false

	child := NewBTree(3)
	k = &item{
		key: 21,
		value: "child 21",
	}
	child.insertNonFull(*k)

	k = &item{
		key: 22,
		value: "child 22",
	}
	child.insertNonFull(*k)

	k = &item{
		key: 23,
		value: "child 23",
	}
	child.insertNonFull(*k)

	k = &item{
		key: 28,
		value: "child 28",
	}
	child.insertNonFull(*k)

	k = &item{
		key: 29,
		value: "child 29",
	}
	child.insertNonFull(*k)

	tree.children[2] = child

	child.leaf = false

	childOfChild := NewBTree(3)
	k = &item{
		key: 24,
		value: "child_child 24",
	}
	childOfChild.insertNonFull(*k)

	k = &item{
		key: 27,
		value: "child_child 27",
	}
	childOfChild.insertNonFull(*k)

	child.children[3] = childOfChild

	return tree
}
