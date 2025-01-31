package main

import (
	"fmt"
	"os"
)

const input = `^v>>>vvvv<>><>>>>>><>^v>vvv<<>>vv<><>^>^^^>>>^>vv<><<>^<v^>^v<>>>^^>v^<^^^<^^><^>>>v^>^v<>v^<^v>>^<v>^<^>>><vv>v>>>v<><><v>v>v<<^v<<><>^<<<vvvv<^><v^<<^v><>>><^^<><<^v<v^<>><>^vv<><^<^^<vv<<^v<>>^<>>>>>>v^>>>>><^v><>v^<<^v^>^^^<^<v<>^v<<v>^^^vv>>^>>^^<v^^<v<^v<>>><<>>>>>>>>>>>>>^^<^^<^vv>>><^^v<^v>><v>>^vv>v^>>v>^><>vv<>>^v<<^v>>v^^vvv^^^^vv<<<<<<<<<<>vvvv^<^^^v<<^v^vv^^<<<<<<<<<v<v<<>^^v^v<<^^<vv>><><^v<<>^^<<<>>>>>^<v>vv>^v>>v<`

type point struct {
	x, y int
}

func main() {
	count := 0
	curr := point{0, 0}
	visited := map[point]bool{curr: true}

	for _, c := range input {
		switch c {
		case '^':
			curr.y++
		case 'v':
			curr.y--
		case '<':
			curr.x--
		case '>':
			curr.x++
		default:
			fmt.Println("illegal input: ", c)
			os.Exit(1)
		}
		count++
		visited[curr] = true
	}

	fmt.Printf("curr = %v  |visited| = %d  count = %d\n", curr, len(visited), count)
	fmt.Println(len(input))
}
