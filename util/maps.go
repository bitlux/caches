package util

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"unicode"
)

func printAscending[T cmp.Ordered, U any](format string, m map[T]U) {
	keys := slices.Sorted(maps.Keys(m))
	for _, k := range keys {
		fmt.Printf(format, k, m[k])
	}
}

// PrintAscending prints the keys and values in the map in increasing order of the keys. If it
// determines that each key in the map is a rune, it will print the rune using %c. Other values
// are printed with %v.
func PrintAscending[T cmp.Ordered, U any](m map[T]U) {
	runes := map[rune]U{}
	isUnicode := true

loop:
	for k, v := range m {
		switch r := any(k).(type) {
		case int32:
			if !(unicode.IsGraphic(r) || unicode.IsSpace(r)) {
				isUnicode = false
				break loop
			}
			runes[r] = v
		default:
			isUnicode = false
			break loop
		}
	}

	if isUnicode {
		printAscending("'%c': %v\n", runes)
	} else {
		printAscending("%v: %v\n", m)
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

func AsSet[T comparable](vals []T) map[T]bool {
	m := map[T]bool{}
	for _, v := range vals {
		m[v] = true
	}
	return m
}
