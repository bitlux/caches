package main

import (
	"fmt"
	"strings"

	"github.com/bitlux/caches/lib/cipher"
	"github.com/bitlux/caches/lib/data"
	"github.com/bitlux/caches/lib/util"
)

const raw = `
AR AE XB QM CG CR EL CV WB SB
OB SW XQ PB HI UL CD HU QP HI
KD QO LM IS AX IP VX AY KG NY
QT HV QD DT EP PA HL XV KE SX
AQ ML SL FL PP QC DP HI CX OQ
FM FM TQ QS YB ZV PB AW YC CY
NS CX QS FM YC VX PM PH KR NY
PD XB YB ZV OG DP QY CY ZC QH
EP CW FX AQ HU NS CX QS NS PL
AW YC CY NS XP PM PH PD XB KE
ZV HW DA NP DP QY CY`

var ct string

func init() {
	ct = strings.ReplaceAll(strings.ReplaceAll(raw, "\n", ""), " ", "")
}

func main() {
	var max int
	var best *cipher.FourSquare

	for key1 := range data.OneK() {
		for key2 := range data.OneK() {
			c := cipher.NewFourSquare(key1, key2, false)
			dec, err := c.Decode(ct)
			util.Must(err)
			if score := data.BigramScore(dec); score > max {
				max = score
				best = c
				fmt.Printf("new best [%s, %s]: %d\n", key1, key2, max)
			}
		}
	}

	fmt.Println(best)
	fmt.Println(ct)
	fmt.Println(best.Decode(ct))
}
