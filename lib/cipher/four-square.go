package cipher

import (
	"errors"
	"slices"
	"strings"
)

// ----------------------------------------------------------------------------
// Four-square

type point struct {
	Row, Col int
}

func letterToPoint(b byte) point {
	p := point{-1, -1}
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

var unkeyed = [5][5]rune{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'I', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

func sanitizeText(s string) string {
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

func keywordToMatrix(s string, horizontal bool) [5][5]rune {
	s = sanitizeText(s)
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

// 1 2
// 4 3
type FourSquare struct {
	// two and four are for printing and encoding
	Two, Four [5][5]rune

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
			b.WriteRune(f.Two[row][col])
		}
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for row := range 5 {
		for col := range 5 {
			b.WriteRune(f.Four[row][col])
		}
		b.WriteRune(' ')
		for col := range 5 {
			b.WriteRune(unkeyed[row][col])
		}
		b.WriteRune('\n')
	}
	return b.String()
}

// NewFourSquare creates a new four square cipher instance.
// GC49KY3 uses horizontal = false.
// CacheSleuth uses horizontal = true.
func NewFourSquare(key1, key2 string, horizontal bool) *FourSquare {
	c := &FourSquare{
		Two:     keywordToMatrix(key1, horizontal),
		Four:    keywordToMatrix(key2, horizontal),
		twoMap:  map[rune]point{},
		fourMap: map[rune]point{},
	}

	for row := range 5 {
		for col := range 5 {
			c.twoMap[c.Two[row][col]] = point{Row: row, Col: col}
			c.fourMap[c.Four[row][col]] = point{Row: row, Col: col}
		}
	}
	return c
}

func (f *FourSquare) Encode(s string) (string, error) {
	s = sanitizeText(s)
	if len(s)%2 != 0 {
		return "", errors.New("plaintext length must be even")
	}

	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		nw := letterToPoint(s[i])
		se := letterToPoint(s[i+1])
		out.WriteRune(f.Two[nw.Row][se.Col])
		out.WriteRune(f.Four[se.Row][nw.Col])
	}
	return out.String(), nil
}

func (f *FourSquare) Decode(s string) (string, error) {
	s = sanitizeText(s)
	if len(s)%2 != 0 {
		return "", errors.New("plaintext length must be even")
	}

	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		ne := f.twoMap[rune(s[i])]
		sw := f.fourMap[rune(s[i+1])]
		out.WriteRune(unkeyed[ne.Row][sw.Col])
		out.WriteRune(unkeyed[sw.Row][ne.Col])
	}
	return out.String(), nil
}
