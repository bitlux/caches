package datastructures

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert([]byte{1, 2, 3})
	trie.Insert([]byte{1, 2, 4})
	trie.Insert([]byte{5})
	trie.Insert([]byte{5, 6, 7})
	trie.Print()

	for _, tc := range []struct {
		value []byte
		want  bool
	}{
		{[]byte{1, 2, 3}, true},
		{[]byte{1, 2, 4}, true},
		{[]byte{1, 2, 2}, false},
		{[]byte{1}, false},
		{[]byte{2}, false},
		{[]byte{0}, false},
		{[]byte{1, 2}, false},
		{[]byte{1, 2, 3, 3}, false},
		{[]byte{5, 6, 7}, true},
	} {
		if got := trie.Contains(tc.value); got != tc.want {
			t.Errorf("Contains(%v) = %t, want %t", tc.value, got, tc.want)
		}
	}

	trie = NewTrie()
	trie.InsertString("adam")
	trie.InsertString("ada")
	trie.InsertString("")
	trie.InsertString("xyz")
	trie.Print()

	for _, tc := range []struct {
		value string
		want  bool
	}{
		{"adam", true},
		{"", true},
		{"adamm", false},
		{"aa", false},
		{"a", false},
		{"xyz", true},
	} {
		if got := trie.ContainsString(tc.value); got != tc.want {
			t.Errorf("ContainsString(%v) = %t, want %t", tc.value, got, tc.want)
		}
	}

	fmt.Println(trie.Elements())
}
