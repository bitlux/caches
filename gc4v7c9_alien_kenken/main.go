package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/bitlux/caches/util"
	"github.com/gnboorse/centipede"
)

func round(n float64) int {
	return int(math.Round(n - 0.5))
}

func clubs(vars *centipede.Variables[int], names centipede.VariableNames) int {
	sum := 0
	for _, name := range names {
		n := vars.Find(name).Value
		sum += n * n
	}

	return sum
}

func hearts(vars *centipede.Variables[int], names centipede.VariableNames) int {
	sum := 0
	for _, name := range names {
		n := vars.Find(name).Value
		sum += n * n * n
	}

	return round(math.Sqrt(float64(sum)))
}

func spades(vars *centipede.Variables[int], names centipede.VariableNames) int {
	sum := 0
	prod := 1
	for _, name := range names {
		n := vars.Find(name).Value
		sum += n
		prod *= n
	}
	return sum + round(math.Sqrt(float64(prod)))
}

func diamonds(vars *centipede.Variables[int], names centipede.VariableNames) int {
	sum := 0
	prod := 1
	denom := 0.0
	for _, name := range names {
		n := vars.Find(name).Value
		sum += n
		prod *= n
		denom += 1.0 / float64(n)
	}

	n := float64(len(names))
	nthRoot := math.Pow(float64(prod), 1.0/n)
	return round(n * (nthRoot + float64(sum)/n + n/denom))
}

const SIZE = 9

func containsAll(vars *centipede.Variables[int], names centipede.VariableNames) bool {
	for _, name := range names {
		if vars.Find(name).Empty {
			return false
		}
	}
	return true
}

func dump(state centipede.CSPState[int]) {
	for i := range SIZE {
		row := i + 1
		for j := range SIZE {
			col := j + 1
			variable := state.Vars.Find(pair(col, row))
			fmt.Printf("%d ", variable.Value)
		}
		fmt.Println()
	}
}

func pair(i, j int) centipede.VariableName {
	return centipede.VariableName(string(util.A1Z26(i)) + strconv.Itoa(j))
}

