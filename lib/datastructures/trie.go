package datastructures

import (
	"fmt"
	"strings"
)

// Trie is a trie that stores byte slices. It is useful for storing (arbitrary-length) integers (one
// digit at a time) or strings. It supports insert and lookup in time O(length) of the slice.
type Trie struct {
	terminal bool
	children [256]*Trie
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(s []byte) {
	curr := t
	for _, e := range s {
		if ch := curr.children[e]; ch == nil {
			curr.children[e] = &Trie{}
		}
		curr = curr.children[e]
	}
	curr.terminal = true
}

func (t *Trie) InsertString(s string) {
	t.Insert([]byte(s))
}

func (t *Trie) Contains(s []byte) bool {
	curr := t
	for _, e := range s {
		if ch := curr.children[e]; ch == nil {
			return false
		}
		curr = curr.children[e]
	}
	return curr.terminal
}

func (t *Trie) ContainsString(s string) bool {
	return t.Contains([]byte(s))
}

func (t *Trie) Elements() [][]byte {
	var ret [][]byte

	var recurse func(t *Trie, soFar []byte)
	recurse = func(t *Trie, soFar []byte) {
		if t.terminal {
			ret = append(ret, soFar)
		}
		for i, n := range t.children {
			if n != nil {
				recurse(n, append(soFar, byte(i)))
			}
		}
	}
	recurse(t, nil)
	return ret
}

func (t *Trie) Print() {
	t.printPrefix(0)
}

// TODO:
//  1. Support strings
//  2. Fix formatting. Ideally it would print top-down, but even one element per line would be better
//     than this.
func (t *Trie) printPrefix(depth int) {
	if t.terminal {
		fmt.Print(" *")
	}
	fmt.Println()
	for i, n := range t.children {
		if n != nil {
			fmt.Printf("%s%d", strings.Repeat("|", depth), i)
			n.printPrefix(depth + 1)
		}
	}
}
