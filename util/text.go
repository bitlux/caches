package util

func Histogram(s string) map[rune]int {
	histo := map[rune]int{}
	for _, r := range s {
		histo[r]++
	}
	return histo
}
