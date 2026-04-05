package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	smallGrid = `.|.F.F--7.L.
F--7||.L|.F7
S7LL-J-.|-||
||.J..7.L-J|
|-F--7|J.-.|
|.|J.L--7.-|
L-J.-.|.L--J`

	bigGrid = `--F---7.-..F--7|L.......F--------77......F--7.JJ.J........||..F-7..F------7.....|J7-..JFL........|..
J.|J..|L.F.|.||.7F-7...|L---7.J..L77LF-7.|-FJ.L.F---7.......F.|.|.FL--7..FJ..F--7|....LF..L..F..L||.
.FJ.F-J....L7.L7||J|FF--7L..|F-7.||F.|.L-J.|....|.|FJ..F---77.|FJ..7.J|.FJJ.FJ..|..F--7L...J.F7..7..
FJ..|.F-7...|.FJ.|FJ.L7.L-7.||.L7L|.FJ..F7L|..F-J..|J..|.|.L7.||-.7.F-J.|...|-.L|F.|..|.|..|FJL7-...
|J-FJ.|.L---J.|F-JL7..L-7.L-JL7JL-J.|J..|L-J.-|...7L---J|...L-JL----J.7.L---J..FJF-J7.|7-.LFJ.FJ...7
|..|7FJ.FF7..-LJ-..L-7..|..F-7|JF---J..||.F---J....F7....F...-...L.L7........|FJFJJ..J|....|.FJ.....
L7.|.L7|-F---7FF.F-77L7.L7.|FJ|.L-77F..FJ.|.F.....FJL7|.L.F-----7........J..F.|.|.L..FJ...FJ.|.-L...
.|.L77|.FJ...|...|.||.||.L7||.|-..|-..||..|.F..-FL|.J|..F-J...7.|..F......F---J.|-.F-J7..FJJ.|..-|.L
.|-.|.L7|7..FJ..J|.L7.|.F.||L-J.-FJ.JF-JF-J|......|.LL--J.-..F.JL7|.JL.F--J..L.FJ|L|...|FJ..FJ..F...
.|..|..|||LFJ-..FJ.||.L7.J||F-77.|...|..|...F-7.|.L7.-|L7F-7.F-7JL77.F-J....F--J...L----J...|.......
FJ..|..LJ..|...FJ...|..L7.|LJLL--J...|.7|...|.L-7J-|.L.F-J.|.|.|..L--J.....FJ...L...JF-7L.L.L7.|....
|..FJ.F---7L7..L77|.L7..L-JJF--------J.-L7|-|...|..L---J...|.|.|.J|....J|.7|L.......FJ-L--7.FL7F-7..
|..|..L-7.||L7.-|....L7...7.L--------7...|..|...L-----7..L.L-J.L-----------J..J-|..FJ..-F.|...LJ.L7.
S..L-7.||.|..|..L---7.L--7.F--7...7.|L---J..|....-F7-.L7.......LFF--7...7..........|...7.LL-7.7.L.|F
|L...L7.|LL7.L-----7|F.F-J.|J.|--F---7.|.|.L|..-|FJ|...L-7...F.F-J.-L-7L|.......F--J.F--7...L-----JJ
L7....|JL-7L7.7F-7.LJ|.|.F-JF-J|.|||.L7-|F..L7.L.|-L7.-F.L-----J.J....|.F......FJJ7..L7.|....|..J|..
.|.F-7|...|.|.L|.L-7...L-J..|.F-FJLF7.L7-...L|.-.|J.|..F......L-.F-7..L-7L...F-J......|.L-----7..L..
.|.|.LJ...|.|.-L-7|||F--7L.7|..FJ7.||.JL--7..L-7||.JL----7...F---J.|....L----J......F-J.....J.L-7..L
|L7|....-FJ.L7.J7|.|-|..L7.|L--JF--JL-7...|..L-|7|FJ..J..|..FJ.J...L7.........JF----J.7.LF---7.JL--7
..|L7..|FJ7..L---J.||L7..L7..F-7|F--7.|...|..-.L-J-7J..F-J..|.J..-..|.-...F----J..F..-..L|..F|F....|
F.|JL---J..JF----7J|..|...|.J|.|||.FJ.|-F-JF----7......|L..F|JF-7|..L--7.7|..F-7.JL..F---JL..L-----J
..L-7.L.J.F-JF-7.L-J..|...|.FJ.||L7|.J|.L--JF---J..F7J.L----J.|.|.--...|.||..|||..F--J..............
.J.|L-7.L.L7.L7|....F-JLF-J.|..||.||..L-7...L-7...F--------7..|.L7..-F-J.7|..|F|..|....7..F-7F7.J..L
-7.|.7L--7F|F.|L----JL..|...|.FJ|.|L--7.L7J..L|F--JJF--7L..|.7|..L7|||.7-.|.L|.|..|.....F-J7L---7.7.
F--7J..-.|.|..|J.-..F---J-.||.|.L-J-..|..L--7.LJ.F--J.LL---J..|...|F.L----JJ.|.|L.L-7L..|..7.F..|...
||.L-7.JF|.L7.L7J-F-J....L|.|JL-7.F...L-7.F.||F--J..F---7....FJ..-||..F......|.|-J..L---JL.F|F--J.7.
|...|L7.L|..|..|F.L----7.F--J.--L-7..-F.L-7.|.|..|LFJL..|..L.|....|...-.F..F-J|L-7.J.....F---J..J|JJ
L7.J.F|.FJ..||F|.F--7..L-J..L.....L---7...|.|.L7..FJ7..FJ7.-FJ....L7..F....|.L|-.L-------JL...F.....
7L7...L-J.-FJ..L-J..L------7..F---7...|.7FJ.|L.|.-|.7LJ|.L..|..-|L.L-7....L|...F---7....F----7..L...
.JL--7.F|L.|7.F.J..7.......L--J...|.|.|..|.FJ.L|..|..|-|.F--J....7..|L-7...L---J...L----J....|.-...|
..|L.||..L.L----7..F-----7..F...7.L7..|.7|.|F..|7.||..FJF|......F-7....L-7..L7...7J.....-....L-7F.7|
.F-7.L--7JF---7.|..|.-...L7.F---7..|F.|..|.L---J.FJ...|.F|....F-J.|.-...FJJ...7.L.J....F---7.L.|....
.|.|.F..L-J.F-|.|..L--7...|.|..F|..|.FJ..L--7...FJ...||..L----J|.-|L..F-J...........7F-JJ..L-7.L7...
FJ.L---7..F---J7|L..L.|..FJ.L7..L7J|J|.F-...|.|.|.....|..7.......J|F..|...F..F-7...F-J.-7..7.|..|...
|..|.|JL--J....FJFF.F-J..|...|...L-J.L-7.|F-J7.FJ.|...L---7......FJJ..|.J..F-J7|.|.|.F-7..JF-J..|.7.
|F.F--7...F---7L----J....|..L|L7F---77.|.FJ....|.-.-...7.FJF-----J|..FJ...F|...|...L-JL|.J.||F--J...
|.||..L---J...L----7...F-J..-L--J.L.|..L-J....FJ.F---7.7.|J|...7J7..FJJ..F-J.F-J.J.L-..|...|.|.F--7.
L7.|F.F---7.F|.....L7|.|.F------7.J7|....J...7|7-|J..L-7.|JL--7.|7..|..-.|...|..-.F.F|FJ.-.|.|-|..L7
7L-J|FJ7..L----7.|..|L.L-J......L--7|...F--7J.|..|.F7..L-J.77.|.....|....|...L7......FJ.FLFJ.L7L7-.|
F--7L|F........L--7.L7......F----7.||.F-J.J|..|L.L-7...|L...F.L-7...|.J.FJ7...|..JF--J...F|...|.L-7|
|F-|F|....|F---7..|.-L7...F-J....|LLJ.|....|L.|.|J.L---7.......7|.-7|...|....JL7..|.LF-7F.L-7.L-7.||
|..|||.F---J.L7|.FJ..FL---J.LF---J....L7.7.|..|.||..7F.|........|.FFJ.--|.L....|..|..|.L7...L7..|.||
|FFJ.|FJ....F--J.|.F-----7.J.|......F.F|..L|..L7.F--7.FL-7FJ.7F-J..|....L7F.-..|.||..L7.L77..|..||||
|7|.-||.|F--JJ..-L-J....-|...L-7...F---J..-L7..|.|J.L-7..|L..7|.FF-J.77..|..JJ.|..|...L7.L77F|..L-J|
|.|..LJL.|L|...F------7|.L-7..-L---J..F|.-..|FLL-J..-.|.7L----J|.|....L.LL7..F-JF.|....L7.L-7|.|LF7|
|LL-7-.F-J-..F-J.-..|.|F..-|F----7.F---7....L-7...L...|...F....J-|.7..F---J77|...7|..-..L7..||...F-J
|J..L--J...F-J.7F--7..L-7..||...LL-J...L7F----J......7L-7..F--7.|L7...|F|.JF.L7..7L7.-|L.L7.LJ.F-J.J
L-7.F......L----J..|....|J.LJ.F-----7.L.LJ...J--F.JLJ.L.L--J..|...L7-.L-7.....L7-.-|.F-7..L7...L--7.
..L--7.F-----7.F-7.|....L-----JJ...-L-7.-J.|F---7JLF-----77...|....L7...|..J.7.|...L-J.L7..|..-.-FJL
J.F..L-J..|.|L-J.L-J|.F..JL...........L-----J...L--J|.|..L----J..|-.L---JJ..F.FL--------J..L-----J.7`
)

