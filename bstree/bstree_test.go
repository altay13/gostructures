package bstree

import (
	"testing"
)

func TestBSTreePutGet(t *testing.T) {
	b := NewBSTree()

	b.Put(5, "Item 5")
	b.Put(7, "Item 7")
	b.Put(1, "Item 1")
	b.Put(14, "Item 14")
	b.Put(10, "Item 10")

	it := b.Get(14)
	if it != "Item 14" {
		t.Errorf("Failed to put item. Expected 'Item 14', instead %s", it)
	}

	it = b.Get(10)
	if it != "Item 10" {
		t.Errorf("Failed to put item. Expected 'Item 10', instead %s", it)
	}

	it = b.Get(100)
	if it != nil {
		t.Errorf("Failed to put item. Expected 'nil', instead %s", it)
	}

}

func TestBSTReeRemove(t *testing.T) {
	b := NewBSTree()

	b.Put(5, "Item 5")
	b.Put(7, "Item 7")
	b.Put(2, "Item 2")
	b.Put(1, "Item 1")
	b.Put(14, "Item 14")
	b.Put(10, "Item 10")
	b.Put(11, "Item 11")
	b.Put(9, "Item 9")
	b.Put(6, "Item 6")

	b.Remove(9)
	if b.Get(9) != nil {
		t.Errorf("Failed to remove item. Expected 'nil', instead %s", b.Get(9))
	}

	b.Put(9, "Item 9")

	b.Remove(2)
	if b.left.value != "Item 1" {
		t.Errorf("Failed to remove item. Expected 'Item 1', instead %s", b.left.value)
	}

	b.Remove(14)
	if b.right.right.value != "Item 10" {
		t.Errorf("Failed to remove item. Expected 'Item 10', instead %s", b.right.right.value)
	}

	b.Remove(7)
	if b.right.value != "Item 9" {
		t.Errorf("Failed to remove item. Expected 'Item 9', instead %s", b.right.value)
	}

	node := b.get(9)
	if node.left.value != nil {
		t.Errorf("Failed to remove item. Expected 'nil', instead %s", node.left.value)
	}
	if node.right.value != "Item 10" {
		t.Errorf("Failed to remove item. Expected 'Item 10', instead %s", node.right.value)
	}
}

func TestBSTreeGetSortedArray(t *testing.T) {
	b := NewBSTree()

	b.Put(5, 5)
	b.Put(7, 7)
	b.Put(1, 1)
	b.Put(14, 14)
	b.Put(10, 10)
	b.Put(18, 18)
	b.Put(3, 3)

	var a []interface{}
	b.GetSortedArray(&a)

	res := [7]int{1, 3, 5, 7, 10, 14, 18}

	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Errorf("Failed to sort. %d", a)
		}
	}
}

func TestBSTreeGetSortedArrayOfKeys(t *testing.T) {
	b := NewBSTree()

	b.Put("A", 5.0)
	b.Put("D", 5.5)
	b.Put("C", 1.2)
	b.Put("F", 14.4)
	b.Put("E", 13.9)
	b.Put("G", 18.0)
	b.Put("B", 3.1)

	var a []interface{}
	b.GetSortedArrayOfKeys(&a)

	res := [7]string{"A", "B", "C", "D", "E", "F", "G"}

	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Errorf("Failed to sort. %d", a)
		}
	}
}

func TestBSTreeContains(t *testing.T) {
	b := NewBSTree()

	b.Put(5, 5)
	b.Put(7, 7)
	b.Put(1, 1)
	b.Put(14, 14)
	b.Put(10, 10)
	b.Put(18, 18)
	b.Put(3, 3)

	if !b.Contains(5) {
		t.Errorf("Failed to locate '5' key. Contains method fails.")
	}

	if !b.Contains(18) {
		t.Errorf("Failed to locate '18' key. Contains method fails.")
	}

	if b.Contains(100) {
		t.Errorf("Failed to locate '100' key. Contains method fails. Should not find the '100' key.")
	}
}
