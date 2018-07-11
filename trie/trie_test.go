package trie

import (
	"testing"

	"github.com/altay13/gostructures/hashtable"
)

func TestNewTrie(t *testing.T) {
	hash := hashtable.NewHash(26)
	trie := NewTrie(hash, 26)

	if trie.children.Size() != 0 {
		t.Fatalf("Failed to create a new Trie. Expecting size = 0, instead %d", trie.children.Size())
	}

	if len(trie.value) != 0 {
		t.Fatalf("Failed to create a new Trie. Expecting value to be empty")
	}
}

func TestInsert(t *testing.T) {
	hash := hashtable.NewHash(26)
	trie := NewTrie(hash, 26)

	trie.Insert("abc")

	if trie.children.Size() != 1 {
		t.Fatalf("Failed to Insert. Expecting 0, instead %d", trie.children.Size())
	}

	if trie.children.Get('a') == nil {
		t.Fatalf("Failed to Insert. Expecting not nil, instead nil")
	}
	if len(trie.value) > 0 {
		t.Fatalf("Failed to Insert. Expecting 0, instead %d", len(trie.value))
	}

	a := trie.children.Get('a').(*Trie)
	if a.children.Get('b') == nil {
		t.Fatalf("Failed to Insert. Expecting not nil, instead nil")
	}
	if len(trie.value) > 0 {
		t.Fatalf("Failed to Insert. Expecting 0, instead %d", len(trie.value))
	}

	b := a.children.Get('b').(*Trie)
	if b.children.Get('c') == nil {
		t.Fatalf("Failed to Insert. Expecting not nil, instead nil")
	}
	if len(trie.value) > 0 {
		t.Fatalf("Failed to Insert. Expecting 0, instead %d", len(trie.value))
	}

	c := b.children.Get('c').(*Trie)
	if c.value != "abc" {
		t.Fatalf("Failed to Insert. Expecting 'abc', instead %s", c.value)
	}
}

func TestFind(t *testing.T) {
	hash := hashtable.NewHash(26)
	trie := NewTrie(hash, 26)

	trie.Insert("abc")
	trie.Insert("abcd")
	trie.Insert("qwerty")
	trie.Insert("a")

	if v, ok := trie.Find("abc"); !ok {
		t.Fatalf("Failed to Find. No string found")
	} else if v != "abc" {
		t.Fatalf("Failed to Find. Expecting 'abc', instead %s", v)
	}

	if v, ok := trie.Find("qwerty"); !ok {
		t.Fatalf("Failed to Find. No string found")
	} else if v != "qwerty" {
		t.Fatalf("Failed to Find. Expecting 'qwerty', instead %s", v)
	}

	if v, ok := trie.Find("a"); !ok {
		t.Fatalf("Failed to Find. No string found")
	} else if v != "a" {
		t.Fatalf("Failed to Find. Expecting 'a', instead %s", v)
	}

	if v, ok := trie.Find("ab"); ok {
		t.Fatalf("Failed to Find. Expecting not found, instead %s", v)
	}
}
