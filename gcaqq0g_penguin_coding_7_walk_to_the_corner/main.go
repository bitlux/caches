package main

import (
	"container/heap"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const gridStr = `48DDFFAFAF7CEFCBCDACFDE887C77DE402EB9E75
E4EEEF9DC59EEA7E34EAFDE7554E4F44FC4C454E
21FEEFDFDE67FEACE9A86FA932FDDFDFEEFEFF9B
0DC204EDEDBA9DCB7EC8EDE121EAD9ECFFFBEEFD
48417A55E9A99EAD775ACCF25DFFA9ACDEDDFA9B
BFC2BCF56DFFAF7B7D7FEFE0E8FEDECAFFEECDAC
EFFFDF53AEDA9D59E79BBAB434EDFABEAFEFBEDA
DDEFEC46CBBDEE9EFC9AEEDEAFDEDEEDCEFEDFFF
6CDAEB2E6FF9FC99CBBFFFCEA4EDDEACBDFFEEFF
5EDCFC04AEAFEDEF9EEEEBECEFDDDECDABFDEEFE
CBF5DEEBFCDDCAEBCBCBBFDFC5244FFEFCFDEEFF
5FEEECF4FDCC9C87C9EADFFBADDF4EDFDDFFDEFD
7CFD4449BDADEBFC5EA85BBDFDDF9FFFECFEFEFE
7F9E4FFFFFE8EE79766EF9FDDE994FEEF9FFFEDE
F8BE4DFFFFFEEFFFB89BDECCDCEBE415F9EEECDB
6EF49DDFEFDFEDDFEEDDAEDDEBCDEE94FEEDEDFD
5ADEFFEED9FFFDEEEFD7AEADCAB9DDB455DCFDBE
6AE4FFFDDEFDED7AAB6769FD8BEDEDFB8320D9AC
5DD72C2F79ACCEE96D58FABAED5BDFFCB964E9FF
AE7FDE0AB5956CA89FAD68EDEA8AACDEDA71BDBC
AD47B64DDFDE796F9BBF6D9EAF9CE8EDD425CCE9
DFEDB4FFF9CCF75EFE5FF77EC79EFA98D5CFCD9D
F57699FFF9BE88B65777EEECADFD89BCE204EBAC
D875C4E442053484F52677F6EEBEBEDEEEFE33FF
A66DBACDDB9CD99BD9544E755FAD88ADFFDFA3EB
9F3FC6567BFD8B59ABEB5F5B8CFCFEEEDFCDA720
AB7ACCBCABA5C6A7A7CE5F5041E7975FCFECAE43
951B95E98CDE655E899A3E0CB544859BEFDF8D81
FAC9CECAD9A5BCD8ABCA5429EDC926AFEF9D9B93
CEABE5AD9EAD9E9B5BFEDEFFFEFE0BEF32343E28
F97DC9D7885DC5E997F9DDFDDDFF457E5FDEAEC5
AB6CFEB5BDCF55CCB9FCEDDCCB8D68C52CD6C8C6
DEBD9AEE9CFFECAABABA1302453554750FEEFEEE
49502EEC6956FF7AEACE3EBEDBDD2FEFFFEFEFAF
1FCAE243FBA5869BDF9C2C9FB9CF4DDCFFFFFFCE
5FDCDFA6303E223030223FDBF994FFEFFEFE0538
999DBB9EECD8B788FCDB9DD9CBFDEEFEF44F2959
2FBE7AACBA0BDFD565E755CBFEF29DEC057EDF73
1E6F8ED924EE1917FBA59E6CFAB04E4529AB9F21`

type Point struct {
	row, col int
}

func (p Point) Add(q Point) Point {
	return Point{p.row + q.row, p.col + q.col}
}

type Grid struct {
	contents         string
	numRows, numCols int
}

func makeGrid(s string) Grid {
	g := Grid{contents: s}
	g.numRows = strings.Count(s, "\n") + 1
	g.numCols = strings.Index(s, "\n")
	return g
}

func (g Grid) At(p Point) int {
	offset := p.row*(g.numCols+1) + p.col
	s := g.contents[offset : offset+1]
	n, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		panic(err)
	}
	return int(n)
}

var grid Grid

type State struct {
	loc  Point
	cost int
	path []Point
}

func (s State) Dump() {
	for i := 0; i < len(s.path); i += 2 {
		hi := grid.At(s.path[i])
		lo := grid.At(s.path[i+1])
		fmt.Printf("%c", 16*hi+lo)
	}
	fmt.Println()
}

type Heap []State

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(State))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	grid = makeGrid(gridStr)
	visited := map[Point]bool{}

	h := &Heap{}
	heap.Init(h)
	start := Point{0, grid.numCols - 1}
	finish := Point{grid.numRows - 1, grid.numCols - 1}
	heap.Push(h, State{start, grid.At(start), []Point{start}})

	for h.Len() > 0 {
		curr := heap.Pop(h).(State)
		if visited[curr.loc] {
			continue
		}
		visited[curr.loc] = true

		if curr.loc == finish {
			curr.Dump()
			return
		}

		for _, offset := range []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			next := curr.loc.Add(offset)
			if next.row < 0 || next.row >= grid.numRows {
				continue
			}
			if next.col < 0 || next.col >= grid.numCols {
				continue
			}

			cost := curr.cost + grid.At(next)
			path := append(slices.Clone(curr.path), next)
			heap.Push(h, State{next, cost, path})
		}
	}
}
