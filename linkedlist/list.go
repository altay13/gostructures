package linkedlist

import "fmt"

// List implementation
type List struct {
	top  *Node
	last *Node
	size int
}

// Node ...
type Node struct {
	item interface{}
	next *Node
}

// NewList method creates the stack
func NewList() *List {
	l := &List{
		size: 0,
	}
	return l
}

// AddNode adds node to the end of the list
func (l *List) AddNode(item interface{}) {
	n := &Node{
		item: item,
		next: nil,
	}

	if l.size == 0 {
		l.top = n
		l.last = n
	} else {
		l.last.next = n
		l.last = n
	}

	l.size++
}

// GetNode ...
func (l *List) GetNode(item interface{}) *Node {
	for n := l.top; ; n = n.next {
		if item == n.item {
			return n
		}

		if n.next == nil {
			break
		}
	}

	return nil
}

// InsertAfter ...
func (l *List) InsertAfter(node interface{}, item interface{}) error {
	n := &Node{
		item: item,
		next: nil,
	}

	switch node.(type) {
	case *Node:
		n.next = node.(*Node).next
		node.(*Node).next = n
		l.size++
	default:
		if fv := l.GetNode(node); fv != nil {
			n.next = fv.next
			fv.next = n
			l.size++
		} else {
			return fmt.Errorf("[ ERR ] --> Failed to add item: Failed to allocate the node")
		}
	}

	if n.next == nil {
		l.last = n
	}

	return nil
}

// InsertFirst ...
func (l *List) InsertFirst(item interface{}) {
	n := &Node{
		item: item,
		next: nil,
	}
	n.next = l.top
	l.top = n
	l.size++
}

// GetLastNode ...
func (l *List) GetLastNode() *Node {
	return l.last
}
