package stable

import (
	"strings"
)

// createBorders create borders for given style
func createBorders(bs *BorderStyle, columnSizeList []int) (string, string, string, string) {
	generic := buildGenericSeparator(columnSizeList)
	topSeperator := styleTheBorder(generic,
		bs.get(_BSITopLeft),
		bs.get(_BSITopCenter),
		bs.get(_BSITopRight),
		bs.get(_BSIVertical),
	)
	midSeperator := styleTheBorder(generic,
		bs.get(_BSIMidLeft),
		bs.get(_BSIMidCenter),
		bs.get(_BSIMidRight),
		bs.get(_BSIVertical),
	)
	botSeperator := styleTheBorder(generic,
		bs.get(_BSIBotLeft),
		bs.get(_BSIBotCenter),
		bs.get(_BSIBotRight),
		bs.get(_BSIVertical),
	)
	return generic, topSeperator, midSeperator, botSeperator
}

func styleTheBorder(generic string, left, center, right, horizontal string) string {
	sep := generic
	sep = strings.Replace(sep, _SPHHorizontal, horizontal, -1)
	sep = strings.Replace(sep, _SPHStart, left, -1)
	sep = strings.Replace(sep, _SPHMiddle, center, -1)
	sep = strings.Replace(sep, _SPHEnd, right, -1)
	return sep
}

func buildGenericSeparator(columnSizeList []int) string {
	s := _SPHStart
	for i, l := range columnSizeList {
		for j := 0; j < l; j++ {
			s += _SPHHorizontal
		}
		if i != len(columnSizeList)-1 {
			s += _SPHMiddle
		}
	}
	s += _SPHEnd
	return s
}
