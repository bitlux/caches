package main

import (
	"fmt"

	"github.com/bitlux/caches/lib/data"
	"github.com/bitlux/caches/lib/util"
)

func main() {
	tau := data.TauDigits()

	// Stage 1
	var primes []byte
	for _, ln := range []int{1, 2, 3, 4, 5} {
		for i := range tau {
			n := tau[i : i+ln]
			if util.IsPrime(util.FromDigits(n)) {
				primes = append(primes, n...)
				break
			}
		}
	}
	fmt.Println(util.ToCoord(util.Digits(util.FromDigits(primes) + 89161547483435)))

	// Stage 2
	start := 0
	length := 1
	var roots []int
	for range 15 {
		n := util.FromDigits(tau[start : start+length])
		dr := util.DigitalRoot(n)
		roots = append(roots, dr)
		start = start + length
		length = dr
	}
	fmt.Println(util.ToCoord(util.Digits(util.FromDigits(roots) - 319293325722075)))

	// Stage 3
	var dets []int
	corr := []int{35, 45, 11, -1, 49, 35, 17, 95}
	for i := range 8 {
		det := (int(tau[i]) * int(tau[i+3])) - (int(tau[i+2]) * int(tau[i+1]))
		dets = append(dets, det+corr[i])
	}

	var digs []int
	for _, d := range dets {
		digs = append(digs, util.Digits(d)...)
	}
	fmt.Println(util.ToCoord(digs))
}
