package util

import "slices"

// D contains the digits 0-9. It is a common input to Explode.
var D = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// Explode takes a list of options, where each option is a list of ints. It returns the cross
// product of all options. Explode({1, 2}, {3, 4}) = {{1, 3}, {1, 4}, {2, 3}, {2, 4}}
func Explode(starting [][]int) [][]int {
	return explodeHelper(0, starting, [][]int{{}})
}

func explodeHelper(index int, values [][]int, soFar [][]int) [][]int {
	if index == len(values) {
		return soFar
	}

	temp := [][]int{}
	for _, value := range values[index] {
		for _, sf := range soFar {
			cl := slices.Clone(sf)
			temp = append(temp, append(cl, value))
		}
	}
	return explodeHelper(index+1, values, temp)
}
