package main

import (
	"container/heap"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

const MaxPathCost = 509

var grid Grid

type Step struct {
	dir   Direction
	reset bool // Turn resets, continue and bounce do not.
}

type State struct {
	loc      Point
	dir      Direction
	cost     int
	path     Path
	slideLen int
}

func New(p Point) State {
	return State{
		loc:      p,
		dir:      SE,
		cost:     grid.At(p),
		path:     &ListPath{point: p, next: nil},
		slideLen: 0,
	}
}

func (s State) Clone(step Step, p Point) State {
	newState := State{
		loc:  p,
		dir:  step.dir,
		cost: s.cost + grid.At(p),
		path: s.path.Append(p),
	}
	if step.reset {
		newState.slideLen = 1
	} else {
		newState.slideLen = s.slideLen + 1
	}
	return newState
}

func (s State) Dump() {
	for i, row := range grid.contents {
		for j, char := range row {
			c := string(char)
			if s.path.Contains(Point{int8(i), int8(j)}) {
				c = color.GreenString(string(char))
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
}

type Heap []State

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) LockedLen(mu *sync.Mutex) int {
	mu.Lock()
	defer mu.Unlock()
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	switch {
	case h[i].cost < h[j].cost:
		return true
	case h[i].cost > h[j].cost:
		return false
	case h[i].loc.row > h[j].loc.row:
		return true
	case h[i].loc.col > h[j].loc.col:
		return true
	default:
		return false
	}
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *Heap) Push(x any) {
	*h = append(*h, x.(State))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func DirectPathCost(p Point) int {
	cost := 0
	for p.row < grid.numRows-1 && p.col < grid.numCols-1 {
		p = p.Move(SE)
		cost += grid.At(p)
	}

	// Move down along right side.
	for p.row < grid.numRows-1 {
		p = p.Move(SW)
		cost += grid.At(p)
		p = p.Move(SE)
		cost += grid.At(p)
	}

	// Move right along bottom.
	for p.col < grid.numCols-1 {
		p = p.Move(NE)
		cost += grid.At(p)
		p = p.Move(SE)
		cost += grid.At(p)
	}

	return cost
}

func init() {
	grid = Grid{contents: strings.Split(bigGrid, "\n")}
	grid.numRows = int8(len(grid.contents))
	grid.numCols = int8(len(grid.contents[0]))
}

func main() {
	start := Point{0, 0}
	finish := Point{grid.numRows - 1, grid.numCols - 1}

	var mu sync.Mutex
	h := &Heap{}
	heap.Init(h)
	heap.Push(h, New(start))
	ch := make(chan bool)

	// Stats
	maxCost := 0
	t0 := time.Now()

	for range 5 { // 5 goroutines
		go func() {
			for h.LockedLen(&mu) > 0 {
				mu.Lock()
				curr := heap.Pop(h).(State)
				mu.Unlock()

				if curr.cost+DirectPathCost(curr.loc) > MaxPathCost+200 {
					continue
				}

				mu.Lock()
				if curr.cost > maxCost && curr.cost > 100 {
					maxCost = curr.cost
					runtime.GC()
					fmt.Printf("%s %d\n", time.Since(t0).Round(time.Second), curr.cost)
				}
				mu.Unlock()

				if curr.slideLen > 10 {
					continue
				}

				if curr.loc == finish {
					if curr.slideLen >= 5 {
						fmt.Printf("\nFound optimal path with cost %s:\n", color.HiRedString("%d", curr.cost))
						curr.Dump()
						fmt.Println()
						ch <- true
						return
					} else {
						continue
					}
				}

				var steps []Step
				if curr.slideLen < 10 {
					steps = append(steps, Step{curr.dir, false})
				}
				if curr.slideLen >= 5 &&
					// Pablo cannot turn when on the edge, only bounce.
					curr.loc.row != 0 && curr.loc.row != grid.numRows-1 &&
					curr.loc.col != 0 && curr.loc.col != grid.numCols-1 {
					switch curr.dir {
					case NE:
						steps = append(steps, Step{SE, true}, Step{NW, true})
					case SE:
						steps = append(steps, Step{NE, true}, Step{SW, true})
					case SW:
						steps = append(steps, Step{SE, true}, Step{NW, true})
					case NW:
						steps = append(steps, Step{SW, true}, Step{NW, true})
					}
				}

				if curr.slideLen < 10 {
					switch curr.dir {
					case NE:
						if curr.loc.row == 0 {
							steps = append(steps, Step{SE, false})
						}
						if curr.loc.col == grid.numCols-1 {
							steps = append(steps, Step{NW, false})
						}
					case SE:
						if curr.loc.row == grid.numRows-1 {
							steps = append(steps, Step{NE, false})
						}
						if curr.loc.col == grid.numCols-1 {
							steps = append(steps, Step{SW, false})
						}
					case SW:
						if curr.loc.row == grid.numRows-1 {
							steps = append(steps, Step{NW, false})
						}
						if curr.loc.col == 0 {
							steps = append(steps, Step{SE, false})
						}
					case NW:
						if curr.loc.row == 0 {
							steps = append(steps, Step{SW, false})
						}
						if curr.loc.col == 0 {
							steps = append(steps, Step{NE, false})
						}
					}
				}

				for _, step := range steps {
					next := curr.loc.Move(step.dir)
					if !grid.IsValid(next) {
						continue
					}
					if curr.path.Contains(next) {
						continue
					}
					mu.Lock()
					heap.Push(h, curr.Clone(step, next))
					mu.Unlock()
				}
			}
		}()
	}

	<-ch
}
