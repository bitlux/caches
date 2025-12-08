package util

import (
	"iter"
	"slices"
)

// RuneCount returns a map containing each rune in s and how many times it occurs.
func RuneCount(s string) map[rune]int {
	ret := map[rune]int{}
	for _, r := range s {
		ret[r]++
	}
	return ret
}

// SortLetters sorts the letters of ASCII strings. SortLetter("asdf") == "adfs". This is useful for
// finding anagrams.
func SortLetters(s string) string {
	b := []byte(s)
	slices.Sort(b)
	return string(b)
}

// Alphabet returns an iter.Seq over all lowercase letters.
func Alphabet() iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for r := 'a'; r <= 'z'; r++ {
			if !yield(r) {
				return
			}
		}
	}
}
