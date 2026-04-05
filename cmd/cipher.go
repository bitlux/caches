package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bitlux/caches/lib/cipher"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	fmt.Println(args)

	fmt.Println("\nA1Z26")
	for _, r := range args {
		fmt.Printf("%d ", cipher.A1Encode(r))
	}

	fmt.Println("\n\nROT")
	for i := 1; i < 26; i++ {
		fmt.Printf("%2d %s\n", i, cipher.ROT(i, args))
	}

}
