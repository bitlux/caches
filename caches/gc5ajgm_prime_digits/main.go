package main

import (
	"fmt"

	"github.com/bitlux/caches/lib/util"
)

const MAX = 4000

func main() {
	var copelandErdos []int
	for i := 2; len(copelandErdos) < MAX; i++ {
		if util.IsPrime(i) {
			copelandErdos = append(copelandErdos, util.Digits(i)...)
		}
	}

	wanted := map[int]bool{
		2: true, 53: true, 59: true, 61: true, 71: true, 131: true, 157: true, 223: true, 277: true,
		617: true, 641: true, 1613: true, 1721: true, 2029: true, 2521: true}

	series := make([][]int, 10)
	for i := range 10 {
		series[i] = append(series[i], i)
	}
	for i := 2; i <= MAX; i++ {
		if util.IsPrime(i) {
			d := copelandErdos[i-1]
			series[d] = append(series[d], i)
		}
	}

	var coords []int

	for row := range 16 {
		for digit := range 10 {
			location := series[digit][row]
			if row > 0 && wanted[location] {
				coords = append(coords, digit)
			}
		}
	}

	fmt.Println(util.ToCoord(coords))
}
