package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bitlux/caches/util"
)

// cbf encodes a string into a number
func cbf(s string) string {
	var ret string
	s = strings.ToLower(s)
	for _, c := range s {
		digit := ((c - 'a' + 1) % 10)
		ret += strconv.Itoa(int(digit))
	}
	return ret
}

func main() {
	m := map[string]string{}
	words := util.ReadLines("../data/wordlist-20210729.txt")
	for _, w := range words {
		m[cbf(w)] = w
	}

	for _, target := range []string{"4896058", "3853116", "3851947", "2166184", "3856658", "321398522",
		"321998553", "29954", "591935",
		"3546583990", "3546945435", "358585435",
		"3536145", "3547857105", "3543188540"} {
		fmt.Println(m[target])
	}
}
