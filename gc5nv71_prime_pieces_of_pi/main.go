package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bitlux/caches/util"
)

func main() {
	digits, err := os.ReadFile("pi")
	util.Must(err)
	digits = digits[2:]

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
