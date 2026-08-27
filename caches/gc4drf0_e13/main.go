package main

import (
	"fmt"
	"hash/adler32"
	"os"
)

var alphabet = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'E', 'F',
	'G', 'H', 'I', 'L', 'M', 'N', 'O', 'R', 'S', 'T', 'U', 'W', 'X', 'Y'}

const target = 0x016700aa

func main() {
	candidate := make([]byte, 3)

	for _, a := range alphabet {
		candidate[0] = a
		for _, b := range alphabet {
			candidate[1] = b

			for _, c := range alphabet {
				candidate[2] = c

				csum := adler32.Checksum(candidate)
				if csum == target {
					fmt.Println(string(candidate))
					os.Exit(0)
				}
			}
		}
	}
}
