package util

import (
	"slices"
	"strings"

	"golang.org/x/exp/constraints"
)

// ----------------------------------------------------------------------------
// Substitution ciphers

// A1Encode encodes a rune in the range [A-Za-z] using the A=1, ..., Z=26 substitution cipher.
func A1Encode[T constraints.Integer](n T) int {
	if n >= 'A' && n <= 'Z' {
		return int(n - 'A' + 1)
	}
	return int(n - 'a' + 1)
}

// A1Decode decodes a number in the range [1-26] using the A=1, ..., Z=26 substitution cipher.
func A1Decode(n int) rune {
	return rune(n + 'A' - 1)
}

// CBF encodes a string into a slice of integers. CBF encoding is similar to A1Encode, but done
// mod 10.
func CBF(s string) []int {
	var ret []int
	for _, c := range s {
		ret = append(ret, A1Encode(c)%10)
	}
	return ret
}

// ROT rotates w by n letter. ROT(13, "terra") = "green". Currently only handles lowercase letters.
func ROT(n int, w string) string {
	ret := ""
	for _, l := range w {
		ret += string(rune(int(l)-'a'+n)%26 + 'a')
	}
	return ret
}

// ----------------------------------------------------------------------------
// Four-square

type point struct {
	row, col int
}

func letterToPoint(b byte) point {
	p := point{-1, -1}
	switch b {
	case 'A', 'B', 'C', 'D', 'E':
		p.row = 0
	case 'F', 'G', 'H', 'I', 'K':
		p.row = 1
	case 'L', 'M', 'N', 'O', 'P':
		p.row = 2
	case 'Q', 'R', 'S', 'T', 'U':
		p.row = 3
	case 'V', 'W', 'X', 'Y', 'Z':
		p.row = 4
	}
	switch b {
	case 'A', 'F', 'L', 'Q', 'V':
		p.col = 0
	case 'B', 'G', 'M', 'R', 'W':
		p.col = 1
	case 'C', 'H', 'N', 'S', 'X':
		p.col = 2
	case 'D', 'I', 'O', 'T', 'Y':
		p.col = 3
	case 'E', 'K', 'P', 'U', 'Z':
		p.col = 4
	}
	return p
}

var unkeyed = [5][5]rune{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'I', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

// 1 2
// 4 3
type FourSquare struct {
	// two and four are for printing and encoding
	two, four [5][5]rune

	// twoMap and fourMap are for decoding
	twoMap, fourMap map[rune]point
}

func (f *FourSquare) String() string {
	var b strings.Builder
	for row := range 5 {
		for col := range 5 {
			b.WriteRune(unkeyed[row][col])
		}
		b.WriteRune(' ')
		for col := range 5 {
			b.WriteRune(f.two[row][col])
		}
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for row := range 5 {
		for col := range 5 {
			b.WriteRune(f.four[row][col])
		}
		b.WriteRune(' ')
		for col := range 5 {
			b.WriteRune(unkeyed[row][col])
		}
		b.WriteRune('\n')
	}
	return b.String()
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

func keywordToMatrix(s string) [5][5]rune {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "J", "I")
	s = strings.ReplaceAll(s, "'", "")

	alphabet := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	for i, c := range removeDuplicates(s) {
		index := slices.Index(alphabet, c)
		alphabet = slices.Concat(alphabet[:i], []rune{c}, alphabet[i:index], alphabet[index+1:])
	}
	var m [5][5]rune
	for row := range 5 {
		for col := range 5 {
			m[col][row] = alphabet[5*row+col]
		}
	}
	return m
}

func NewFourSquare(key1, key2 string) *FourSquare {
	c := &FourSquare{
		two:     keywordToMatrix(key1),
		four:    keywordToMatrix(key2),
		twoMap:  map[rune]point{},
		fourMap: map[rune]point{},
	}

	for row := range 5 {
		for col := range 5 {
			c.twoMap[c.two[row][col]] = point{row, col}
			c.fourMap[c.four[row][col]] = point{row, col}
		}
	}
	return c
}

func (f *FourSquare) Encode(s string) string {
	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		nw := letterToPoint(s[i])
		se := letterToPoint(s[i+1])
		out.WriteRune(f.two[nw.row][se.col])
		out.WriteRune(f.four[se.row][nw.col])
	}
	return out.String()
}

func (f *FourSquare) Decode(s string) string {
	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		ne := f.twoMap[rune(s[i])]
		sw := f.fourMap[rune(s[i+1])]
		out.WriteRune(unkeyed[ne.row][sw.col])
		out.WriteRune(unkeyed[sw.row][ne.col])
	}
	return out.String()
}
