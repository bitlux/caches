package util

import (
	"iter"
)

// D contains the digits 0-9. It is a common input to Explode.
var D = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// Explode takes a list of options, where each option is a list of ints. It returns the cross
// product of all options. Explode({1, 2}, {3, 4}) = {{1, 3}, {1, 4}, {2, 3}, {2, 4}}
func Explode(vals ...[]int) iter.Seq[[]int] {
	type fn = func([]int) bool

	var helper func(yield fn, index int, soFar []int)
	helper = func(yield fn, index int, soFar []int) {
		if index == len(vals) {
			yield(soFar)
			return
		}

		for _, val := range vals[index] {
			helper(yield, index+1, append(soFar, val))
		}
	}

	return func(yield fn) {
		helper(yield, 0, []int{})
	}
}
