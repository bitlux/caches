package main

import (
	"fmt"
	"slices"
	"strconv"
)

const (
	smallGrid = `4 E 1 1 4 D D 0 0 1 7 8 F C A 8
 D 0 F 4 5 7 E 4 5 9 4 8 8 1 E
3 E 0 F E D 3 D 0 1 0 9 4 F 0 0
 7 F 9 E 1 2 D D F 5 E E D 4 4
8 F 9 6 D 3 F 0 F D D E 1 D E C
 F 1 9 D 4 1 0 1 2 0 6 9 9 1 3
4 4 0 0 8 1 B 4 3 D 2 7 F 8 9 F
 3 0 4 8 D F A 7 0 B 0 F 5 9 4
D 3 8 F 0 9 8 2 2 D F E 1 4 2 F
 E 4 6 9 D 4 7 E E A 2 F F A D
D D C 4 F E 6 F E 8 1 F 1 2 7 1
 1 D 6 B 5 4 1 3 1 1 A F D 6 D
D 2 2 7 1 7 0 E F D 9 1 8 3 F E
 E 7 2 D 8 B 0 0 4 D 6 2 1 F D
9 6 2 5 3 F 1 E 3 0 D 1 9 5 2 F
 6 5 2 1 0 D 7 8 A 2 7 4 3 3 7
1 2 1 0 3 D E 0 E 1 3 0 5 E D 4`

	bigGrid = `E 9 A E B D 3 F 4 4 1 2 E E 4 5 7 4 E 7 9 1 0 0 D D 9 F 0 F
 F A 4 E 9 7 7 9 A B D 8 E E 0 F 3 7 D 1 7 D E 1 3 1 2 D 9
D A 2 5 4 D D F F F 8 B 2 8 1 9 D F B 8 3 8 2 3 0 9 A E D F
 7 3 0 2 1 F D 0 1 7 2 D A 5 2 0 1 A 3 1 B E 5 2 1 E 2 0 F
2 2 D 8 8 2 3 6 1 9 5 1 D 2 8 9 7 5 F 4 4 B 5 4 9 7 3 2 A 8
 4 4 9 E 4 E 4 3 D A 2 5 1 2 2 5 F 8 F 2 3 2 8 0 E 2 A F D
7 5 3 2 F E 0 1 3 6 7 F 7 4 E 7 7 7 8 A 2 0 B 0 9 D 0 E E F
 F 4 4 3 3 F 8 A E 1 2 E 3 B 4 3 1 2 3 8 F 7 6 D 5 A B 4 E
2 3 2 1 7 F 4 1 3 3 3 D 5 F F F 9 5 F 2 7 E 9 8 8 4 8 2 4 E
 1 2 9 E D E 0 6 5 4 0 D 8 E 2 4 E 7 F E 0 D F 2 1 6 9 D 4
0 E D E 3 D F 5 7 4 3 A 0 9 E 7 E 0 6 9 3 E 3 6 2 3 0 2 1 4
 3 4 D 9 0 2 D F D 8 4 1 9 F 4 9 E 2 6 0 0 7 E 6 0 F E 4 3
F D 1 3 F 7 0 E B 6 E B 6 E 4 6 3 6 0 9 0 F F 1 4 6 4 4 0 4
 5 D 4 F 3 1 9 3 4 F 3 1 7 E 2 3 E 5 D 4 3 7 E 6 E E F 2 5
6 4 1 1 F 3 4 F F 1 6 4 6 0 4 E 0 3 9 E F A D F 4 3 7 2 E 5
 A 4 9 4 9 E D 6 F E 4 9 5 F 3 0 1 F 2 D F 2 6 0 6 3 4 F 7
0 E B B 2 F E 5 9 1 F D 5 E 4 A 4 F 2 9 2 B E B 3 4 1 0 F 9
 D 3 E D 0 D 7 1 2 A 4 6 3 4 1 6 4 4 D 3 2 E F 0 E E 2 1 B
F 5 7 5 1 9 5 8 0 3 5 4 0 E 4 E D 5 2 2 8 A 4 B E E 1 5 3 E
 9 0 E 5 F E E 3 A 3 E E 1 9 6 D F 5 E E 6 D B F D F 1 7 8
5 0 3 F 2 1 E E E 2 4 F D 7 2 3 8 8 2 F 8 0 D 1 2 0 0 2 1 B
 A 4 2 4 9 1 A 6 1 2 4 5 2 9 F F 6 1 F E 5 5 F E B 7 5 7 9
5 D D 2 1 4 2 D D 6 F 0 E 3 A 2 F 3 1 4 F B 9 2 9 4 4 A F F
 1 B 2 E 4 3 9 8 6 E F 1 4 E F 4 D D 0 E 9 3 D D E 0 4 F E
E F 6 D 6 D D D E 3 5 8 B 2 9 4 D 9 2 A 4 0 0 A 9 0 4 6 2 5
 0 9 4 7 1 A E 1 A 3 B 6 A E F D B 1 3 0 9 E 3 2 7 0 0 0 5
0 1 E 0 0 D E 3 5 8 7 8 E 5 0 4 E 2 2 E 2 D 0 0 E 4 3 A 9 E
 4 0 F F 0 F A E D 9 4 B B F 6 9 E 0 6 E F 1 0 1 0 D 8 8 9
5 F 7 7 D D 7 E F 4 5 B 9 6 3 5 7 0 0 F 7 8 D D 5 2 2 F A D
 1 F D 1 D 3 8 2 F E 4 D 1 9 1 D 4 E 1 A 0 3 7 3 D F F 1 1
3 B 4 0 6 4 1 F 2 7 D 9 2 A 9 8 9 E F 8 D 0 F D B 3 2 5 7 F
 D 2 3 2 3 F D E F 7 9 A 4 7 D 1 1 3 0 D F 3 D E F A 2 8 F
1 B D 0 3 5 8 2 3 0 F 5 F F 1 7 E 3 1 A E 3 B 6 5 B A 2 A 9
 A F 8 D 2 6 1 B 1 3 F 0 F E 2 D 0 0 2 6 F F D 1 F F 5 D D
4 1 6 F 5 D 5 4 D 0 8 3 A E 3 F F 3 0 F D 2 6 B 1 F B 3 F 0
 E E E 4 D D E A E F 5 5 D B 4 E 1 3 2 F 0 1 3 D D 0 E D F
E 2 D 8 9 2 3 F 2 4 4 0 8 F 0 8 7 4 1 F 0 1 F 1 B 3 1 2 E 3
 F 5 5 B 9 1 0 A 1 E F 9 5 E 2 0 A 9 F E 4 F 4 B D 8 F B 1
D 2 2 8 E F 2 A E D D 0 5 E E 7 9 7 7 0 0 D F 1 B 4 9 E F 4
 5 D 2 E D D A 0 E F 1 F 4 5 7 4 A F 3 1 6 0 A 7 9 A 2 4 0
A F B 2 E 8 5 3 E D 5 A 7 8 6 0 D 1 4 4 0 F 2 E 9 1 1 7 F D
 0 1 3 0 E F 7 7 3 E D A 4 E F 5 2 D 4 E 6 E D 4 F F 2 2 1
F E 0 1 E 0 1 8 6 4 3 B 8 7 5 E 4 A B D F 4 7 3 6 D 3 3 B B
 4 0 3 2 9 E 8 4 8 E D 0 2 F F F 1 2 2 B 0 E E 1 D A E D D
F 3 7 1 F F D 1 D F D D 6 2 4 B 0 4 4 5 2 A 9 4 D F F 8 4 D
 0 E 0 7 6 F E 1 6 0 A 3 9 4 B 4 6 1 F F 0 9 D E A D 4 F E
2 F 9 6 5 9 6 9 9 B 2 F D F D F 4 B 0 F 0 F F 0 5 4 F D 2 1
 3 4 F 8 F 4 2 5 3 F F E E 4 F 9 1 3 6 1 4 E F A 3 3 1 3 F
7 8 2 F 4 E E D 9 4 5 0 4 D 1 4 9 0 7 5 4 7 4 7 2 D 5 3 2 4
 1 2 B F 8 7 A 8 7 2 E 4 4 2 D 6 8 D 9 B E D E D 1 2 0 4 F
9 8 6 2 0 F 9 F F F F 9 1 3 0 4 B F E 8 5 2 D 0 F D 3 1 8 A
 6 6 3 0 0 D E 3 E 4 F 2 3 D 0 B 0 D 3 7 8 E 5 A 4 3 4 4 F
4 7 0 4 7 A B 5 2 D 2 0 F 5 1 7 4 F E F 9 4 0 5 9 E E 3 3 2
 0 6 6 2 0 1 5 2 2 1 8 F 1 D 6 A 1 4 D 3 F 5 2 F F F 1 9 7
4 F D 1 E 8 1 8 E D 6 3 0 B 2 0 A E 7 F E E A D B 6 1 B 2 2
 5 A 1 2 0 6 2 9 6 C E 3 F F D 0 9 D 4 1 3 2 D A 4 E 4 6 8
7 1 D A 2 2 2 4 E 8 2 D F 0 3 0 0 3 9 3 D E 3 4 3 E F 0 B 2
 D 2 7 E 0 B D 3 E 3 0 E 2 D 1 3 D 0 F 1 D 3 2 4 F 1 4 1 0
9 7 D F F D 4 7 1 D E 0 E 9 0 0 3 9 4 0 5 5 7 1 F 2 9 0 9 4
 6 F 9 F D 1 F 3 A 2 E 4 8 3 2 4 0 D 2 D 7 D 1 D 9 2 F B 9
9 4 9 3 9 D 4 D 7 E 7 3 3 1 2 0 D 4 2 2 E D A 8 2 E A E F E`
)

