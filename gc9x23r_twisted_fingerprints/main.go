package main

import (
	"fmt"

	"github.com/bitlux/caches/util"
)

func main() {
	m := map[int]string{}
	words := util.ReadLines("../data/wordlist-20210729.txt")
	for _, w := range words {
		m[util.FromDigits(util.CBF(w))] = w
	}

	for _, target := range []int{4896058, 3853116, 3851947, 2166184, 3856658, 321398522, 321998553,
		29954, 591935, 3546583990, 3546945435, 358585435, 3536145, 3547857105, 3543188540} {
		fmt.Println(m[target])
	}
}
