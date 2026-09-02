package util

import (
	"slices"
	"testing"
)

func TestFactor(t *testing.T) {
	for _, test := range []struct {
		n       int
		factors []int
	}{
		{3, []int{3}},
		{27, []int{3, 3, 3}},
		{4, []int{2, 2}},
		{10, []int{2, 5}},
	} {
		got := Factor(test.n)
		if !slices.Equal(got, test.factors) {
			t.Errorf("Factor(%d): got: %v, want: %v", test.n, got, test.factors)
		}
	}
}

func TestIsPrime(t *testing.T) {
	for _, test := range []struct {
		n    int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{27, false},
		{4, false},
		{10, false},
	} {
		got := IsPrime(test.n)
		if got != test.want {
			t.Errorf("IsPrime(%d): got: %t, want: %t", test.n, got, test.want)
		}
	}
}

func TestDigits(t *testing.T) {
	for _, test := range []struct {
		n      int
		digits []int
	}{
		{131, []int{1, 3, 1}},
		{27, []int{2, 7}},
		{4, []int{4}},
		{10, []int{1, 0}},
	} {
		got := Digits(test.n)
		if !slices.Equal(got, test.digits) {
			t.Errorf("Digits(%d): got: %v, want: %v", test.n, got, test.digits)
		}
	}
}

func TestFromDigits(t *testing.T) {
	for _, tc := range []int{1, 101, 1618033, 2468} {
		num := FromDigits(Digits(tc))
		if num != tc {
			t.Errorf("FromDigits(Digits(%d)) = %d, want %d", tc, num, tc)
		}
	}
}

func TestDigitalRoot(t *testing.T) {
	for _, tc := range []struct {
		input, want int
	}{
		{10, 1},
		{11, 2},
		{1, 1},
		{9, 9},
		{33, 6},
		{1234, 1},
		{9999, 9},
	} {
		got := DigitalRoot(tc.input)
		if got != tc.want {
			t.Errorf("DigitalRoot(%d) = %d, want %d", tc.input, got, tc.want)
		}
	}
}

func TestCollatzStoppingTime(t *testing.T) {
	for _, tc := range []struct {
		input, want int
	}{
		{1, 0},
		{19, 20},
		{11, 14},
	} {
		if got := CollatzStoppingTime(tc.input); got != tc.want {
			t.Errorf("CollatzStoppingTime(%d) = %d, want %d", tc.input, got, tc.want)
		}
	}
}

func TestConvertBase(t *testing.T) {
	for _, tc := range []struct {
		input, want string
		from, to    int
	}{
		{"3", "11", 10, 2},
		{"31", "13", 4, 10},
		{"101010", "42", 2, 10},
	} {
		got, err := ConvertBase(tc.input, tc.from, tc.to)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if got != tc.want {
			t.Errorf("ConvertBase(%q, %d, %d) = %q, want %q", tc.input, tc.from, tc.to, got, tc.want)
		}
	}

}