var _, _ = smallGrid, bigGrid

type Direction int

const (
	NE Direction = iota
	SE
	SW
	NW
)

// Upper left is (0, 0). Use int8 to save space.
type Point struct {
	row, col int8
}

func (p Point) Move(d Direction) (ret Point) {
	switch d {
	case NE:
		return Point{p.row - 1, p.col + 1}
	case SE:
		return Point{p.row + 1, p.col + 1}
	case SW:
		return Point{p.row + 1, p.col - 1}
	default:
		return Point{p.row - 1, p.col - 1}
	}
}

// Grid assumes the first and last rows are both "long" rows
type Grid struct {
	contents         []string
	numRows, numCols int8
}

func (g Grid) At(p Point) int {
	if !g.IsValid(p) {
		panic(fmt.Errorf("invalid At(%d, %d)", p.row, p.col))
	}

	s := g.contents[p.row][p.col : p.col+1]
	n, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		panic(fmt.Errorf("invalid rune at (%v)", p))
	}
	return int(n)
}

func (g Grid) IsValid(p Point) bool {
	if p.row < 0 || p.row >= g.numRows {
		return false
	}

	if p.col < 0 || p.col >= int8(len(g.contents[p.row])) {
		return false
	}
	return true
}

type Path interface {
	Append(p Point) Path
	Elements() []Point
	Contains(p Point) bool
}

// SlicePath is a naive implementation of a path: just a slice of Points.
type SlicePath []Point

// Append clones the underlying slice so that it can be modified.
func (sp SlicePath) Append(p Point) Path {
	return append(slices.Clone(sp), p)
}

func (sp SlicePath) Elements() []Point {
	return sp
}

func (sp SlicePath) Contains(p Point) bool {
	return slices.Contains(sp, p)
}

// ListPath represents a path as a linked list, but prefixes are shared, so
// multiple calls to Append with the same prefix create a tree. Append does not
// clone the preceeding path.
type ListPath struct {
	point Point
	next  *ListPath
}

func (lp *ListPath) Append(p Point) Path {
	return &ListPath{p, lp}
}

func (lp *ListPath) Elements() []Point {
	var ret []Point
	for lp != nil {
		ret = append(ret, lp.point)
		lp = lp.next
	}
	slices.Reverse(ret)
	return ret
}

func (lp *ListPath) Contains(p Point) bool {
	for lp != nil {
		if lp.point == p {
			return true
		}
		lp = lp.next
	}
	return false
}
