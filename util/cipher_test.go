package util

import (
	"slices"
	"strings"
	"testing"
)

func TestA1EncodeDecode(t *testing.T) {
	for _, tc := range []struct {
		r    rune
		want int
	}{
		{'a', 1},
		{'A', 1},
		{'z', 26},
		{'Z', 26},
	} {
		enc := A1Encode(tc.r)
		if enc != tc.want {
			t.Errorf("A1Encode(%c): got %d, want %d", tc.r, enc, tc.want)
		}
		if got := A1Decode(enc); got != []rune(strings.ToUpper(string(tc.r)))[0] {
			t.Errorf("A1Decode(%d): got %c, want %c", enc, got, tc.r)
		}
	}
}

func TestCBF(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []int
	}{
		{"abc", []int{1, 2, 3}},
		{"xyz", []int{4, 5, 6}},
	} {
		if got := CBF(tc.s); !slices.Equal(got, tc.want) {
			t.Errorf("CBF(%s) = %v, want %v", tc.s, got, tc.want)
		}
	}
}
