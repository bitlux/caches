package util

import "testing"

func TestIsUnique(t *testing.T) {
	for _, tc := range []struct {
		slice []int
		want  bool
	}{
		{
			[]int{1, 2, 3},
			true,
		},
	} {
		if got := IsUnique(tc.slice...); got != tc.want {
			t.Errorf("IsUnique(%v) = %t, want %t", tc.slice, got, tc.want)
		}
	}
}
