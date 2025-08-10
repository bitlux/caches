package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Float32frombits(0x42181c8d))
	fmt.Println(math.Float32frombits(0xc2f51a8a))
}
