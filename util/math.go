package util

import (
	"errors"
	"fmt"
	"maps"
	"slices"
)

// Factor returns the prime factors of n. n must be greater than 1.
func Factor(n int) []int {
	if n < 2 {
		Must(errors.New("can only factor integers > 1"))
	}
	var factors []int
	for i := 2; i <= n; {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
		} else {
			i++
		}
	}
	return factors
}

// IsPrime returns whether n is prime.
func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	return len(Factor(n)) == 1
}

// Digits returns a slice of the digits of n. Digits(1234) returns [1 2 3 4].
func Digits(n int) []int {
	var d []int
	for n > 0 {
		d = append(d, n%10)
		n /= 10
	}
	slices.Reverse(d)
	return d
}

// RuneCount returns a map containing each rune in s and how many times it occurs.
func RuneCount(s string) map[rune]int {
	ret := map[rune]int{}
	for _, r := range s {
		ret[r]++
	}
	return ret
}

// PrintIncreasing prints the keys and values in the map in increasing order of the keys.
func PrintIncreasing(m map[rune]int) {
	// TODO: use generics. May not be possible to print a rune as a char with %v.
	keys := slices.Sorted(maps.Keys(m))
	for _, k := range keys {
		fmt.Printf("%c: %d\n", k, m[k])
	}
}