func main() {
	ctx := context.Background()

	var vars centipede.Variables[int]
	var constraints []centipede.Constraint[int]

	// Create variables, enforce uniqueness in a row.
	for i := range SIZE {
		var inRow []centipede.VariableName
		for j := range SIZE {
			name := pair(j+1, i+1)
			vars = append(vars, centipede.NewVariable(name, centipede.IntRange(1, SIZE+1)))
			inRow = append(inRow, name)
		}
		//		constraints = append(constraints, centipede.AllUnique[int](inRow...)...)
	}

	// Enforce uniqueness in a column.
	for j := range SIZE {
		var inCol []centipede.VariableName
		for i := range SIZE {
			inCol = append(inCol, pair(j+1, i+1))
		}
		//		constraints = append(constraints, centipede.AllUnique[int](inCol...)...)
	}

	// TODO: remove
	for _, row := range []int{ /*1, 5, */ 9} {
		var inRow []centipede.VariableName
		for col := range SIZE {
			inRow = append(inRow, pair(row, col+1))
		}
		constraints = append(constraints, centipede.AllUnique[int](inRow...)...)
	}

	for row := range SIZE {
		var inCol []centipede.VariableName
		for _, col := range []int{1 /*5,*/, 9} {
			inCol = append(inCol, pair(row+1, col))
		}
		constraints = append(constraints, centipede.AllUnique[int](inCol...)...)
	}
	// end TODO:

	for _, c := range []struct {
		f      func(*centipede.Variables[int], centipede.VariableNames) int
		target int
		names  centipede.VariableNames
	}{
		//		{clubs, 237, centipede.VariableNames{"A1", "B1", "C1", "A2", "A3"}},
		//		{diamonds, 29, centipede.VariableNames{"D1", "E1", "F1"}},
		//		{diamonds, 72, centipede.VariableNames{"G1", "H1", "I1", "I2", "I3"}},

		//		{hearts, 36, centipede.VariableNames{"B2", "C2", "D2", "B3", "B4"}},
		// redundant with domain
		//		{diamonds, 13, centipede.VariableNames{"E2", "E3"}},
		//		{hearts, 34, centipede.VariableNames{"F2", "G2", "H2", "H3", "H4"}},

		//		{hearts, 39, centipede.VariableNames{"C3", "D3", "C4", "D4"}},
		//		{diamonds, 20, centipede.VariableNames{"F3", "G3", "F4", "G4"}},

		//	{hearts, 9, centipede.VariableNames{"A4", "A5", "A6"}},
		// {spades, 104, centipede.VariableNames{"E4", "D5", "E5", "F5", "E6"}},
		//		{hearts, 18, centipede.VariableNames{"I4", "I5", "I6"}},

		// redundant with domain
		// {clubs, 34, centipede.VariableNames{"B5", "C5"}},
		// redundant with domain
		// {spades, 15, centipede.VariableNames{"G5", "H5"}},

		//	{clubs, 130, centipede.VariableNames{"B6", "B7", "B8", "C8", "D8"}},
		//	{spades, 21, centipede.VariableNames{"C6", "D6", "C7", "D7"}},
		//	{diamonds, 79, centipede.VariableNames{"F6", "G6", "F7", "G7"}},
		//	{clubs, 91, centipede.VariableNames{"H6", "H7", "F8", "G8", "H8"}},

		{diamonds, 84, centipede.VariableNames{"A7", "A8", "A9", "B9", "C9"}},
		// redundant with domain
		// {hearts, 28, centipede.VariableNames{"E7", "E8"}},
		{hearts, 35, centipede.VariableNames{"I7", "I8", "G9", "H9", "I9"}},

		{spades, 24, centipede.VariableNames{"D9", "E9", "F9"}},
	} {
		constraints = append(constraints, centipede.Constraint[int]{
			Vars: c.names,
			ConstraintFunction: func(vars *centipede.Variables[int]) bool {
				if !containsAll(vars, c.names) {
					return true
				}
				return c.f(vars, c.names) == c.target
			},
		})
	}

	// There must be two 9's in 237, not in A1
	constraints = append(constraints, centipede.Constraint[int]{
		Vars: []centipede.VariableName{"B1", "C1", "A2", "A3"},
		ConstraintFunction: func(vars *centipede.Variables[int]) bool {
			names := []centipede.VariableName{"B1", "C1", "A2", "A3"}
			if !containsAll(vars, names) {
				return true
			}
			count := 0
			for _, name := range names {
				if vars.Find(name).Value == 9 {
					count++
				}
			}
			return count == 2
		},
	})

	vars.SetDomain("A1", centipede.Domain[int]{1, 5})
	vars.SetDomain("B1", centipede.Domain[int]{1, 5, 9})
	vars.SetDomain("C1", centipede.Domain[int]{1, 5, 9})
	vars.SetDomain("D1", centipede.Domain[int]{2, 3, 6})
	vars.SetDomain("E1", centipede.Domain[int]{3, 6})
	vars.SetDomain("F1", centipede.Domain[int]{2, 3, 6})
	vars.SetDomain("G1", centipede.Domain[int]{4, 7})
	vars.SetDomain("H1", centipede.Domain[int]{4, 7})
	vars.SetDomain("I1", centipede.Domain[int]{8})

	vars.SetDomain("A2", centipede.Domain[int]{7, 9})
	vars.SetDomain("E2", centipede.Domain[int]{1, 5})
	vars.SetDomain("I2", centipede.Domain[int]{3, 4})

	vars.SetDomain("A3", centipede.Domain[int]{7, 9})
	vars.SetDomain("E3", centipede.Domain[int]{1, 5})
	vars.SetDomain("I3", centipede.Domain[int]{3, 4})

	vars.SetDomain("A4", centipede.Domain[int]{1, 2, 3})
	vars.SetDomain("E4", centipede.Domain[int]{2, 8})
	vars.SetDomain("I4", centipede.Domain[int]{2, 7})

	vars.SetDomain("A5", centipede.Domain[int]{4})
	vars.SetDomain("B5", centipede.Domain[int]{3, 5})
	vars.SetDomain("C5", centipede.Domain[int]{3, 5})
	vars.SetDomain("D5", centipede.Domain[int]{6, 7, 8})
	vars.SetDomain("E5", centipede.Domain[int]{6, 7})
	vars.SetDomain("F5", centipede.Domain[int]{6, 7, 8})
	vars.SetDomain("G5", centipede.Domain[int]{2, 9})
	vars.SetDomain("H5", centipede.Domain[int]{2, 9})
	vars.SetDomain("I5", centipede.Domain[int]{1})

	vars.SetDomain("A6", centipede.Domain[int]{1, 2, 3})
	vars.SetDomain("E6", centipede.Domain[int]{2, 8})
	vars.SetDomain("I6", centipede.Domain[int]{2, 7})

	vars.SetDomain("A7", centipede.Domain[int]{2, 5, 6, 8})
	vars.SetDomain("E7", centipede.Domain[int]{4, 9})
	vars.SetDomain("I7", centipede.Domain[int]{5, 6, 9})

	vars.SetDomain("A8", centipede.Domain[int]{2, 5, 6, 8})
	vars.SetDomain("E8", centipede.Domain[int]{4, 9})
	vars.SetDomain("I8", centipede.Domain[int]{5, 6, 9})

	vars.SetDomain("A9", centipede.Domain[int]{2, 5, 6, 8})
	vars.SetDomain("E9", centipede.Domain[int]{3, 6, 7})
	vars.SetDomain("G9", centipede.Domain[int]{1, 2, 4, 5, 6})
	vars.SetDomain("H9", centipede.Domain[int]{1, 2, 4, 5, 6})
	vars.SetDomain("I9", centipede.Domain[int]{5, 6, 9})

	solver := centipede.NewBackTrackingCSPSolver(vars, constraints)
	solver.State.MakeArcConsistent(ctx)

	t := time.Now()
	go func() {
		for range time.Tick(10 * time.Second) {
			fmt.Println("Running for", time.Since(t))
			fmt.Println(solver.State.Failures, "failures")
			dump(solver.State)
			fmt.Println()
		}
	}()

	success, err := solver.Solve(ctx)
	util.Must(err)
	util.MustBool(success)

	dump(solver.State)

	x := clubs(&solver.State.Vars, centipede.VariableNames{"B6", "C9", "I3", "E3"})
	y := clubs(&solver.State.Vars, centipede.VariableNames{"F9", "D4", "G6", "B2", "G4", "F1"})
	fmt.Printf("N 37 23.%03d W 122 08.%03d", x, y)
}
