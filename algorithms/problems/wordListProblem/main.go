package main

import (
	"fmt"

	"github.com/altay13/gostructures/hashtable"
)

type Trie struct {
	children *hashtable.Hash
	value    string
	hashSize int
}

func NewTrie(hash *hashtable.Hash, hashSize int) *Trie {
	t := &Trie{
		children: hash,
		value:    "",
		hashSize: hashSize,
	}
	return t
}

func (t *Trie) Insert(str string) {
	node := t
	for _, v := range str {
		raw := node.children.Get(v)
		if raw != nil {
			node = raw.(*Trie)
		} else {
			// adding a node
			hash := hashtable.NewHash(t.hashSize)
			tmp := NewTrie(hash, t.hashSize)
			node.children.Add(v, tmp)
			node = tmp
		}
	}
	node.value = str
}

func (t *Trie) Find(str string) (string, bool) {
	node := t
	for _, v := range str {
		raw := node.children.Get(v)
		if raw != nil {
			node = raw.(*Trie)
		} else {
			return "", false
		}
		if len(node.value) > 0 {
			return node.value, true
		}
	}

	if len(node.value) > 0 {
		return node.value, true
	}

	return "", false
}

func main() {
	str := []string{"rockstar", "rock", "star", "rocks", "tar", "star", "rockstars", "super", "highway", "high", "way", "superhighway"}

	hash := hashtable.NewHash(len(str))
	trie := NewTrie(hash, len(str))

	for _, word := range str {
		trie.Insert(word)
	}

	result := make([]string, 0)
	var tmp []string

	for _, word := range str {
		tmp = make([]string, 0)
		for len(word) != 0 {
			substr, ok := trie.Find(word)
			if ok {
				tmp = append(tmp, substr)
				word = word[len(substr):]
			} else {
				tmp = nil
				word = ""
			}
		}

		if len(tmp) <= 1 {
			continue
		}

		for _, v := range tmp {
			result = append(result, v)
		}
	}

	fmt.Println(result)
}
