package main

import (
	"fmt"

	"github.com/bitlux/caches/util"
)

func main() {
	x := 1
	for x%3 != 1 || x%5 != 2 || x%11 != 10 || x%17 != 9 || x%23 != 19 || x%29 != 10 || x%37 != 5 {
		x++
	}

	y := 1
	for y%7 != 1 || y%13 != 1 || y%19 != 10 || y%31 != 23 || y%41 != 21 || y%43 != 20 {
		y++
	}

	fmt.Println(util.ToCoord(append(util.Digits(x), util.Digits(y)...)))
}
