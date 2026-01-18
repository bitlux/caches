package util

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExplode(t *testing.T) {
	for _, tc := range []struct {
		name   string
		inputs [][]int
		want   [][]int
	}{
		{
			name:   "2x2",
			inputs: [][]int{{0, 1}, {1, 2}},
			want:   [][]int{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		},
		{
			name:   "1x2x3",
			inputs: [][]int{{3}, {1, 2}, {4, 5, 6}},
			want:   [][]int{{3, 1, 4}, {3, 1, 5}, {3, 1, 6}, {3, 2, 4}, {3, 2, 5}, {3, 2, 6}},
		},
	} {
		coll := slices.Collect(Explode2(tc.inputs...))
		diff := cmp.Diff(coll, tc.want)
		if diff != "" {
			t.Errorf("Explode(%v): %s", tc.inputs, diff)
		}
	}
}
