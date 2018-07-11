package trie

import (
	"github.com/altay13/gostructures/hashtable"
)

// Trie ...
type Trie struct {
	children *hashtable.Hash
	value    string
	hashSize int
}

// NewTrie creates a root node and returns *Trie object
func NewTrie(hash *hashtable.Hash, hashSize int) *Trie {
	trie := &Trie{
		value:    "",
		hashSize: hashSize,
	}
	trie.children = hash

	return trie
}

// Insert method inserts string into the tree
func (t *Trie) Insert(str string) error {
	node := t
	for _, ch := range str {
		raw := node.children.Get(ch)
		if raw != nil {
			node = raw.(*Trie)
			continue
		} else {
			hash := hashtable.NewHash(node.hashSize)
			new := NewTrie(hash, node.hashSize)
			node.children.Add(ch, new)
			node = new
		}
	}

	node.value = str

	return nil
}

// Find method searches for the string in tree
// Returns the string if it is in tree
func (t *Trie) Find(str string) (string, bool) {

	node := t
	for _, ch := range str {
		raw := node.children.Get(ch)
		if raw != nil {
			node = raw.(*Trie)
			continue
		} else {
			return "", false
		}
	}

	if len(node.value) > 0 {
		return node.value, true
	}

	return "", false
}
