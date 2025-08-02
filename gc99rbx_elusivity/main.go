package main

import (
	"fmt"

	"github.com/bitlux/caches/util"
)

func index(s string, b byte, n byte) int {
	var count byte = 0
	for i := range len(s) {
		if s[i] == b {
			count++
			if count == n {
				return i
			}
		}
	}
	return -1
}

const (
	alphabet   = "OFANYWEIRDELUSIVEGHOSTFISH"
	ciphertext = `W1I2G1  O2I1Y1  W1R1S1O1L1  E1H1  Y1S1O2Y1G1
S1I2G1O2I1  W1I2S2G1  H2Y1G1I2  O2F2Y1S1O2S3
N1I2O2  H1Y1T1Y1S1  Y1L1Y1T1Y1S1  F2Y1H1O2
O2F2Y1L1T1Y1  I2S1Y1  O2I1R1G1O2S3  W1R1T1Y1
N1I2O2  W1I2G1O2S3  S1R1S1Y1  Y1R1E1I1O2`
)

func main() {
	for i := 0; i < len(ciphertext); {
		c := ciphertext[i]
		if c >= 'A' && c <= 'Z' {
			n := ciphertext[i+1] - '0'
			ind := index(alphabet, c, n) + 1
			fmt.Printf("%c", util.A1Decode(ind))
			i += 2
		} else {
			fmt.Printf(" ")
			i++
		}
	}
	fmt.Println()
}
