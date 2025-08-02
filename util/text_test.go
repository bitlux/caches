package util

import (
	"slices"
	"testing"
)

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
