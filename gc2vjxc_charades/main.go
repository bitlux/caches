package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, s := range []string{
		"245B71",
		"840C5B",
		"3C7238",
		"491AA8",
		"515541",
		"494C46",
		"10304E",
		"E0F57B",
	} {
		i, _ := strconv.ParseInt(s, 16, 32)
		fmt.Printf("%024s\n", strconv.FormatInt(i, 2))
	}

	for _, s := range []string{
		"000000000001",
		"000000100000",
		"000001000111",
		"000000011001",
		"000000110101",
		"000001010100",
		"000000100111",
		"000000001111",
	} {
		i, _ := strconv.ParseInt(s, 2, 32)
		fmt.Printf("%2s\n", strconv.FormatInt(i, 10))
	}
}
