package util

import (
	"maps"
	"slices"
	"strconv"
)

// Permutations returns all permutations of the elements of s. If the elements
// are not unique, then the return value with contain duplicates.
func Permutations(s []int) [][]int {
	var ret [][]int
	var inner func(curr []int, index int)
	inner = func(curr []int, index int) {
		if index == len(s) {
			ret = append(ret, slices.Clone(curr))
			return
		}

		for i := index; i < len(s); i++ {
			s[index], s[i] = s[i], s[index]
			inner(s, index+1)
			s[index], s[i] = s[i], s[index]
		}
	}

	inner(s, 0)
	return ret
}

// Combinations returns all combinations of the digits 1-9 of the specified length.
// Returned combinations are unique up to ordering (if [1, 2] is included, [2, 1] will not be).
// Combinations(2) returns [[1, 1], [1, 2], ..., [8, 9], [9, 9]] (45 elements).
// TODO: generalize to support a supplied alphabet.
func Combinations(length int) [][]int {
	return combos(length, [][]int{{}})
}

// TODO: This should not sort.
func unique(nums [][]int) [][]int {
	m := map[string]bool{}
	for _, slice := range nums {
		var s string
		slices.Sort(slice)
		for _, n := range slice {
			s += strconv.Itoa(n)
		}
		m[s] = true
	}

	var ret [][]int
	for _, k := range slices.Sorted(maps.Keys(m)) {
		n, _ := strconv.Atoi(k)
		ret = append(ret, Digits(n))
	}
	return ret
}

func combos(remaining int, soFar [][]int) [][]int {
	if remaining == 0 {
		return unique(soFar)
	}
	var ret [][]int
	for n := 1; n < 10; n++ {
		for _, curr := range soFar {
			ret = append(ret, append(slices.Clone(curr), n))
		}
	}
	return combos(remaining-1, ret)
}
