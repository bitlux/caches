package main

import (
	"fmt"

	"github.com/bitlux/caches/lib/cipher"
	"github.com/bitlux/caches/lib/util"
)

func main() {
	fmt.Println(util.FromDigits(cipher.CBF("wgvclpe")))
	fmt.Println(util.FromDigits(cipher.CBF("ubadrksw")))
}
