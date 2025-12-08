package util

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/exp/constraints"
)

// A1Encode encodes a rune in the range [A-Za-z] using the A=1, ..., Z=26 substitution cipher.
func A1Encode[T constraints.Integer](n T) int {
	if n >= 'A' && n <= 'Z' {
		return int(n - 'A' + 1)
	}
	return int(n - 'a' + 1)
}

// A1Decode decodes a number in the range [1-26] using the A=1, ..., Z=26 substitution cipher.
func A1Decode(n int) rune {
	return rune(n + 'A' - 1)
}

// CBF encodes a string into a slice of integers. CBF encoding is similar to A1Encode, but done
// mod 10.
func CBF(s string) []int {
	var ret []int
	for _, c := range s {
		ret = append(ret, A1Encode(c)%10)
	}
	return ret
}

// ROT rotates w by n letter. ROT(13, "terra") = "green". Currently only handles lowercase letters.
func ROT(n int, w string) string {
	ret := ""
	for _, l := range w {
		ret += string(rune(int(l)-'a'+n)%26 + 'a')
	}
	return ret
}

// SHA256 returns the SHA-256 hash of the input as a hex-encoded string.
func SHA256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