type Point struct {
	Row, Col int
}

func (p Point) Neighbors() []Point {
	return []Point{{p.Row - 1, p.Col}, {p.Row + 1, p.Col}, {p.Row, p.Col - 1}, {p.Row, p.Col + 1}}
}

var (
	gridStr          = bigGrid
	grid             [][]rune
	numRows, numCols int
	onPath           = map[Point]bool{}
	startRow         int

	extMap = map[rune]rune{
		'|': '\u2503',
		'S': '\u2503',
		'-': '\u2501',
		'F': '\u250F',
		'7': '\u2513',
		'L': '\u2517',
		'J': '\u251B',
	}
)

func init() {
	for i, line := range strings.Split(gridStr, "\n") {
		var row []rune
		for _, r := range line {
			row = append(row, r)
			if r == 'S' {
				startRow = i
			}
		}
		grid = append(grid, row)
	}

	numRows = len(grid)
	numCols = len(grid[0])
}

func findPath() {
	currRow, currCol := startRow, 0
	dir := 'S'
	for {
		onPath[Point{currRow, currCol}] = true
		switch dir {
		case 'S':
			switch grid[currRow][currCol] {
			case '|', 'S':
				currRow++
			case 'J':
				currCol--
				dir = 'W'
			case 'L':
				currCol++
				dir = 'E'
			default:
				panic(fmt.Errorf("S: (%d, %d) = %c", currRow, currCol, grid[currRow][currCol]))
			}
		case 'E':
			switch grid[currRow][currCol] {
			case '-':
				currCol++
			case 'J':
				currRow--
				dir = 'N'
			case '7':
				currRow++
				dir = 'S'
			default:
				panic(fmt.Errorf("E: (%d, %d) = %c", currRow, currCol, grid[currRow][currCol]))
			}
		case 'N':
			switch grid[currRow][currCol] {
			case '|':
				currRow--
			case '7':
				currCol--
				dir = 'W'
			case 'F':
				currCol++
				dir = 'E'
			default:
				panic(errors.New("N"))
			}
		case 'W':
			switch grid[currRow][currCol] {
			case '-':
				currCol--
			case 'F':
				currRow++
				dir = 'S'
			case 'L':
				currRow--
				dir = 'N'
			default:
				panic(errors.New("W"))
			}
		default:
			panic(errors.New("?"))
		}
		if currRow == startRow && currCol == 0 {
			break
		}
	}
}

