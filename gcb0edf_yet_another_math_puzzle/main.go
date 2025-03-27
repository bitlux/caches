package main

import (
	"fmt"
)

const needle = 0.5463865691989018

func concat(args ...int) float64 {
	r := args[0]
	for i := 1; i < len(args); i++ {
		r = r*10 + args[i]
	}
	return float64(r)
}

func main() {
	n0 := 3
	n1 := 7
	n2 := 2

	w0 := 1
	w1 := 2
	w2 := 1
	w3 := 5

	for n3 := 2; n3 <= 4; n3++ {
		for n4 := range 10 {
			for n5 := range 10 {
				for n6 := range 10 {
					nf := concat(n0, n1, n2, n3, n4, n5, n6)
					nb := concat(n6, n5, n4, n3, n2, n1, n0)
					for w4 := range 10 {
						for w5 := range 10 {
							for w6 := range 10 {
								for w7 := range 10 {
									wf := concat(w0, w1, w2, w3, w4, w5, w6, w7)
									wb := concat(w7, w6, w5, w4, w3, w2, w1, w0)
									sum := nf/nb + wf/wb

									if sum == needle {
										fmt.Printf("N %d%d %d%d.%d%d%d W %d%d%d %d%d.%d%d%d\n", n0, n1, n2, n3, n4, n5, n6, w0, w1, w2, w3, w4, w5, w6, w7)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
