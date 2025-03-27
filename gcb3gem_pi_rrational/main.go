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
	E int
	F int
	G int
	H int
	I int
	J int
)

func main() {
	for i := range 1000 {
		f := float64(i)
		if math.Abs(333/f-math.Pi) < .0001 {
			E = (i / 10) % 10
			F = i % 10
			break
		}
	}

	for i := range 10 {
		f := float64(300 + 11*i)
		if math.Abs(f/113-math.Pi) < .00001 {
			G = i
			break
		}
	}

	for i := range 10 {
		f := float64(103003 + 10000*E + 110*i)
		if math.Abs(f/(33102+10*float64(E))-math.Pi) < .000000001 {
			H = i
			break
		}
	}

Outer:
	for i := range 10 {
		for j := range 10 {
			num := float64(100300 + 1010*i + j)
			if math.Abs(num/(33210+float64(G))-math.Pi) < .000000001 {
				I = i
				J = j
				break Outer
			}
		}
	}

	fmt.Printf("N 37 %d%d.%d%d%d W 121 %d%d.%d%d%d\n", C, C, J, C, B, G, H, J, A, E)

}
