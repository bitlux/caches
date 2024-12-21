package util

import (
	"errors"
	"slices"
)

func Factor(n int) []int {
	if n < 2 {
		Must(errors.New("can only factor integers > 1"))
	}
	factors := []int{}
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

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	return len(Factor(n)) == 1
}

func Digits(n int) []int {
	d := []int{}
	for n > 0 {
		d = append(d, n%10)
		n /= 10
	}
	slices.Reverse(d)
	return d
}
