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
