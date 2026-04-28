package main

import (
	"fmt"
	"math"

	"github.com/bitlux/caches/lib/util"
)

func digit(index int) int {
	for i := 1; ; i++ {
		span := 9 * int(math.Pow10(i-1)) * i

		if span >= index {
			num := int(math.Pow10(i-1)) + (index+i-1)/i - 1
			return util.Digits(num)[(index-1)%i]
		}
		index -= span
	}
}

func main() {
	var coords []int
	for _, i := range []int{
		37412, 87654320, 3738, 60264027, 77777809,
		8777777787, 24910780, 38890, 616163, 19888834,
		488894, 30770774, 680309142, 2221, 86527065} {
		coords = append(coords, digit(i))
	}
	fmt.Println(util.ToCoord(coords))
}
