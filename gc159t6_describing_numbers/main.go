package main

import (
	"fmt"

	"github.com/bitlux/caches/util"
)

func square(n int) int {
	return n * n
}

func main() {
	for A := range 10 {
		for B := range 10 {
			for C := range 10 {
				ABC := 100*A + 10*B + C
				XYZ := square(A + B + C)
				if XYZ < 100 {
					continue
				}
				// fmt.Println(A, B, C, XYZ)
				digits := util.Digits(XYZ)
				if ABC-A-B-C == square(digits[0]+digits[1]+digits[2]) {
					fmt.Println(A, B, C, " ", digits[0], digits[1], digits[2])
				}
			}
		}
	}
}
