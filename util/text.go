package util

// Histogram produces a count of all of the letters that appear in s.
func Histogram(s string) map[rune]int {
	histo := map[rune]int{}
	for _, r := range s {
		histo[r]++
	}
	return histo
}

// RuneCount returns a map containing each rune in s and how many times it occurs.
func RuneCount(s string) map[rune]int {
	ret := map[rune]int{}
	for _, r := range s {
		ret[r]++
	}
	return ret
}
