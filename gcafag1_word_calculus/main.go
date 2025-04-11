package main

import (
	"fmt"
	"slices"

	"github.com/bitlux/caches/util"
)

var (
	ones = map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'l': true,
		'n': true,
		's': true,
		't': true,
		'r': true,
	}

	threes = map[byte]bool{
		'b': true,
		'c': true,
		'm': true,
		'p': true,
	}

	fours = map[byte]bool{
		'f': true,
		'h': true,
		'v': true,
		'w': true,
		'y': true,
	}
)

func runeToOrd(r rune) int {
	return int(r) - 'a' + 1
}

func isAllOnes(s string) bool {
	for _, r := range s {
		if t := ones[byte(r)]; !t {
			return false
		}
	}
	return true
}

func is341111(s string) bool {
	f := func(m map[byte]bool) int {
		count := 0
		for _, r := range s {
			if m[byte(r)] {
				count++
			}
		}
		return count
	}

	return f(ones) == 4 && f(threes) == 1 && f(fours) == 1
}

func sum(s string) int {
	total := 0
	for _, r := range s {
		total += runeToOrd(r)
	}
	return total
}

func Range(s string) int {
	min := 100
	max := 0
	for _, r := range s {
		i := runeToOrd(r)
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	r := max - min
	return r
}

func median(s string) float64 {
	var ints []int
	for _, r := range s {
		ints = append(ints, runeToOrd(r))
	}
	slices.Sort(ints)
	middle := len(ints) / 2
	if len(ints)%2 == 0 {
		return float64(ints[middle-1]+ints[middle]) / 2
	}
	return float64(ints[middle])
}

func main() {
	lines := util.ReadLines("/home/abender/Adam/projects/wordlist-20210729.txt")
	for _, word := range lines {
		if len(word) == 4 && isAllOnes(word) && sum(word) == 58 && Range(word) == 15 {
			fmt.Println("[2]", word)
		}
		if len(word) == 5 && isAllOnes(word) && sum(word) == 64 && Range(word) == 19 && median(word) == 19 {
			fmt.Println("[1]", word)
		}
		if len(word) == 6 && sum(word) == 58 && Range(word) == 18 && median(word) == 8.5 && is341111(word) {
			fmt.Println("[3]", word)
		}
	}
}
