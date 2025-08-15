package main

import (
	"fmt"
	"strings"

	"github.com/bitlux/caches/util"
)

func main() {
	words := util.ReadLines("../data/wordlist-20210729.txt")

	for _, codeword := range []string{
		"mjdau", "apgla", "zlbsj", "saqdq", "hsjvq", "hyxub", "qscaf", "fshks", "qktul", "lvxm",
		"suzla", "foawd", "scdjv", "daysj", "fipsf",
		// "YFWLA", "NEJGFLK"
	} {
		codeword = strings.ToLower(codeword)
		fmt.Println(codeword)
		for n := range 25 {
			rotten := util.ROT(n+1, codeword)
			for _, term := range words {
				if strings.Contains(term, rotten) {
					fmt.Printf(" %2d %-5s %-15s %d\n", n+1, rotten, term, util.CBF(term[:1])[0])
				}
			}
		}
	}
}
