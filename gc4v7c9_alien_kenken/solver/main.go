package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/bitlux/caches/gc4v7c9_alien_kenken/common"
	"github.com/bitlux/caches/util"
)

const SIZE = common.SIZE

var cages = []Cage{
	{
		candidates: [][]int{{1, 5, 9}},
		cells:      []string{"A1", "B1", "C1"},
	},
	{
		candidates: [][]int{{2, 3, 6}},
		cells:      []string{"D1", "E1", "F1"},
	},
	{
		candidates: [][]int{{4, 7, 8}},
		cells:      []string{"G1", "H1", "I1"},
	},

	{
		candidates: [][]int{{2, 7}},
		cells:      []string{"A2", "A3"},
	},
	// {
	// 	candidates: [][]int{{1, 5}},
	// 	cells:      []string{"E2", "E3"},
	// },
	{
		candidates: [][]int{{3, 4}},
		cells:      []string{"I2", "I3"},
	},

	{
		candidates: [][]int{{1, 3}, {2, 3}},
		cells:      []string{"A4", "A6"},
	},
	{
		candidates: [][]int{{2, 7}, {5, 6}},
		cells:      []string{"I4", "I6"},
	},
}

func sheetsToIndices(s string) (int, int) {
	i, err := strconv.Atoi(s[1:])
	util.Must(err)
	return i - 1, util.A1Z26(int(s[0]) - 1)
}

type Cage struct {
	// len(cells) == len(i) for _, i := range list
	candidates [][]int
	cells      []string
}

type Board struct {
	// Current board state
	values [SIZE][SIZE]int
}

func (b *Board) dump() {
	for i := range SIZE {
		for j := range SIZE {
			fmt.Printf("%d ", b.values[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) clone() *Board {
	ret := &Board{}
	for i := range SIZE {
		for j := range SIZE {
			ret.values[i][j] = b.values[i][j]
		}
	}
	return ret
}

func (b *Board) isValid() bool {
	for row := range SIZE {
		for colA := range SIZE {
			a := b.values[row][colA]
			if a == 0 {
				continue
			}
			for colB := colA + 1; colB < SIZE; colB++ {
				b := b.values[row][colB]
				if b == 0 {
					continue
				}
				if a == b {
					return false
				}
			}
		}
	}

	for col := range SIZE {
		for rowA := range SIZE {
			a := b.values[rowA][col]
			if a == 0 {
				continue
			}
			for rowB := rowA + 1; rowB < SIZE; rowB++ {
				b := b.values[rowB][col]
				if b == 0 {
					continue
				}
				if a == b {
					return false
				}
			}
		}
	}

	return true
}

func (b *Board) recurse(index int) {
	if index == len(cages) {
		b.dump()
		return
	}
	cage := cages[index]
	for _, candidate := range cage.candidates {
		perms := permutations(candidate)
		for _, p := range perms {
			// Assign permutation of candidate to cells in cage.
			clone := b.clone()
			for index, cell := range cage.cells {
				i, j := sheetsToIndices(cell)
				clone.values[i][j] = p[index]
			}
			if clone.isValid() {
				clone.recurse(index + 1)
			}
		}
	}
}

func permutations(s []int) [][]int {
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

func main() {
	b := &Board{}
	b.values[0][4] = 4
	b.values[8][4] = 1

	b.recurse(0)
}
