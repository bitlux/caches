package main

import (
	"fmt"
	"os"

	"github.com/bitlux/caches/util"
)

func main() {
	fmt.Print("       ")
	for _, w := range os.Args[1:] {
		fmt.Printf("%s ", w)
	}
	fmt.Println()

	fmt.Print("A1Z26: ")
	for _, w := range os.Args[1:] {
		for _, r := range w {
			fmt.Printf("%d ", util.A1Encode(r))
		}
	}
	fmt.Println()
}
