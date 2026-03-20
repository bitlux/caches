package main

import (
	"fmt"
	"strconv"

	"github.com/bitlux/caches/lib/data"
	"github.com/bitlux/caches/lib/util"
)

func main() {
	digits := data.PiString()[2:]

	for targetLen := 1; targetLen <= 5; targetLen++ {
		for index := 0; ; index++ {
			candidate := digits[index : index+targetLen]
			n, err := strconv.Atoi(string(candidate))
			util.Must(err)

			if util.IsPrime(n) {
				fmt.Print(n)
				break
			}
		}
	}
	fmt.Println()
}
