package style

import (
	"strings"
)

// CreateBorders create borders for given style
func CreateBorders(bs *BorderStyle, columnSizeList []int) (string, string, string, string) {
	generic := buildGenericSeparator(columnSizeList)
	topSeperator := StyleTheBorder(generic,
		bs.Get(BorderStyleIndexTopLeft),
		bs.Get(BorderStyleIndexTopCenter),
		bs.Get(BorderStyleIndexTopRight),
		bs.Get(BorderStyleIndexVertical),
	)
	midSeperator := StyleTheBorder(generic,
		bs.Get(BorderStyleIndexMidLeft),
		bs.Get(BorderStyleIndexMidCenter),
		bs.Get(BorderStyleIndexMidRight),
		bs.Get(BorderStyleIndexVertical),
	)
	botSeperator := StyleTheBorder(generic,
		bs.Get(BorderStyleIndexBotLeft),
		bs.Get(BorderStyleIndexBotCenter),
		bs.Get(BorderStyleIndexBotRight),
		bs.Get(BorderStyleIndexVertical),
	)
	return generic, topSeperator, midSeperator, botSeperator
}

func StyleTheBorder(generic string, left, center, right, horizontal string) string {
	sep := generic
	sep = strings.Replace(sep, SeperatorHorizontalPlaceHolder, horizontal, -1)
	sep = strings.Replace(sep, SeperatorStartPlaceHolder, left, -1)
	sep = strings.Replace(sep, SeperatorMiddlePlaceHolder, center, -1)
	sep = strings.Replace(sep, SeperatorEndPlaceHolder, right, -1)
	return sep
}

func buildGenericSeparator(columnSizeList []int) string {
	s := SeperatorStartPlaceHolder
	for i, l := range columnSizeList {
		for j := 0; j < l; j++ {
			s += SeperatorHorizontalPlaceHolder
		}
		if i != len(columnSizeList)-1 {
			s += SeperatorMiddlePlaceHolder
		}
	}
	s += SeperatorEndPlaceHolder
	return s
}
