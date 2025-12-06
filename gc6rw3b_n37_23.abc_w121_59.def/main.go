package main

import (
	"fmt"
	"strings"

	"github.com/bitlux/caches/util"
)

func main() {
	text := "Frequently puzzles produce only 6 digits (ABCDEF). Count me out on liking this encoding but maybe you do"
	counts := util.RuneCount(strings.ToLower(text))
	fmt.Println(util.ToCoord([]int{3, 7, 2, 3, counts['a'], counts['b'], counts['c'], 1, 2, 1, 5, 9, counts['d'], counts['e'], counts['f']}))
}
