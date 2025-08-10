package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/bitlux/caches/util"
)

func main() {
	fake := util.ReadLines("pi.htm")[0]
	f, _, err := big.ParseFloat(fake, 10, 500_000, big.ToNearestEven)
	util.Must(err)

	var base11 strings.Builder
	for range len(fake) {
		i, _ := f.Int64()
		if i == 10 {
			base11.WriteString("A")
		} else {
			s := strconv.FormatInt(i, 10)
			base11.WriteString(s)
		}
		f.Sub(f, new(big.Float).SetInt64(i))
		f.Mul(f, new(big.Float).SetInt64(11))
	}

	binary := base11.String()[28735:30944]
	for i := 0; i < len(binary)/47; i++ {
		fmt.Println(binary[i*47 : (i+1)*47])
	}
}
