package main

import (
	"fmt"
	"math"

	"github.com/bitlux/caches/gc4v7c9_alien_kenken/solver"
	"github.com/bitlux/caches/util"
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

func main() {
	p := solver.New([]solver.Cage{
		// row 1
		{
			Candidates: [][]int{{1, 5, 9}},
			Cells:      []string{"A1", "B1", "C1"},
		},
		{
			Candidates: [][]int{{2, 3, 6}},
			Cells:      []string{"D1", "E1", "F1"},
		},
		{
			Candidates: [][]int{{4, 7, 8}},
			Cells:      []string{"G1", "H1", "I1"},
		},

		// row 2
		{
			Candidates: [][]int{{7, 9}},
			Cells:      []string{"A2", "A3"},
		},
		{
			Candidates: solver.Candidates(Hearts, 36, 5),
			Cells:      []string{"B2", "C2", "D2", "B3", "B4"},
		},
		{
			Candidates: [][]int{{1, 5}},
			Cells:      []string{"E2", "E3"},
		},
		{
			Candidates: solver.Candidates(Hearts, 34, 5),
			Cells:      []string{"F2", "G2", "H2", "H3", "H4"},
		},
		{
			Candidates: [][]int{{3, 4}},
			Cells:      []string{"I2", "I3"},
		},

		// row 3
		{
			Candidates: solver.Candidates(Hearts, 39, 4),
			Cells:      []string{"C3", "D3", "C4", "D4"},
		},
		{
			Candidates: solver.Candidates(Hearts, 20, 4),
			Cells:      []string{"F3", "G3", "F4", "G4"},
		},

		// row 4
		{
			Candidates: [][]int{{1, 3}, {2, 3}},
			Cells:      []string{"A4", "A6"}, // 9h
		},
		{
			Candidates: [][]int{{2, 8}},
			Cells:      []string{"E4", "E6"}, // 104s
		},
		{
			Candidates: [][]int{{2, 7}, {5, 6}},
			Cells:      []string{"I4", "I6"}, // 18h
		},

		// row 5
		{
			Candidates: [][]int{{4}},
			Cells:      []string{"A5"},
		},
		{
			Candidates: [][]int{{3, 5}},
			Cells:      []string{"B5", "C5"}, // 34c
		},
		{
			Candidates: [][]int{{6, 7, 8}},
			Cells:      []string{"D5", "E5", "F5"}, // 104s
		},
		{
			Candidates: [][]int{{2, 9}},
			Cells:      []string{"G5", "H5"}, // 15s
		},
		{
			Candidates: [][]int{{1}},
			Cells:      []string{"I5"},
		},

		// row 6
		{
			Candidates: solver.Candidates(Clubs, 130, 5),
			Cells:      []string{"B6", "B7", "B8", "C8", "D8"},
		},
		{
			Candidates: solver.Candidates(Spades, 21, 4),
			Cells:      []string{"C6", "D6", "C7", "D7"},
		},
		{
			Candidates: solver.Candidates(Diamonds, 79, 4),
			Cells:      []string{"F6", "G6", "F7", "G7"},
		},
		{
			Candidates: solver.Candidates(Clubs, 91, 5),
			Cells:      []string{"H6", "H7", "F8", "G8", "H8"},
		},

		// row 7
		{
			Candidates: solver.Candidates(Diamonds, 84, 5),
			Cells:      []string{"A7", "A8", "A9", "B9", "C9"},
		},
		{
			Candidates: [][]int{{4, 9}},
			Cells:      []string{"E7", "E8"}, // 28h
		},
		{
			Candidates: solver.Candidates(Hearts, 35, 5),
			Cells:      []string{"I7", "I8", "G9", "H9", "I9"},
		},

		// row 9
		{
			Candidates: [][]int{{1, 9}, {4, 8}},
			Cells:      []string{"D9", "F9"}, // 24s
		},
		{
			Candidates: [][]int{{3}, {7}},
			Cells:      []string{"E9"}, // 24s
		},
	})

	fmt.Println(len(util.Combinations(2)))

	b := p.Solve()
	fmt.Printf("N 37 23.%03d\n", Clubs([]int{b[5][1], b[8][2], b[2][8], b[2][4]}))
	fmt.Printf("W 122 08.%03d\n", Clubs([]int{b[8][5], b[3][3], b[5][6], b[1][1], b[3][6]}))
}
