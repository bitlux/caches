package util

import "testing"

func TestToCoord(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want string
	}{
		{
			3724123,
			"N 37 24.123",
		},
		{
			12204789,
			"W 122 04.789",
		},
	} {
		got := ToCoord(Digits(tc.n))
		if got != tc.want {
			t.Errorf("ToCoord(%d) = %s, want %s", tc.n, got, tc.want)
		}
	}
}
