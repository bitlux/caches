package main

import (
	"fmt"
	"math/big"

	"github.com/bitlux/caches/util"
)

const modulus = 100_000_000

var fibs = make([]int, 0, 15_000_000)

func main() {
	fibs = append(fibs, 0, 1)
	for _, coord := range []string{
		"161803398874989484820458683436563811772030917980576286213544862270526046289040334",
		"3141592653589793238462643383279502884197169399375105820974944592307939959760"} {
		bn, ok := new(big.Int).SetString(coord, 10)
		util.MustBool(ok)

		bn.Mod(bn, big.NewInt(15_000_000))
		n64 := bn.Int64()

		for i := len(fibs); i <= int(n64); i++ {
			fibs = append(fibs, (fibs[i-1]+fibs[i-2])%modulus)
		}
		fmt.Println(fibs[n64])
	}
}
