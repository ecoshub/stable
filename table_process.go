package stable

import (
	"strings"
)

// createBorders create borders for given style
func createBorders(bs *BorderStyle, columnSizeList []int) (string, string, string, string) {
	generic := buildGenericSeparator(columnSizeList)
	topSeperator := styleTheBorder(generic,
		bs.get(borderStyleIndexTopLeft),
		bs.get(borderStyleIndexTopCenter),
		bs.get(borderStyleIndexTopRight),
		bs.get(borderStyleIndexVertical),
	)
	midSeperator := styleTheBorder(generic,
		bs.get(borderStyleIndexMidLeft),
		bs.get(borderStyleIndexMidCenter),
		bs.get(borderStyleIndexMidRight),
		bs.get(borderStyleIndexVertical),
	)
	botSeperator := styleTheBorder(generic,
		bs.get(borderStyleIndexBotLeft),
		bs.get(borderStyleIndexBotCenter),
		bs.get(borderStyleIndexBotRight),
		bs.get(borderStyleIndexVertical),
	)
	return generic, topSeperator, midSeperator, botSeperator
}

func styleTheBorder(generic string, left, center, right, horizontal string) string {
	sep := generic
	sep = strings.Replace(sep, seperatorHorizontalPlaceHolder, horizontal, -1)
	sep = strings.Replace(sep, seperatorStartPlaceHolder, left, -1)
	sep = strings.Replace(sep, seperatorMiddlePlaceHolder, center, -1)
	sep = strings.Replace(sep, seperatorEndPlaceHolder, right, -1)
	return sep
}

func buildGenericSeparator(columnSizeList []int) string {
	s := seperatorStartPlaceHolder
	for i, l := range columnSizeList {
		for j := 0; j < l; j++ {
			s += seperatorHorizontalPlaceHolder
		}
		if i != len(columnSizeList)-1 {
			s += seperatorMiddlePlaceHolder
		}
	}
	s += seperatorEndPlaceHolder
	return s
}
