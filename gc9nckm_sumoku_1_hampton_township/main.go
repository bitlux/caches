package main

import (
	"fmt"

	"github.com/bitlux/caches/gc4v7c9_alien_kenken/solver"
	"github.com/bitlux/caches/util"
)

func Sum(nums []int) int {
	// Optimization: if nums contains duplicates, return an impossible sum so that it is not
	// considered as a candidate.
	if !util.IsUnique(nums...) {
		return -1
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	p := solver.New([]solver.Cage{
		// row 1
		{
			Candidates: solver.Candidates(Sum, 9, 2),
			Cells:      []string{"A1", "B1"},
		},
		{
			Candidates: solver.Candidates(Sum, 15, 2),
			Cells:      []string{"C1", "D1"},
		},
		{
			Candidates: solver.Candidates(Sum, 9, 2),
			Cells:      []string{"E1", "F1"},
		},
		{ // pseudo
			Candidates: [][]int{{3}},
			Cells:      []string{"G1"},
		},
		{
			Candidates: [][]int{{4, 5}},
			Cells:      []string{"H1", "I1"},
		},
		// row 2
		{
			Candidates: solver.Candidates(Sum, 14, 2),
			Cells:      []string{"A2", "B2"},
		},
		{
			Candidates: solver.Candidates(Sum, 4, 2),
			Cells:      []string{"C2", "C3"},
		},
		{
			Candidates: solver.Candidates(Sum, 6, 2),
			Cells:      []string{"D2", "E2"},
		},
		{
			Candidates: solver.Candidates(Sum, 15, 2),
			Cells:      []string{"F2", "F3"},
		},
		{ // pseudo
			Candidates: solver.Candidates(Sum, 12, 2),
			Cells:      []string{"G2", "G3"},
		}, {
			Candidates: solver.Candidates(Sum, 15, 2),
			Cells:      []string{"H2", "H3"},
		},
		{
			Candidates: solver.Candidates(Sum, 3, 2),
			Cells:      []string{"I2", "I3"},
		},
		// row 3
		{
			Candidates: solver.Candidates(Sum, 10, 2),
			Cells:      []string{"A3", "A4"},
		},
		{
			Candidates: solver.Candidates(Sum, 12, 2),
			Cells:      []string{"B3", "B4"},
		},
		{
			Candidates: solver.Candidates(Sum, 9, 2),
			Cells:      []string{"D3", "E3"},
		},
		// row 4
		{
			Candidates: solver.Candidates(Sum, 12, 2),
			Cells:      []string{"C4", "C5"},
		},
		{
			Candidates: solver.Candidates(Sum, 10, 2),
			Cells:      []string{"D4", "E4"},
		},
		{
			Candidates: solver.Candidates(Sum, 4, 2),
			Cells:      []string{"F4", "F5"},
		},
		{
			Candidates: solver.Candidates(Sum, 3, 2),
			Cells:      []string{"G4", "G5"},
		},
		{
			Candidates: solver.Candidates(Sum, 7, 2),
			Cells:      []string{"H4", "H5"},
		},
		{
			Candidates: solver.Candidates(Sum, 15, 2),
			Cells:      []string{"I4", "I5"},
		},
		// row 5
		{
			Candidates: solver.Candidates(Sum, 12, 2),
			Cells:      []string{"A5", "B5"},
		},
		{
			Candidates: solver.Candidates(Sum, 17, 2),
			Cells:      []string{"D5", "E5"},
		},
		// row 6
		{
			Candidates: solver.Candidates(Sum, 10, 2),
			Cells:      []string{"A6", "A7"},
		},
		{
			Candidates: solver.Candidates(Sum, 9, 2),
			Cells:      []string{"B6", "C6"},
		},
		{
			Candidates: solver.Candidates(Sum, 5, 2),
			Cells:      []string{"D6", "E6"},
		},
		{
			Candidates: solver.Candidates(Sum, 12, 2),
			Cells:      []string{"F6", "F7"},
		},
		{
			Candidates: solver.Candidates(Sum, 10, 2),
			Cells:      []string{"G6", "G7"},
		},
		{
			Candidates: solver.Candidates(Sum, 18, 3),
			Cells:      []string{"H6", "H7", "H8"},
		},
		{
			Candidates: solver.Candidates(Sum, 11, 2),
			Cells:      []string{"I6", "I7"},
		},
		// row 7
		{
			Candidates: solver.Candidates(Sum, 16, 3),
			Cells:      []string{"B7", "B8", "B9"},
		},
		{
			Candidates: solver.Candidates(Sum, 7, 2),
			Cells:      []string{"C7", "C8"},
		},
		{
			Candidates: solver.Candidates(Sum, 14, 2),
			Cells:      []string{"D7", "D8"},
		},
		{
			Candidates: solver.Candidates(Sum, 9, 2),
			Cells:      []string{"E7", "E8"},
		},
		// row 8
		{
			Candidates: solver.Candidates(Sum, 5, 2),
			Cells:      []string{"A8", "A9"},
		},
		{
			Candidates: solver.Candidates(Sum, 12, 2),
			Cells:      []string{"F8", "G8"},
		},
		{
			Candidates: solver.Candidates(Sum, 11, 2),
			Cells:      []string{"I8", "I9"},
		},
		// row 9
		{
			Candidates: solver.Candidates(Sum, 8, 2),
			Cells:      []string{"C9", "D9"},
		},
		{
			Candidates: solver.Candidates(Sum, 13, 2),
			Cells:      []string{"E9", "F9"},
		},
		{
			Candidates: solver.Candidates(Sum, 10, 2),
			Cells:      []string{"G9", "H9"},
		},
	})

	fmt.Println(p.Solve())
}
