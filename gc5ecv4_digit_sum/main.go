package main

import (
	"fmt"

	"github.com/bitlux/caches/util"
)

func sum(body []byte) int {
	sum := 0
	for _, b := range body {
		sum += int(b) - 48
	}
	return sum
}

func main() {
	lat := util.Wget("http://techmanski.net/geocaching/GC5ECV4/latitude.txt")
	fmt.Println(sum(lat))
	long := util.Wget("http://techmanski.net/geocaching/GC5ECV4/longitude.txt")
	fmt.Println(sum(long) * 3)
}
