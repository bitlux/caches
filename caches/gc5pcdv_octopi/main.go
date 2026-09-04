package main

import (
	"fmt"

	"github.com/claygod/PiHex"
)

var byteToBits = map[byte]string{
	0:   "0000",
	1:   "0001",
	2:   "0010",
	3:   "0011",
	4:   "0100",
	5:   "0101",
	6:   "0110",
	7:   "0111",
	8:   "1000",
	9:   "1001",
	0xA: "1010",
	0xB: "1011",
	0xC: "1100",
	0xD: "1101",
	0xE: "1110",
	0xF: "1111",
}

var bitsToOctal = map[string]string{
	"000": "0",
	"001": "1",
	"010": "2",
	"011": "3",
	"100": "4",
	"101": "5",
	"110": "6",
	"111": "7",
}

func main() {
	pi := PiHex.New()
	bytes := pi.Get(3958428, 40)

	bits := ""
	for _, b := range bytes {
		bits += byteToBits[b]
	}

	for i := 0; i < len(bits)-2; i += 3 {
		fmt.Print(bitsToOctal[bits[i:i+3]])
	}
	fmt.Println()
}
