package main

import (
	"fmt"
	"strings"
)

func penguncode(r rune) rune {
	m := map[rune]rune{
		'C': 'y',
		'D': 'g',
		'E': 'a',
		'F': 'z',
		'H': 'v',
		'I': 't',
		'J': 'o',
		'M': 'u',
		'O': 'n',
		'L': 's',
		'P': 'r',
		'Q': 'h',
		'T': 'd',
		'V': 'w',
		'X': 'e',
		'Y': 'i',
		'Z': 'p',
	}

	if ret, ok := m[r]; ok {
		return ret
	}
	return r
}

func main() {
	const ciphertext = "OJPIQ IQYPI CLXHX OTXDP XXLOY OXIXX OZJYO IJOXI VJFXP JVXLI JOXQM OTPXT EOTIV XOICI VJTXD PXXLI VJZJY OIIQP XXOYO XFXPJ"
	fmt.Println(ciphertext)
	fmt.Println(strings.Map(penguncode, ciphertext))
}
