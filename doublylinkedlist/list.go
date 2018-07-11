package doublylinkedlist

import "fmt"

// List implementation
type List struct {
	top  *Node
	last *Node
	size int
}

// Node ...
type Node struct {
	Item interface{}
	next *Node
	prev *Node
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
		Item: item,
		next: nil,
		prev: nil,
	}

	if l.size == 0 {
		l.top = n
		l.last = n
	} else {
		n.prev = l.last
		l.last.next = n
		l.last = n
	}

	l.size++
}

// GetNode ...
func (l *List) GetNode(item interface{}) *Node {
	for n := l.top; ; n = n.next {
		if item == n.Item {
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
		Item: item,
		next: nil,
	}

	switch node.(type) {
	case *Node:
		n.prev = node.(*Node)

		n.next = node.(*Node).next
		node.(*Node).next = n

		node.(*Node).next.prev = n

		l.size++
	default:
		if fv := l.GetNode(node); fv != nil {
			n.prev = fv

			n.next = fv.next
			fv.next = n

			fv.next.prev = n

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
	if l.top != nil {
		n := &Node{
			Item: item,
			next: nil,
			prev: nil,
		}

		l.top.prev = n
		n.next = l.top
		l.top = n
		l.size++
	} else {
		l.AddNode(item)
	}

}

// GetLastNode ...
func (l *List) GetLastNode() *Node {
	return l.last
}

// GetFirstNode ...
func (l *List) GetFirstNode() *Node {
	return l.top
}

// RemoveLastNode ...
func (l *List) RemoveLastNode() {
	if l.last == nil {
		return
	}

	if l.last == l.top {
		l.last = nil
		l.top = nil
	} else {
		n := l.last.prev
		l.last = nil
		l.last = n
	}

	l.size--
}

// RemoveFirstNode ...
func (l *List) RemoveFirstNode() {
	if l.top == nil {
		return
	}

	if l.last == l.top {
		l.last = nil
		l.top = nil
	} else {
		n := l.top.next
		l.top = nil
		l.top = n
	}

	l.size--
}
