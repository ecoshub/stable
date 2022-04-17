package stable

import (
	"math"
)

// doPadding padding to a value with space size and padding type
func doPadding(value string, columnSize int, alignment alignment) string {
	valueLength := len([]rune(value))
	padding := columnSize - valueLength

	switch alignment {
	// nil string is representing no alignment
	case AlignmentLeft, "":
		s := value + nSpace(padding)
		return s
	case AlignmentCenter:
		left := int(math.Ceil(float64(padding) / 2.0))
		s := nSpace(left)
		s += value
		s += nSpace(padding - left)
		return s
	case AlignmentRight:
		if valueLength > columnSize {
			return "err"
		}
		columnSize = padding
		s := nSpace(columnSize)
		s += value
		return s
	}
	return "wrong-alignment"

}

func nSpace(n int) string {
	if n < 0 {
		return ""
	}
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = ' '
	}
	return string(s)
}
