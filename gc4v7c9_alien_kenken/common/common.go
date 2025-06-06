package common

import (
	"maps"
	"math"
	"slices"
	"strconv"
)

func round(n float64) int {
	return int(math.Round(n - 0.5))
}

func Clubs(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n * n
	}

	return sum
}

func Hearts(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n * n * n
	}

	return round(math.Sqrt(float64(sum)))
}

func Spades(nums []int) int {
	sum := 0
	prod := 1
	for _, n := range nums {
		sum += n
		prod *= n
	}
	return sum + round(math.Sqrt(float64(prod)))
}

func Diamonds(nums []int) int {
	sum := 0
	prod := 1
	denom := 0.0
	for _, n := range nums {
		sum += n
		prod *= n
		denom += 1.0 / float64(n)
	}

	n := float64(len(nums))
	nthRoot := math.Pow(float64(prod), 1.0/n)
	return round(n * (nthRoot + float64(sum)/n + n/denom))
}

const SIZE = 9

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
		ret = append(ret, digits(n))
	}
	return ret
}

func digits(n int) []int {
	var d []int
	for n > 0 {
		d = append(d, n%10)
		n /= 10
	}
	return d
}

func recurse(length int, soFar [][]int) [][]int {
	if length == 0 {
		return unique(soFar)
	}
	var ret [][]int
	for n := 1; n < 10; n++ {
		for _, curr := range soFar {
			ret = append(ret, append(curr, n))
		}
	}
	return recurse(length-1, ret)
}

func Candidates(f func([]int) int, target int, size int) [][]int {
	var ret [][]int
	for _, cand := range recurse(size, [][]int{{}}) {
		if f(cand) == target {
			ret = append(ret, cand)
		}
	}
	return ret
}
