package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, c := range []string{"A37841", "26B3B11"} {
		i, _ := strconv.ParseInt(c, 13, 32)
		fmt.Println(i)
	}
}
