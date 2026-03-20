package main

import "fmt"

var targets = [][]int{
	{1920, 768},
	{4608, 1472},
	{1792, 768},
	{3008, 3648},
}

func f(x, y int) (int, int) {
	for range 12 {
		x, y = x+y, x-y
	}
	return x, y
}

func main() {
OUTER:
	for _, t := range targets {
		for x := range 1000 {
			for y := range 1000 {
				a, b := f(x, y)
				if a == t[0] && b == t[1] {
					fmt.Println(x, y)
					continue OUTER
				}
			}
		}
	}
}
