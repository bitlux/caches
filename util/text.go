package util

// Histogram produces a count of all of the letters that appear in s.
func Histogram(s string) map[rune]int {
	histo := map[rune]int{}
	for _, r := range s {
		histo[r]++
	}
	return histo
}

// CBF encodes a string into a slice of integers. CBF encoding is similar to A1Encode, but done
// mod 10.
func CBF(s string) []int {
	var ret []int
	for _, c := range s {
		ret = append(ret, A1Encode(c)%10)
	}
	return ret
}
