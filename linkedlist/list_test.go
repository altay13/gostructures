package linkedlist

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	l := NewList()

	l.AddNode("Node_1")

	if l.size != 1 {
		t.Errorf("The size of the List is wrong. Expecting 1, instead %d", l.size)
	}

	l.AddNode("Node_2")

	if l.size != 2 {
		t.Errorf("The size of the List is wrong. Expecting 1, instead %d", l.size)
	}

	if l.last.item != "Node_2" {
		t.Errorf("Failed to save the last node. Expecting 'Node_2' %s", l.last.item)
	}

	if l.top.item != "Node_1" {
		t.Errorf("Failed to add the node. Expecting 'Node_1' %s", l.top.item)
	}

	n := l.top.next
	if n.item != "Node_2" {
		t.Errorf("Failed to add the node. Expecting 'Node_2' %s", n.item)
	}

	l.AddNode("Node_3")

	nn := n.next
	if nn.item != "Node_3" {
		t.Errorf("Failed to add the node. Expecting 'Node_3' %s", nn.item)
	}

}

func TestGetNode(t *testing.T) {
	l := NewList()

	l.AddNode("Node_1")
	l.AddNode("Node_2")
	l.AddNode("Node_3")
	l.AddNode("Node_4")

	n := l.GetNode("Node_3")
	if n == nil || n.item != "Node_3" {
		t.Errorf("Failed to get a node")
	}

	if n.next.item != "Node_4" {
		t.Errorf("Wrong next Node")
	}

	n = nil

	n = l.GetNode("Node_4")
	if n == nil || n.item != "Node_4" {
		t.Errorf("Failed to get a node")
	}
}

func TestInsertAfter(t *testing.T) {
	l := NewList()

	l.AddNode("Node_1")
	l.AddNode("Node_2")
	l.AddNode("Node_3")
	l.AddNode("Node_4")

	l.InsertAfter("Node_2", "Node_2.5")
	if l.size != 5 {
		t.Errorf("Failed to insert node after.")
	}

	n := l.GetNode("Node_2")
	if n.next.item != "Node_2.5" {
		t.Errorf("Failed to insert node after.")
	}

	n = nil
	n = l.GetNode("Node_2.5")
	if n.next.item != "Node_3" {
		t.Errorf("Failed to insert node after.")
	}

	l.InsertAfter("Node_4", "Node_5")
	if l.last.item != "Node_5" {
		t.Errorf("Failed to insert node after.")
	}

	n = nil
	n = l.GetNode("Node_2.5")
	l.InsertAfter(n, "Node_2.9")
	if l.size != 7 {
		t.Errorf("Failed to insert node after.")
	}
	n = nil
	n = l.GetNode("Node_2.5")
	if n.next.item != "Node_2.9" {
		t.Errorf("Failed to insert node after.")
	}

	n = nil
	n = l.GetNode("Node_999")
	if n != nil {
		t.Errorf("Failed to insert node after.")
	}
}

func TestInsertFirst(t *testing.T) {
	l := NewList()

	l.AddNode("Node_1")
	l.AddNode("Node_2")

	l.InsertFirst("Node_0")

	if l.size != 3 {
		t.Errorf("Failed to insert first.")
	}
	if l.top.item != "Node_0" {
		t.Errorf("Failed to insert first.")
	}
	n := l.GetNode("Node_0")
	if n.next.item != "Node_1" {
		t.Errorf("Failed to insert first.")
	}

	w := NewList()

	w.InsertFirst("item")
	if w.size != 1 {
		t.Errorf("Failed to insert first.")
	}
	if w.top.item != "item" {
		t.Errorf("Failed to insert first.")
	}
}
