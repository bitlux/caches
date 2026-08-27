package main

import (
	"fmt"

	"github.com/bitlux/caches/lib/util"
)

const puzzle = "64808341397925771754311573787726916936586178620871621541944880412872661527497280940583628987905012418629597510525677043"

func main() {
	var coords []int
	for i := 9; i < len(puzzle); {
		next := int(puzzle[i] - '0')
		fmt.Println(coords, next, i)
		if next == 0 {
			break
		}
		coords = append(coords, int(puzzle[i-1]-'0'))
		i += next
	}

	fmt.Println(util.ToCoord(coords))
}
