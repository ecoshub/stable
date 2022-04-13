package style

import (
	"errors"
	"fmt"
	"math"
)

// DoPadding padding to a value with space size and padding type
func DoPadding(value string, colunSize int, paddingType string) (string, error) {
	valueLength := len([]rune(value))
	padding := colunSize - valueLength

	switch paddingType {
	case AlignementLeft:
		s := value + nSpace(padding)
		return s, nil
	case AlignementCenter:
		left := int(math.Ceil(float64(padding) / 2.0))
		s := nSpace(left)
		s += value
		s += nSpace(padding - left)
		return s, nil
	case AlignementRight:
		if valueLength > colunSize {
			return value, errors.New("value length is overflowing")
		}
		colunSize = padding
		s := nSpace(colunSize)
		s += value
		return s, nil
	default:
		return value, fmt.Errorf("unknown padding '%s'. possible paddings left|center|right", paddingType)
	}
}

func nSpace(n int) string {
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = ' '
	}
	return string(s)
}
