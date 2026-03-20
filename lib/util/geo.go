package util

import "fmt"

func ToCoord(digits []int) string {
	var asAny []any
	for _, d := range digits {
		asAny = append(asAny, d)
	}
	switch len(digits) {
	case 7:
		return fmt.Sprintf("N %d%d %d%d.%d%d%d", asAny...)
	case 8:
		return fmt.Sprintf("W %d%d%d %d%d.%d%d%d", asAny...)
	case 15:
		return fmt.Sprintf("%s %s", ToCoord(digits[:7]), ToCoord(digits[7:]))
	}
	return ""
}
