package util

import (
	"maps"
	"testing"
)

func TestIsUnique(t *testing.T) {
	for _, tc := range []struct {
		slice []int
		want  bool
	}{
		{
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1},
			true,
		},
		{
			[]int{1, 2, 3, 9, 8, 78},
			true,
		},
		{
			[]int{1, 1},
			false,
		},
		{
			[]int{3, 1, 2, 3},
			false,
		},
	} {
		if got := IsUnique(tc.slice...); got != tc.want {
			t.Errorf("IsUnique(%v) = %t, want %t", tc.slice, got, tc.want)
		}
	}
}

func TestAsSet(t *testing.T) {
	slice := []int{1, 2, 3}
	want := map[int]bool{
		1: true,
		2: true,
		3: true,
	}
	if got := AsSet(slice); !maps.Equal(got, want) {
		t.Errorf("AsSet(%v) = %v, want %v", slice, got, want)
	}
}
