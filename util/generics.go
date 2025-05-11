package util

import (
	"fmt"
	"maps"
	"slices"
)

// PrintAscending prints the keys and values in the map in increasing order of the keys.
func PrintAscending(m map[rune]int) {
	// TODO: use generics. May not be possible to print a rune as a char with %v.
	keys := slices.Sorted(maps.Keys(m))
	for _, k := range keys {
		fmt.Printf("%c: %d\n", k, m[k])
	}
}

// IsUnique returns whether the elements of vals are all unique.
func IsUnique[T comparable](vals ...T) bool {
	m := map[T]bool{}
	for _, v := range vals {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}
