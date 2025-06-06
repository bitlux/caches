package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"time"

	"github.com/bitlux/caches/gc4v7c9_alien_kenken/common"
	"github.com/bitlux/caches/util"
)

const SIZE = common.SIZE

var (
	validCount, invalidCount int
	farthest                 map[int]string
)

type Cage struct {
	// len(cells) == len(i) for _, i := range candidates
	candidates [][]int
	cells      []string
}

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
		candidates: [][]int{{7, 9}},
		cells:      []string{"A2", "A3"},
	},
	{
		candidates: common.Candidates(common.Hearts, 36, 5),
		cells:      []string{"B2", "C2", "D2", "B3", "B4"},
	},
	{
		candidates: [][]int{{1, 5}},
		cells:      []string{"E2", "E3"},
	},
	{
		candidates: common.Candidates(common.Hearts, 34, 5),
		cells:      []string{"F2", "G2", "H2", "H3", "H4"},
	},
	{
		candidates: [][]int{{3, 4}},
		cells:      []string{"I2", "I3"},
	},

	{
		candidates: common.Candidates(common.Hearts, 39, 4),
		cells:      []string{"C3", "D3", "C4", "D4"},
	},
	{
		candidates: common.Candidates(common.Hearts, 20, 4),
		cells:      []string{"F3", "G3", "F4", "G4"},
	},

	{
		candidates: [][]int{{1, 3}, {2, 3}},
		cells:      []string{"A4", "A6"}, // 9h
	},
	{
		candidates: [][]int{{2, 8}},
		cells:      []string{"E4", "E6"}, // 104s
	},
	{
		candidates: [][]int{{2, 7}, {5, 6}},
		cells:      []string{"I4", "I6"}, // 18h
	},

	{
		candidates: [][]int{{3, 5}},
		cells:      []string{"B5", "C5"}, // 34c
	},
	{
		candidates: [][]int{{6, 7, 8}},
		cells:      []string{"D5", "E5", "F5"}, // 104s
	},
	{
		candidates: [][]int{{2, 9}},
		cells:      []string{"G5", "H5"}, // 15s
	},

	{
		candidates: common.Candidates(common.Spades, 130, 5),
		cells:      []string{"B6", "B7", "B8", "C8", "D8"},
	},
	{
		candidates: common.Candidates(common.Spades, 21, 4),
		cells:      []string{"C6", "D6", "C7", "D7"},
	},
	{
		candidates: common.Candidates(common.Diamonds, 79, 4),
		cells:      []string{"F6", "G6", "F7", "G7"},
	},
	{
		candidates: common.Candidates(common.Clubs, 91, 5),
		cells:      []string{"H6", "H7", "F8", "G8", "H8"},
	},

	{
		candidates: common.Candidates(common.Diamonds, 84, 5),
		cells:      []string{"A7", "A8", "A9", "B9", "C9"},
	},
	{
		candidates: [][]int{{4, 9}},
		cells:      []string{"E7", "E8"}, // 28h
	},
	{
		candidates: common.Candidates(common.Hearts, 35, 5),
		cells:      []string{"I7", "I8", "G9", "H9", "I9"},
	},

	// {
	// 	candidates: [][]int{{1, 9}, {4, 8}},
	// 	cells:      []string{"D9", "F9"}, // 24s
	// },
	// {
	// 	candidates: [][]int{{3}, {7}},
	// 	cells:      []string{"E9"}, // 24s
	// },
}

func sheetsToIndices(s string) (int, int) {
	i, err := strconv.Atoi(s[1:])
	util.Must(err)
	return i - 1, util.A1Z26(int(s[0]) - 1)
}

type Board struct {
	// Current board state
	values [SIZE][SIZE]int
}

func (b *Board) String() string {
	s := ""
	for i := range SIZE {
		for j := range SIZE {
			s += fmt.Sprintf("%d ", b.values[i][j])
		}
		s += "\n"
	}
	return s
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
					invalidCount++
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
					invalidCount++
					return false
				}
			}
		}
	}

	return true
}

func (b *Board) recurse(index int) {
	if farthest[index] == "" {
		farthest[index] = b.String()
	}
	if index == len(cages) {
		if validCount%1_000_000 == 0 {
			fmt.Println(b)
		}
		validCount++
		return
	}
	cage := cages[index]
	for _, candidate := range cage.candidates {
		perms := permutations(candidate)
		for _, p := range perms {
			// Assign permutation of candidate to cells in cage.
			clone := b.clone()
			for pc, cell := range cage.cells {
				i, j := sheetsToIndices(cell)
				clone.values[i][j] = p[pc]
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

func dump() {
	fmt.Println("valid:", validCount, "invalid:", invalidCount)
}

func main() {
	farthest = map[int]string{}

	b := &Board{}
	b.values[4][0] = 4
	b.values[4][8] = 1

	go func() {
		for range time.Tick(5 * time.Second) {
			dump()
		}
	}()

	b.recurse(0)
	dump()

	for _, d := range slices.Sorted(maps.Keys(farthest)) {
		fmt.Printf("%d\n%s\n\n", d, farthest[d])
	}
}
