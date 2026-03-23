package cipher

import (
	"errors"
	"regexp"
	"slices"
	"strings"
)

var whitespaceRE = regexp.MustCompile(`\s+`)

// TODO: type polybius [5][5]byte

var unkeyed = [5][5]byte{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'I', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

type point struct {
	Row, Col int
}

func unkeyedLetterToPoint(b byte) point {
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

func keyedLetterToPoint(matrix [5][5]byte, b byte) point {
	for row := range 5 {
		for col := range 5 {
			if matrix[row][col] == byte(b) {
				return point{Row: row, Col: col}
			}
		}
	}
	return point{Row: -1, Col: -1}
}

func sanitizeText(s string) string {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "J", "I")
	s = strings.ReplaceAll(s, "'", "")
	s = whitespaceRE.ReplaceAllString(s, "")
	return s
}

func removeDuplicates(s string) []byte {
	m := map[byte]bool{}
	var ret []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !m[c] {
			ret = append(ret, c)
		}
		m[c] = true
	}
	return ret
}

func KeywordToMatrix(s string, horizontal bool) [5][5]byte {
	s = sanitizeText(s)
	alphabet := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	for i, c := range removeDuplicates(s) {
		index := slices.Index(alphabet, byte(c))
		alphabet = slices.Concat(alphabet[:i], []byte{c}, alphabet[i:index], alphabet[index+1:])
	}
	var m [5][5]byte
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
	two, four [5][5]byte

	// twoMap and fourMap are for decoding
	twoMap, fourMap map[byte]point
}

func (f *FourSquare) String() string {
	var b strings.Builder
	for row := range 5 {
		for col := range 5 {
			b.WriteByte(unkeyed[row][col])
		}
		b.WriteByte(' ')
		for col := range 5 {
			b.WriteByte(f.two[row][col])
		}
		b.WriteByte('\n')
	}
	b.WriteByte('\n')
	for row := range 5 {
		for col := range 5 {
			b.WriteByte(f.four[row][col])
		}
		b.WriteByte(' ')
		for col := range 5 {
			b.WriteByte(unkeyed[row][col])
		}
		b.WriteByte('\n')
	}
	return b.String()
}

// NewFourSquare creates a new four square cipher instance.
// GC49KY3 uses horizontal = false.
// CacheSleuth uses horizontal = true.
func NewFourSquare(key1, key2 string, horizontal bool) *FourSquare {
	c := &FourSquare{
		two:     KeywordToMatrix(key1, horizontal),
		four:    KeywordToMatrix(key2, horizontal),
		twoMap:  map[byte]point{},
		fourMap: map[byte]point{},
	}

	for row := range 5 {
		for col := range 5 {
			c.twoMap[c.two[row][col]] = point{Row: row, Col: col}
			c.fourMap[c.four[row][col]] = point{Row: row, Col: col}
		}
	}
	return c
}

func newFourSquareFromMatrices(two, four [5][5]byte) *FourSquare {
	return &FourSquare{
		two:  two,
		four: four,
	}
}

func (f *FourSquare) Encode(s string) (string, error) {
	s = sanitizeText(s)
	if len(s)%2 != 0 {
		return "", errors.New("plaintext length must be even")
	}

	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		nw := unkeyedLetterToPoint(s[i])
		se := unkeyedLetterToPoint(s[i+1])
		out.WriteByte(f.two[nw.Row][se.Col])
		out.WriteByte(f.four[se.Row][nw.Col])
	}
	return out.String(), nil
}

func (f *FourSquare) Decode(s string) (string, error) {
	s = sanitizeText(s)
	if len(s)%2 != 0 {
		return "", errors.New("ciphertext length must be even")
	}

	var out strings.Builder
	var ne, sw point
	for i := 0; i < len(s); i += 2 {
		if f.twoMap != nil {
			ne = f.twoMap[s[i]]
			sw = f.fourMap[s[i+1]]
		} else {
			ne = keyedLetterToPoint(f.two, s[i])
			sw = keyedLetterToPoint(f.four, s[i+1])
		}
		out.WriteByte(unkeyed[ne.Row][sw.Col])
		out.WriteByte(unkeyed[sw.Row][ne.Col])
	}
	return out.String(), nil
}
