package stable

import "errors"

const (
	seperatorStartPlaceHolder      string = "S"
	seperatorMiddlePlaceHolder     string = "M"
	seperatorEndPlaceHolder        string = "E"
	seperatorVerticalPlaceHolder   string = "V"
	seperatorHorizontalPlaceHolder string = "H"

	borderStyleIndexTopLeft    int = 0
	borderStyleIndexTopCenter  int = 1
	borderStyleIndexTopRight   int = 2
	borderStyleIndexMidLeft    int = 3
	borderStyleIndexMidCenter  int = 4
	borderStyleIndexMidRight   int = 5
	borderStyleIndexBotLeft    int = 6
	borderStyleIndexBotCenter  int = 7
	borderStyleIndexBotRight   int = 8
	borderStyleIndexVertical   int = 9
	borderStyleIndexHorizontal int = 10
)

// borderStyleName border style name type
type borderStyleName string

const (
	// BorderStyleDoubleLine double stripped border style
	BorderStyleDoubleLine borderStyleName = "style-double-line"

	// BorderStyleSingleLine single stripped border style
	BorderStyleSingleLine borderStyleName = "style-single-line"

	// BorderStylePrintableLine pritable border style
	BorderStylePrintableLine borderStyleName = "style-printable-line"
)

var (

	// doubleLine predefined border styl
	doubleLine *BorderStyle = &BorderStyle{
		name: BorderStyleDoubleLine,
		runes: [25]rune{
			'╔', '╦', '╗',
			'╠', '╬', '╣',
			'╚', '╩', '╝',
			'═',
			'║',
		},
	}

	// singleLine predefined border styl
	singleLine *BorderStyle = &BorderStyle{
		name: BorderStyleSingleLine,
		runes: [25]rune{
			'┌', '┬', '┐',
			'├', '┼', '┤',
			'└', '┴', '┘',
			'─',
			'│',
		},
	}
	// printableLine predefined border styl
	printableLine *BorderStyle = &BorderStyle{
		name: BorderStyleSingleLine,
		runes: [25]rune{
			'+', '-', '+',
			'|', '+', '|',
			'+', '+', '+',
			'-',
			'|',
		},
	}
)

// BorderStyle border style main strcut
type BorderStyle struct {
	name  borderStyleName
	runes [25]rune
}

// get get char value
func (bs *BorderStyle) get(index int) string {
	return string(bs.runes[index])
}

// getStyle get predefined border style
func getStyle(name borderStyleName) (*BorderStyle, error) {
	switch name {
	case BorderStyleDoubleLine:
		return doubleLine, nil
	case BorderStyleSingleLine:
		return singleLine, nil
	case BorderStylePrintableLine:
		return printableLine, nil
	default:
		return nil, errors.New("border style not found")
	}
}
