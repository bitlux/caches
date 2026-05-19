package cipher

import (
	"errors"
	"strings"
)

// The TwoSquare cipher is based on https://www.cryptogram.org/downloads/aca.info/ciphers/TwoSquare.pdf.
type TwoSquare struct {
	one, two [5][5]byte
}

func NewTwoSquare(key1, key2 string) *TwoSquare {
	return &TwoSquare{
		one: KeywordToMatrix(key1, true),
		two: KeywordToMatrix(key2, true),
	}
}

func (t *TwoSquare) Encode(s string) (string, error) {
	s = sanitizeText(s)
	if len(s)%2 != 0 {
		return "", errors.New("plaintext length must be even")
	}

	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		west := keyedLetterToPoint(t.one, s[i])
		east := keyedLetterToPoint(t.two, s[i+1])
		if west.Row == east.Row {
			out.WriteByte(s[i+1])
			out.WriteByte(s[i])
		} else {
			out.WriteByte(t.two[west.Row][east.Col])
			out.WriteByte(t.one[east.Row][west.Col])
		}
	}
	return out.String(), nil
}

func (t *TwoSquare) Decode(s string) (string, error) {
	s = sanitizeText(s)
	if len(s)%2 != 0 {
		return "", errors.New("ciphertext length must be even")
	}

	var out strings.Builder
	for i := 0; i < len(s); i += 2 {
		west := keyedLetterToPoint(t.one, s[i+1])
		east := keyedLetterToPoint(t.two, s[i])
		if west.Row == east.Row {
			out.WriteByte(s[i+1])
			out.WriteByte(s[i])
		} else {
			out.WriteByte(t.one[east.Row][west.Col])
			out.WriteByte(t.two[west.Row][east.Col])
		}
	}
	return out.String(), nil
}
