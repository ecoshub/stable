package stable

import "errors"

const (

	// seperator place holders
	_SPHStart      string = "S"
	_SPHMiddle     string = "M"
	_SPHEnd        string = "E"
	_SPHVertical   string = "V"
	_SPHHorizontal string = "H"

	// Border Style Indices
	_BSITopLeft    int = 0
	_BSITopCenter  int = 1
	_BSITopRight   int = 2
	_BSIMidLeft    int = 3
	_BSIMidCenter  int = 4
	_BSIMidRight   int = 5
	_BSIBotLeft    int = 6
	_BSIBotCenter  int = 7
	_BSIBotRight   int = 8
	_BSIVertical   int = 9
	_BSIHorizontal int = 10
)

const (
	// BorderStyleDoubleLine double stripped border style
	BorderStyleDoubleLine borderStyleName = 1

	// BorderStyleSingleLine single stripped border style
	BorderStyleSingleLine borderStyleName = 2

	// BorderStylePrintableLine printable border style
	BorderStylePrintableLine borderStyleName = 3
)

// borderStyleName border style name type
type borderStyleName int

// BorderStyle border style main strcut
type BorderStyle struct {
	name  borderStyleName
	chars [25]string
}

var (

	// doubleLine predefined border style
	doubleLine *BorderStyle = &BorderStyle{
		name: BorderStyleDoubleLine,
		chars: [25]string{
			"╔", "╦", "╗",
			"╠", "╬", "╣",
			"╚", "╩", "╝",
			"═",
			"║",
		},
	}

	// singleLine predefined border style
	singleLine *BorderStyle = &BorderStyle{
		name: BorderStyleSingleLine,
		chars: [25]string{
			"┌", "┬", "┐",
			"├", "┼", "┤",
			"└", "┴", "┘",
			"─",
			"│",
		},
	}
	// printableLine predefined border style
	printableLine *BorderStyle = &BorderStyle{
		name: BorderStyleSingleLine,
		chars: [25]string{
			"+", "-", "+",
			"|", "+", "|",
			"+", "+", "+",
			"-",
			"|",
		},
	}

	// DefaultLineStyle default line style
	DefaultLineStyle *BorderStyle = printableLine

	// all supported styles are mapped with corresponding name *bs pairs
	styles map[borderStyleName]*BorderStyle = map[borderStyleName]*BorderStyle{
		BorderStyleDoubleLine:    doubleLine,
		BorderStyleSingleLine:    singleLine,
		BorderStylePrintableLine: printableLine,
	}
)

// get get char value
func (bs *BorderStyle) get(index int) string {
	return bs.chars[index]
}

// getStyle get predefined border style
func getStyle(name borderStyleName) (*BorderStyle, error) {
	style, exists := styles[name]
	if !exists {
		return nil, errors.New("stable error. border style not found")
	}
	return style, nil
}
