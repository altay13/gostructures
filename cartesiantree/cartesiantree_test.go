package cartesiantree

import (
	"testing"
)

func TestBSTreeAddGetSortedArray(t *testing.T) {
	tree := NewCartesianTree()

	tree.Add(5)
	tree.Add(7)
	tree.Add(1)
	tree.Add(14)
	tree.Add(10)
	tree.Add(18)
	tree.Add(3)

	var a []uint32
	tree.GetSortedArray(&a)

	res := [7]uint32{1, 3, 5, 7, 10, 14, 18}

	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Errorf("Failed to sort.")
		}
	}
}

func TestBSTreeRemove(t *testing.T) {
	tree := NewCartesianTree()

	tree.Add(5)
	tree.Add(7)
	tree.Add(1)
	tree.Add(14)
	tree.Add(10)
	tree.Add(18)
	tree.Add(3)

	var a []uint32
	tree.GetSortedArray(&a)

	res := [7]uint32{1, 3, 5, 7, 10, 14, 18}

	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Errorf("Failed to sort. %d", a)
		}
	}

	tree.Remove(3)

	a = nil
	tree.GetSortedArray(&a)

	res2 := [6]uint32{1, 5, 7, 10, 14, 18}
	for i := 0; i < len(res2); i++ {
		if res2[i] != a[i] {
			t.Errorf("Failed to sort. %d", a)
		}
	}
}
