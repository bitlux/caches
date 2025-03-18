package main

import (
	"fmt"
	"math"
)

var (
	A = 3
	B = 1
	C = 2
	D = 7
	E = 0
	F = 6
	G = 5
	H = 9
	I = -1
	J = -1
)

/*

6. The sixth rational approximation I will mention is BEIAIJ for the numerator and AACBG as the
denominator. This is slighly more accurate than the preceding approximation, but still only to 9 places past the decimal.

*/

func main() {
	for i := range 1000 {
		f := float64(i)
		if math.Abs(333/f-math.Pi) < .0001 {
			fmt.Println("BEF =", i)
		}
	}

	for i := range 10 {
		num := float64(300 + 11*i)
		if math.Abs(num/113-math.Pi) < .00001 {
			fmt.Println("G =", i)
		}
	}

	for i := range 10 {
		num := float64(103003 + 110*i)
		if math.Abs(num/33102-math.Pi) < .000000001 {
			fmt.Println("H =", i)
		}
	}

	for i := range 10 {
		for j := range 10 {
			num := float64(100300 + 1010*i + j)
			if math.Abs(num/33215-math.Pi) < .000000001 {
				fmt.Println("I J =", i, j)
			}
		}
	}
}
