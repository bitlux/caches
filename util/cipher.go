package util

import (
	"strings"

	"github.com/bitlux/caches/util/foursquare"
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

// 1 2
// 4 3
type FourSquare struct {
	// two and four are for printing and encoding
	Two, Four [5][5]rune

	// twoMap and fourMap are for decoding
	twoMap, fourMap map[rune]foursquare.Point
}

func (f *FourSquare) String() string {
	var b strings.Builder
	for row := range 5 {
		for col := range 5 {
			b.WriteRune(foursquare.Unkeyed[row][col])
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
			b.WriteRune(foursquare.Unkeyed[row][col])
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
		Two:     foursquare.KeywordToMatrix(key1, horizontal),
		Four:    foursquare.KeywordToMatrix(key2, horizontal),
		twoMap:  map[rune]foursquare.Point{},
		fourMap: map[rune]foursquare.Point{},
	}

	for row := range 5 {
		for col := range 5 {
			c.twoMap[c.Two[row][col]] = foursquare.Point{row, col}
			c.fourMap[c.Four[row][col]] = foursquare.Point{row, col}
		}
	}
	return c
}

func (f *FourSquare) Encode(s string) string {
	s = foursquare.SanitizeText(s)
	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		nw := foursquare.LetterToPoint(s[i])
		se := foursquare.LetterToPoint(s[i+1])
		out.WriteRune(f.Two[nw.Row][se.Col])
		out.WriteRune(f.Four[se.Row][nw.Col])
	}
	return out.String()
}

func (f *FourSquare) Decode(s string) string {
	s = foursquare.SanitizeText(s)
	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		ne := f.twoMap[rune(s[i])]
		sw := f.fourMap[rune(s[i+1])]
		out.WriteRune(foursquare.Unkeyed[ne.Row][sw.Col])
		out.WriteRune(foursquare.Unkeyed[sw.Row][ne.Col])
	}
	return out.String()
}
