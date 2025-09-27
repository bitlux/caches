package util

import (
	"slices"
	"testing"
)

func TestPermutations(t *testing.T) {
	want := []string{"c", "a", "b"}
	perms := Permutations([]string{"a", "b", "c"})

	if ln := len(perms); ln != 6 {
		t.Errorf("len is %d, want 6", ln)
	}

	if !slices.ContainsFunc(perms, func(s []string) bool { return slices.Equal(s, want) }) {
		t.Errorf("%v missing from perms", want)
	}
}

func BenchmarkPermutations(b *testing.B) {
	var p [][]int
	for b.Loop() {
		p = Permutations([]int{1, 2, 3, 4, 5, 6, 7, 8})
	}
	_ = p
}
