package main

import (
	"fmt"
)

func uniq(ints ...int) bool {
	m := map[int]bool{}
	for _, n := range ints {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = true
	}
	return true
}

func main() {
	for a := 1; a <= 20; a++ {
		for b := 1; b <= 20; b++ {
			for c := 1; c <= 20; c++ {
				for d := 1; d <= 20; d++ {
					for e := 1; e <= 20; e++ {
						if b+c+d+e != 26 { // A
							continue
						}
						for f := 18; f <= 20; f++ {
							if a+d+f != 33 { // B
								continue
							}
							for g := 1; g <= 20; g++ {
								if a+d+g != 21 { // C
									continue
								}
								for h := 1; h <= 20; h++ {
									if c+d+h != 15 { // G
										continue
									}
									for i := 18; i <= 20; i++ {
										if f+i != 38 { // J
											continue
										}
										if g+i != 26 { // H
											continue
										}
										if a+b+c+g+i != 52 { // D
											continue
										}
										if a+f+i != 45 { // E
											continue
										}
										for j := 1; j <= 20; j++ {
											if b+e+j != 35 { // F
												continue
											}
											if d+e+h+j != 31 { // I
												continue
											}
											if !uniq(a, b, c, d, e, f, g, h, i, j) {
												continue
											}
											fmt.Printf("%c%c%c%c%c%c%c%c%c%c\n", a+64, b+64, c+64, d+64, e+64, f+64, g+64, h+64, i+64, j+64)
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
}
