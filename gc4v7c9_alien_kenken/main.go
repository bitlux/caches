package main

import (
	"fmt"
	"math"
)

func round(n float64) int {
	return int(math.Round(n - 0.5))
}

func clubs(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n * n
	}

	return sum
}

func hearts(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n * n * n
	}

	return round(math.Sqrt(float64(sum)))
}

func spades(nums []int) int {
	sum := 0
	prod := 1
	for _, n := range nums {
		sum += n
		prod *= n
	}
	return sum + round(math.Sqrt(float64(prod)))
}

func diamonds(nums []int) int {
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

func gen2() [][]int {
	var ret [][]int
	for a := 1; a < 9; a++ {
		for b := a + 1; b < 10; b++ {
			ret = append(ret, []int{a, b})
		}
	}
	return ret
}

func gen3() [][]int {
	var ret [][]int
	for a := 1; a < 8; a++ {
		for b := a + 1; b < 9; b++ {
			for c := b + 1; c < 10; c++ {
				ret = append(ret, []int{a, b, c})
			}
		}
	}
	return ret
}

func gen4() [][]int {
	var ret [][]int
	for a := 1; a < 7; a++ {
		for b := a + 1; b < 8; b++ {
			for c := b + 1; c < 9; c++ {
				for d := c + 1; d < 10; d++ {
					ret = append(ret, []int{a, b, c, d})
				}
			}
		}
	}
	return ret
}

func gen5() [][]int {
	var ret [][]int
	for a := 1; a < 6; a++ {
		for b := a + 1; b < 7; b++ {
			for c := b + 1; c < 8; c++ {
				for d := c + 1; d < 9; d++ {
					for e := d + 1; e < 10; e++ {
						ret = append(ret, []int{a, b, c, d, e})
					}
				}
			}
		}
	}
	return ret
}

func main() {
	for _, tc := range []struct {
		op     func([]int) int
		gen    func() [][]int
		target int
	}{
		// {clubs, gen5, 237},
		// {diamonds, gen3, 29},
		// {diamonds, gen5, 72},
		// {hearts, gen4, 36},
		// {hearts, gen4, 34},
		// {hearts, gen4, 39},
		// {hearts, gen4, 20},
		{spades, gen5, 104},
		// {hearts, gen3, 18},
		{clubs, gen5, 130},
		{spades, gen4, 21},
		{diamonds, gen4, 79},
		{clubs, gen5, 91},
		{diamonds, gen5, 84},
		// {hearts, gen5, 35},
		{spades, gen3, 24},
	} {
		for _, cand := range tc.gen() {
			result := tc.op(cand)
			//			fmt.Println(cand, op)
			if result == tc.target {
				fmt.Println(cand, result)
			}
		}
		fmt.Println()
	}
}
