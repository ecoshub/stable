package style

import "errors"

const (

	// SeperatorStartPlaceHolder seperator placeholder for Start
	SeperatorStartPlaceHolder string = "S"
	// SeperatorMiddlePlaceHolder seperator placeholder for Middle
	SeperatorMiddlePlaceHolder string = "M"
	// SeperatorEndPlaceHolder seperator placeholder for End
	SeperatorEndPlaceHolder string = "E"
	// SeperatorVerticalPlaceHolder seperator placeholder for Vertical
	SeperatorVerticalPlaceHolder string = "V"
	// SeperatorHorizontalPlaceHolder seperator placeholder for Horizontal
	SeperatorHorizontalPlaceHolder string = "H"

	// BorderStyleIndexTopLeft border char index for Top Left
	BorderStyleIndexTopLeft int = 0
	// BorderStyleIndexTopCenter border char index for Top Center
	BorderStyleIndexTopCenter int = 1
	// BorderStyleIndexTopRight border char index for Top Right
	BorderStyleIndexTopRight int = 2
	// BorderStyleIndexMidLeft border char index for Mid Left
	BorderStyleIndexMidLeft int = 3
	// BorderStyleIndexMidCenter border char index for Mid Center
	BorderStyleIndexMidCenter int = 4
	// BorderStyleIndexMidRight border char index for Mid Right
	BorderStyleIndexMidRight int = 5
	// BorderStyleIndexBotLeft border char index for Bot Left
	BorderStyleIndexBotLeft int = 6
	// BorderStyleIndexBotCenter border char index for Bot Center
	BorderStyleIndexBotCenter int = 7
	// BorderStyleIndexBotRight border char index for Bot Right
	BorderStyleIndexBotRight int = 8
	// BorderStyleIndexVertical border char index for Vertical
	BorderStyleIndexVertical int = 9
	// BorderStyleIndexHorizontal border char index for Horizontal
	BorderStyleIndexHorizontal int = 10
)

// BorderStyleName border style name type
type BorderStyleName string

const (
	// BorderStyleDoubleLine double stripped border style
	BorderStyleDoubleLine BorderStyleName = "bs double line"

	// BorderStyleSingleLine single stripped border style
	BorderStyleSingleLine BorderStyleName = "bs single line"

	// BorderStylePrintableLine pritable border style
	BorderStylePrintableLine BorderStyleName = "bs printable line"
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
	name  BorderStyleName
	runes [25]rune
}

// Get get char value
func (bs *BorderStyle) Get(index int) string {
	return string(bs.runes[index])
}

// GetStyle get predefined border style
func GetStyle(name BorderStyleName) (*BorderStyle, error) {
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
