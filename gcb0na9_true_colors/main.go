package main

import (
	"fmt"
	"strings"

	"github.com/bitlux/caches/util"
)

var matrix [][]rune

func init() {
	rows := strings.Split(`🟢🔴🔵🔴🔵🔴🔵🔵🔵🔵🔴🔵🟢🟢🔴
🔵🟢🔵🔵🔵🟢🔵🔵🔵🔵🔵🔵🔵🔵🔵
🔴🔴🟢🔴🔵🔴🔵🔵🟢🔵🔴🟢🔴🔴🔴
🔵🔴🔵🟢🔵🔴🔵🔵🔵🔵🔴🔵🔵🔵🔴
🔴🔴🔴🔴🟢🔴🔴🔴🔴🔴🔴🔴🔴🔴🔴
🔵🟢🔵🔵🔵🟢🔵🔵🔵🔵🔵🔵🔵🔵🔵
🔴🔴🔴🔴🔵🔴🟢🟢🔴🟢🔴🔴🔴🔴🔴
🔴🔴🔴🔴🔵🔴🟢🟢🔴🟢🔴🔴🔴🔴🔴
🔴🔴🟢🔴🔵🔴🔵🔵🟢🔵🔴🟢🔴🔴🔴
🔴🔴🔴🔴🔵🔴🟢🟢🔴🟢🔴🔴🔴🔴🔴
🔵🔴🔵🔵🔵🔴🔵🔵🔵🔵🟢🔵🔵🔵🟢
🔴🔴🟢🔴🔵🔴🔵🔵🟢🔵🔴🟢🔴🔴🔴
🟢🔴🔵🔴🔵🔴🔵🔵🔵🔵🔴🔵🟢🟢🔴
🟢🔴🔵🔴🔵🔴🔵🔵🔵🔵🔴🔵🟢🟢🔴
🔵🔴🔵🔵🔵🔴🔵🔵🔵🔵🟢🔵🔵🔵🟢`, "\n")
	for _, row := range rows {
		var newRow []rune
		for _, r := range row {
			switch r {
			case '🟢':
				newRow = append(newRow, '=')
			case '🔴':
				newRow = append(newRow, '<')
			case '🔵':
				newRow = append(newRow, '>')
			}
		}
		matrix = append(matrix, newRow)
	}
}

func testRow(index int, candidate []int) bool {
	row := matrix[index]
	cand := candidate[index]
	for j, candJ := range candidate {
		switch row[j] {
		case '=':
			if cand != candJ {
				return false
			}
		case '<':
			if cand >= candJ {
				return false
			}
		case '>':
			if cand <= candJ {
				return false
			}
		}
	}
	return true
}

func main() {
	entries := util.Explode(
		[]int{3}, []int{7}, []int{2}, []int{2, 3, 4}, util.D, []int{7}, util.D,
		[]int{1}, []int{2}, []int{1}, []int{4, 5}, util.D, []int{3}, []int{3}, util.D)
OUTER:
	for entry := range entries {
		for i := range 15 {
			if !testRow(i, entry) {
				continue OUTER
			}
		}
		fmt.Println(util.ToCoord(entry))
	}
}
