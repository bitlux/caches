package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/bitlux/caches/lib/cipher"
	"github.com/bitlux/caches/lib/data"
	"github.com/bitlux/caches/lib/util"
)

const ciphertext = `VU SA EE TN ON FW RE FD RI MG EZ NE KN EZ QF KF ZP RI LC AE MV UE BA CB WG GF DD LK QQ XL NX MG OC UO QF KF GP RE LX XP HO QW WG TP EA TR RE ED UR DS EW GP TN SC ON WG PV ES VP NX CD QW HF TY RE SM GP HM GL TO BX NB KK QF TS RE SM GP`

var runes = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'K', 'L',
	'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func findCycle(length, numDigits int) map[int]int {
	pi := data.PiString()[1:]
	start := int(math.Pow10(numDigits - 1))
	end := int(math.Pow10(numDigits)) - 1

	values := map[int]int{}

	for i := start; i <= end; i++ {
		n, err := strconv.Atoi(pi[i : i+numDigits])
		util.Must(err)
		values[i] = n
	}

	for start, next := range values {
		seen := map[int]int{start: 1}
		for i := 2; ; i++ {
			if repeat, ok := seen[next]; ok {
				if next == start && i-repeat == length {
					return seen
				}
				break
			}
			seen[next] = i
			var ok bool
			next, ok = values[next]
			if !ok {
				break
			}
		}
	}
	return nil
}

func sortByValues(m map[int]int) []int {
	var ret []int
	for i := 1; i <= len(m); i++ {
		for k, v := range m {
			if v == i {
				ret = append(ret, k)
				break
			}
		}
	}

	start := slices.Min(ret)
	for ret[0] != start {
		ret = append(ret[1:], ret[0])
	}

	return ret
}

func main() {
	cycle := sortByValues(findCycle(25, 5))

	sorted := slices.Clone(cycle)
	slices.Sort(sorted)
	map1 := map[int]rune{}
	map2 := map[int]rune{}

	for i, r := range runes {
		map1[sorted[i]] = r
		map2[cycle[i]] = r
	}

	keyword1, keyword2 := "", ""
	// cycle order
	for _, n := range cycle {
		keyword1 += string(map1[n])
	}
	// numeric order
	for _, n := range sorted {
		keyword2 += string(map2[n])
	}

	fs := cipher.NewTwoSquare(keyword1, keyword2)

	pt, err := fs.Decode(ciphertext)
	util.Must(err)
	fmt.Println(pt)

	cycle2 := sortByValues(findCycle(22, 4))
	latMap := map[int]bool{3: true, 7: true, 8: true, 10: true, 14: true, 16: true, 17: true, 21: true}
	longMap := map[int]bool{3: true, 5: true, 6: true, 12: true, 14: true, 19: true, 20: true, 22: true}

	latSum, longSum := 0, 0
	for i, n := range cycle2 {
		if latMap[i+1] {
			latSum += n
		}
		if longMap[i+1] {
			longSum += n
		}
	}

	fmt.Println(util.ToCoord(slices.Concat([]int{3, 7}, util.Digits(latSum), []int{1, 2, 1}, util.Digits(longSum))))
}
