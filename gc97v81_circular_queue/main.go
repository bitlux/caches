package main

import (
	"fmt"

	"github.com/bitlux/caches/util"
)

func main() {
	offsets := []int{2, 1, 1, -1, 0, 0, -2, -8, -1, -3, 0, 5, -5, 1, -1}
	cbf := util.CBF("APANDEMICNOMORE")
	for i := range len(cbf) {
		fmt.Println(cbf[i] + offsets[i])
	}
}
