package main

import (
	"fmt"

	"github.com/bitlux/caches/data"
	"github.com/bitlux/caches/util"
)

var floats = []float64{
	2.0544150110375275,
	4.3662254361879629e-4,
	8.4984670371344170e-8,
	9.3814246508023106e-12,
	5.5184850887057188e-14,
}

func main() {
	words := make([]string, 5)

	var threes []string
	dict := data.LargeSet()
	for word := range dict {
		if len(word) == 3 {
			threes = append(threes, word)
		}
	}

OUTER:
	for _, w1 := range threes {
		for l4 := 'a'; l4 <= 'z'; l4++ {
			for l5 := 'a'; l5 <= 'z'; l5++ {
				w2 := fmt.Sprintf("%s%c%c", w1, l4, l5)
				if !dict[w2] {
					continue
				}
				for l6 := 'a'; l6 <= 'z'; l6++ {
					w3 := fmt.Sprintf("%s%c", w2, l6)
					if !dict[w3] {
						continue
					}

					for l7 := 'a'; l7 <= 'z'; l7++ {
						w4 := fmt.Sprintf("%s%c%c", w3, w1[1], l7)
						if !dict[w4] {
							continue
						}
						for l8 := 'a'; l8 <= 'z'; l8++ {
							w5 := fmt.Sprintf("%s%c", w4, l8)
							if !dict[w5] {
								continue
							}

							if util.SHA256(w5) == "f350895acb597dcc8a039060a9ce12f0c9c6504f51c2bfc25034f04c4c786675" {
								words = []string{w1, w2, w3, w4, w5}
								break OUTER
							}
						}
					}
				}
			}
		}
	}

	var coords []int
	for i, w := range words {
		var digits []int
		for _, r := range w {
			digits = append(digits, util.Digits(util.A1Encode(r))...)
		}
		n := util.FromDigits(digits)
		coords = append(coords, util.Digits(int64(float64(n)*floats[i]))...)
	}
	fmt.Println(util.ToCoord(coords))
}
