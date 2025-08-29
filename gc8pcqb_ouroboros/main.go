package main

import (
	"fmt"
	"strconv"

	"github.com/bitlux/caches/util"
)

// http://combos.org/bruijn
func main() {
	var digits []int
	for _, o := range []struct {
		n   int
		seq string
		k   int
	}{
		{
			8,
			"redacted",
			125,
		},
		{
			9,
			"redacted",
			22,
		},
		{
			10,
			"redacted",
			947,
		},
		{
			7,
			"redacted",
			53,
		},
		{
			6,
			"redacted",
			37,
		},
		{
			11,
			"redacted",
			988,
		},
	} {
		sub := o.seq[o.k : o.k+o.n]
		n, err := strconv.ParseInt(sub, 2, 0)
		util.Must(err)
		digits = append(digits, util.Digits(n)...)
		fmt.Println(digits)
	}
	fmt.Println(util.ToCoord(digits))
}
