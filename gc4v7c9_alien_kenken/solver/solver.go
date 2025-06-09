package solver

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/bitlux/caches/util"
)

const SIZE = 9

type Board [SIZE][SIZE]int

func (b Board) String() string {
	s := ""
	for i := range SIZE {
		for j := range SIZE {
			s += fmt.Sprintf("%d ", b[i][j])
		}
		s += "\n"
	}
	return s
}

func (b Board) clone() Board {
	ret := Board{}
	for i := range SIZE {
		for j := range SIZE {
			ret[i][j] = b[i][j]
		}
	}
	return ret
}

// isValid checks that all entries in a row and column are unique. 0's are ignored.
func (b Board) isValid() bool {
	for row := range SIZE {
		for colA := range SIZE {
			a := b[row][colA]
			if a == 0 {
				continue
			}
			for colB := colA + 1; colB < SIZE; colB++ {
				b := b[row][colB]
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
			a := b[rowA][col]
			if a == 0 {
				continue
			}
			for rowB := rowA + 1; rowB < SIZE; rowB++ {
				b := b[rowB][col]
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

// len(cells) == len(i) for _, i := range candidates
type Cage struct {
	Candidates [][]int
	Cells      []string
}

// Puzzle contains metadata about a puzzle, such as statistics.
type Puzzle struct {
	cages []Cage

	ch     chan Board
	cancel context.CancelFunc

	invalidCount int
	farthest     map[int]bool
	start        time.Time
}

func New(cages []Cage) *Puzzle {
	return &Puzzle{
		cages:    cages,
		farthest: map[int]bool{},
	}
}

func (p *Puzzle) dump() {
	fmt.Println(time.Since(p.start).Truncate(time.Second), "invalid:", p.invalidCount)
}

func sheetsToIndices(s string) (int, int) {
	i, err := strconv.Atoi(s[1:])
	util.Must(err)
	return i - 1, util.A1Z26(int(s[0]) - 1)
}

func (p *Puzzle) solve(ctx context.Context, index int, b Board) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	if !p.farthest[index] {
		p.farthest[index] = true
		fmt.Println("depth", index)
	}
	if index == len(p.cages) {
		p.ch <- b
		p.cancel()
		return
	}
	cage := p.cages[index]
	for _, candidate := range cage.Candidates {
		perms := util.Permutations(candidate)
		for _, perm := range perms {
			// Assign permutation of candidate to cells in cage.
			clone := b.clone()
			for pc, cell := range cage.Cells {
				i, j := sheetsToIndices(cell)
				clone[i][j] = perm[pc]
			}
			if clone.isValid() {
				p.solve(ctx, index+1, clone)
			} else {
				p.invalidCount++
			}
		}
	}
}

func (p *Puzzle) Solve() Board {
	go func() {
		for range time.Tick(5 * time.Second) {
			p.dump()
		}
	}()
	defer p.dump()

	ctx, cancel := context.WithCancel(context.Background())

	p.cancel = cancel
	p.start = time.Now()
	p.ch = make(chan Board)

	go p.solve(ctx, 0, Board{})

	return <-p.ch
}

// Candidates returns all slices of the specified size containing elements such that
// f(elem) == target.
func Candidates(f func([]int) int, target int, size int) [][]int {
	var ret [][]int
	for _, cand := range util.Combinations(size) {
		if f(cand) == target {
			ret = append(ret, cand)
		}
	}
	return ret
}
