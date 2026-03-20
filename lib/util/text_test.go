package util

import (
	"maps"
	"testing"
)

func TestRuneCount(t *testing.T) {
	s := "bitloox"
	m := RuneCount(s)
	want := map[rune]int{'b': 1, 'i': 1, 't': 1, 'l': 1, 'o': 2, 'x': 1}
	if !maps.Equal(m, want) {
		t.Errorf("RunCount(%q): got %v, want %v", s, m, want)
	}
}
