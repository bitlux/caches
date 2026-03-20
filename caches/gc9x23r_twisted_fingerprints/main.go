package main

import (
	"fmt"

	"github.com/bitlux/caches/lib/cipher"
	"github.com/bitlux/caches/lib/data"
	"github.com/bitlux/caches/lib/util"
)

func main() {
	m := map[int]string{}
	words := data.Large()
	for w := range words {
		m[util.FromDigits(cipher.CBF(w))] = w
	}

	for _, target := range []int{4896058, 3853116, 3851947, 2166184, 3856658, 321398522, 321998553,
		29954, 591935, 3546583990, 3546945435, 358585435, 3536145, 3547857105, 3543188540} {
		fmt.Println(m[target])
	}
}
