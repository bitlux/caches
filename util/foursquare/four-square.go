// Package foursquare contains implementation details of the Four-square cipher that are not useful
// to users of the util package but are useful to certain other repos.
package foursquare

import (
	"slices"
	"strings"
)

type Point struct {
	Row, Col int
}

func LetterToPoint(b byte) Point {
	p := Point{-1, -1}
	switch b {
	case 'A', 'B', 'C', 'D', 'E':
		p.Row = 0
	case 'F', 'G', 'H', 'I', 'K':
		p.Row = 1
	case 'L', 'M', 'N', 'O', 'P':
		p.Row = 2
	case 'Q', 'R', 'S', 'T', 'U':
		p.Row = 3
	case 'V', 'W', 'X', 'Y', 'Z':
		p.Row = 4
	}
	switch b {
	case 'A', 'F', 'L', 'Q', 'V':
		p.Col = 0
	case 'B', 'G', 'M', 'R', 'W':
		p.Col = 1
	case 'C', 'H', 'N', 'S', 'X':
		p.Col = 2
	case 'D', 'I', 'O', 'T', 'Y':
		p.Col = 3
	case 'E', 'K', 'P', 'U', 'Z':
		p.Col = 4
	}
	return p
}

var Unkeyed = [5][5]rune{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'I', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

func SanitizeText(s string) string {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "J", "I")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func removeDuplicates(s string) string {
	m := map[rune]bool{}
	var out strings.Builder
	for _, c := range s {
		if !m[c] {
			out.WriteRune(c)
		}
		m[c] = true
	}
	return out.String()
}

func KeywordToMatrix(s string, horizontal bool) [5][5]rune {
	s = SanitizeText(s)
	alphabet := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	for i, c := range removeDuplicates(s) {
		index := slices.Index(alphabet, c)
		alphabet = slices.Concat(alphabet[:i], []rune{c}, alphabet[i:index], alphabet[index+1:])
	}
	var m [5][5]rune
	for row := range 5 {
		for col := range 5 {
			if horizontal {
				m[row][col] = alphabet[5*row+col]
			} else {
				m[col][row] = alphabet[5*row+col]
			}
		}
	}
	return m
}
