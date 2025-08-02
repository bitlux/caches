package util

import (
	"errors"
	"slices"

	"golang.org/x/exp/constraints"
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

// FromDigitsBase takes a slice of digits in the provided base and returns them as a single number.
// It is the inverse of Digits.
func FromDigitsBase(digits []int, base int) int {
	ret := 0
	for _, d := range digits {
		ret = ret*base + d
	}
	return ret
}

// FromDigits takes a slice of digits and returns them as a single number. It is the inverse of
// Digits.
func FromDigits(digits []int) int {
	return FromDigitsBase(digits, 10)
}

// RuneCount returns a map containing each rune in s and how many times it occurs.
func RuneCount(s string) map[rune]int {
	ret := map[rune]int{}
	for _, r := range s {
		ret[r]++
	}
	return ret
}

// A1Encode encodes a rune in the range [A-Za-z] using the A=1, ..., Z=26 substitution cipher.
func A1Encode[T constraints.Integer](n T) int {
	if n < 91 {
		return int(n - 'A' + 1)
	}
	return int(n - 'a' + 1)
}

// A1Decode decodes a number in the range [1-26] using the A=1, ..., Z=26 substitution cipher.
func A1Decode(n int) rune {
	return rune(n + 'A' - 1)
}
