package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/bitlux/caches/lib/datastructures"
	"github.com/bitlux/caches/lib/util"
)

// p prints a grid
func p(g [][]byte) {
	for i := range 5 {
		for j := range 5 {
			fmt.Print(g[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func reversed(s []byte) []byte {
	return []byte{s[4], s[3], s[2], s[1], s[0]}
}

func equal(a, b []byte) bool {
	return slices.Equal(a, b) || slices.Equal(a, reversed(b))
}

func bidiContains(t *datastructures.Trie, s []byte) bool {
	return t.Contains(s) || t.Contains(reversed(s))
}

func main() {
	ODDS := []byte{1, 3, 5, 7, 9}

	// for iteration, contains primes both backwards and forwards
	candidates := [][]byte{}
	// for lookup
	primes := datastructures.NewTrie()

	for _, a := range ODDS {
		for _, b := range ODDS {
			for _, c := range ODDS {
				for _, d := range ODDS {
					for _, e := range ODDS {
						s := []byte{a, b, c, d, e}
						if n := util.FromDigits(s); util.IsPrime(n) {
							if nRev := util.FromDigits([]byte{e, d, c, b, a}); n != nRev && util.IsPrime(nRev) {
								candidates = append(candidates, s)
								primes.Insert(s)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(len(candidates), "primes")

	start := time.Now()
	count := 0

	grid := [][]byte{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}

	for i, a := range candidates {
		fmt.Printf("%d / %d  %s\n", i, len(candidates), time.Since(start))
		grid[0] = a

		for _, e := range candidates {
			if a[0] != e[4] || a[4] != e[0] {
				continue
			}
			if equal(a, e) {
				continue
			}

			grid[4] = e

			for _, colA := range candidates {
				if colA[0] != a[0] || colA[4] != e[0] {
					continue
				}
				if equal(a, colA) || equal(e, colA) {
					continue
				}

				grid[1][0] = colA[1]
				grid[2][0] = colA[2]
				grid[3][0] = colA[3]

				for _, colB := range candidates {
					if colB[0] != a[1] || colB[4] != e[1] {
						continue
					}
					if equal(a, colB) || equal(e, colB) || equal(colA, colB) {
						continue
					}

					grid[1][1] = colB[1]
					grid[2][1] = colB[2]
					grid[3][1] = colB[3]

					for _, colC := range candidates {
						if colC[0] != a[2] || colC[4] != e[2] {
							continue
						}
						if equal(a, colC) || equal(e, colC) || equal(colA, colC) || equal(colB, colC) {
							continue
						}

						grid[1][2] = colC[1]
						grid[2][2] = colC[2]
						grid[3][2] = colC[3]

						for _, colD := range candidates {
							if colD[0] != a[3] || colD[4] != e[3] {
								continue
							}
							if equal(a, colD) || equal(e, colD) || equal(colA, colD) || equal(colB, colD) || equal(colC, colD) {
								continue
							}

							grid[1][3] = colD[1]
							grid[2][3] = colD[2]
							grid[3][3] = colD[3]

						E:
							for _, colE := range candidates {
								if colE[0] != a[4] || colE[4] != e[4] {
									continue
								}

								grid[1][4] = colE[1]
								grid[2][4] = colE[2]
								grid[3][4] = colE[3]

								d1 := []byte{grid[0][0], grid[1][1], grid[2][2], grid[3][3], grid[4][4]}
								d2 := []byte{grid[0][4], grid[1][3], grid[2][2], grid[3][1], grid[4][0]}

								if !bidiContains(primes, grid[1]) || !bidiContains(primes, grid[2]) ||
									!bidiContains(primes, grid[3]) || !bidiContains(primes, d1) ||
									!bidiContains(primes, d2) {
									continue
								}

								// Now the grid is complete, check all rows and column for uniqueness.
								gridSet := datastructures.NewTrie()
								for _, s := range [][]byte{
									grid[0], grid[1], grid[2], grid[3], grid[4], colA, colB, colC, colD, colE, d1, d2,
								} {
									if bidiContains(gridSet, s) {
										continue E
									}
									// Because bidiContains checks both directions, we need to insert only one direction.
									gridSet.Insert(s)
								}

								count++
								p(grid)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(count)
}