func floodFill(p Point, m map[Point]bool) {
	if onPath[p] {
		return
	}

	queue := []Point{p}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if m[curr] {
			continue
		}
		m[curr] = true

		for _, next := range curr.Neighbors() {
			if m[next] {
				continue
			}
			// TODO: Move into Neighbors
			if next.Row < 0 || next.Row >= numRows || next.Col < 0 || next.Col >= numCols {
				continue
			}
			if onPath[next] {
				continue
			}
			queue = append(queue, next)
		}
	}
}

func printHeader() {
	fmt.Print("   ")
	for col := range numCols {
		fmt.Printf("%0d", col/10)
	}
	fmt.Printf("\n   ")
	for col := range numCols {
		fmt.Print(col % 10)
	}
	fmt.Println()
}

func main() {
	findPath()

	outside := map[Point]bool{}

	for col := range numCols {
		floodFill(Point{0, col}, outside)
		floodFill(Point{numRows - 1, col}, outside)
	}
	for row := range numRows {
		floodFill(Point{row, 0}, outside)
		floodFill(Point{row, numCols - 1}, outside)
	}

	for _, p := range []Point{
		{2, 25}, {5, 4}, {7, 5}, {7, 23}, {7, 31}, {8, 58}, {9, 9}, {9, 53},
		{10, 40}, {11, 18}, {11, 38}, {11, 66}, {12, 11}, {14, 28}, {14, 48}, {15, 14}, {16, 20}, {17, 8}, {17, 33}, {17, 38}, {18, 40},
		{23, 51}, {23, 56}, {25, 25}, {26, 37}, {27, 78}, {29, 13},
		{30, 44}, {32, 12}, {36, 64}, {37, 12}, {37, 34}, {37, 94}, {39, 88},
		{40, 4}, {41, 13}, {41, 27}, {41, 48}, {42, 97}, {43, 8}, {44, 40}, {45, 5}, {46, 15},
	} {
		floodFill(p, outside)
	}
	outCount, inCount, pipeCount := 0, 0, 0

	for row := range numRows {
		if row == 0 {
			printHeader()
		}
		fmt.Printf("%02d ", row)
		for col := range numCols {
			ch := grid[row][col]
			if ch == '.' {
				fmt.Print(".")
				continue
			}
			if onPath[Point{row, col}] {
				fmt.Print(color.HiBlueString("%c", extMap[ch]))
				pipeCount++
				continue
			}
			if outside[Point{row, col}] {
				fmt.Print(color.HiGreenString("%c", '*'))
				outCount++
			} else {
				fmt.Print(color.HiRedString("%c", '#'))
				inCount++
			}
		}
		fmt.Printf(" %02d\n", row)
	}
	printHeader()

	fmt.Printf("\npipe %d  in %d  out %d\n", pipeCount, inCount, outCount)
	fmt.Println(4175 + 7*len(onPath))
	fmt.Println(200 + 3*outCount + 8*inCount)
}
