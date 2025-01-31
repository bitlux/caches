package main

import "fmt"

const (
	n = 18457
	m = 59209
)

var dir = "sw"

func getNextStep(x, y int) (int, int) {
	if dir == "sw" {
		if y == 0 {
			dir = "ne"
			return x + 1, y
		}
		return x + 1, y - 1
	}
	if x == 0 {
		dir = "sw"
		return x, y + 1
	}
	return x - 1, y + 1
}

func main() {
	id := 10001
	x := 0
	y := 0
	for {
		// fmt.Printf("(%d, %d) = %d\n", x, y, id)

		if x == 37 && y == 0 {
			fmt.Printf("(%d, %d) = %d\n", x, y, id)
		}
		if x == 0 && y == 122 {
			fmt.Printf("(%d, %d) = %d\n", x, y, id)
			break
		}

		id = (id + n) % m
		x, y = getNextStep(x, y)
	}
}
