package common

import "math"

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
