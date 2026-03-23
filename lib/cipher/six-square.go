package cipher

import (
	"errors"
	"fmt"
	"strings"
)

type SixSquare struct {
	four *FourSquare
	six  [5][5]byte
}

func (s *SixSquare) String() string {
	var b strings.Builder
	b.WriteString(s.four.String())
	b.WriteByte('\n')
	for row := range 5 {
		for col := range 5 {
			b.WriteByte(unkeyed[row][col])
		}
		b.WriteByte(' ')
		for col := range 5 {
			b.WriteByte(s.six[row][col])
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func NewSixSquare(a, b, c string) *SixSquare {
	return &SixSquare{
		four: NewFourSquare(a, b, false),
		six:  KeywordToMatrix(c, false),
	}
}

func NewSixSquareFromMatrices(two, four, six [5][5]byte) *SixSquare {
	return &SixSquare{
		four: newFourSquareFromMatrices(two, four),
		six:  six,
	}
}

func (six *SixSquare) Encode(s string) (string, error) {
	s = sanitizeText(s)
	if len(s)%3 != 0 {
		return "", fmt.Errorf("plaintext length must be divisible by 3 (is %d)", len(s))
	}

	var out strings.Builder
	for i := 0; i < len(s); i += 3 {
		one := unkeyedLetterToPoint(s[i])
		three := unkeyedLetterToPoint(s[i+1])
		five := unkeyedLetterToPoint(s[i+2])
		out.WriteByte(six.four.two[one.Row][three.Col])
		out.WriteByte(six.four.four[three.Row][one.Col])
		out.WriteByte(six.four.four[three.Row][five.Col])
		out.WriteByte(six.six[five.Row][three.Col])
		out.WriteByte(' ')
	}
	return out.String(), nil
}

func (six *SixSquare) Decode(s string) (string, error) {
	s = sanitizeText(s)
	if len(s)%4 != 0 {
		return "", errors.New("ciphertext length must be divisible by 4")
	}

	var out strings.Builder
	for i := 0; i < len(s); i += 4 {
		twoPoint := keyedLetterToPoint(six.four.two, s[i])
		fourPoint1 := keyedLetterToPoint(six.four.four, s[i+1])
		fourPoint2 := keyedLetterToPoint(six.four.four, s[i+2])
		sixPoint := keyedLetterToPoint(six.six, s[i+3])

		if fourPoint1.Row != fourPoint2.Row || twoPoint.Col != sixPoint.Col {
			return "", fmt.Errorf("%v != %v", fourPoint1, fourPoint2)
		}

		out.WriteByte(unkeyed[twoPoint.Row][fourPoint1.Col])
		out.WriteByte(unkeyed[fourPoint2.Row][sixPoint.Col])
		out.WriteByte(unkeyed[sixPoint.Row][fourPoint2.Col])
		out.WriteByte(' ')
	}

	return out.String(), nil
}
