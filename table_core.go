package stable

import "fmt"

func (st *STable) String() string {
	if !st.isThereAnyChanges() {
		return st.cache
	}

	if len(st.fields) == 0 {
		return "error. Add fields to table"
	}

	st.calculateColumnSizeList()
	generic, topBorder, midBorder, botBorder := createBorders(st.borderStyle, st.columSizeList)

	s := ""
	if st.caption != "" {
		s += createCaptionBar(st.borderStyle, st.caption, generic)
	} else {
		s += topBorder + "\n"
	}

	h := st.createHeader()
	s += h + "\n"
	s += midBorder + "\n"

	for _, r := range st.rowValues {
		c := createColumn(st.borderStyle, st.fields, r, st.columSizeList)
		s += c + "\n"
	}
	s += botBorder + "\n"
	st.cache = s
	st.changed = false
	return s
}

func (st *STable) calculateColumnSizeList() {
	st.columSizeList = make([]int, len(st.fields))
	for i, f := range st.fields {
		st.columSizeList[i] = len(addExtraPadding(f.name, st.generalPadding))
	}
	st.rowValues = make([][]string, len(st.rows))
	for i := range st.rows {
		row := st.rows[i]
		st.rowValues[i] = make([]string, len(row))
		for j := range row {
			f := st.GetField(j)
			val := row[j]
			s := f.toString(val, st.generalPadding)
			if len(s) > st.columSizeList[j] {
				st.columSizeList[j] = len(s)
			}
			st.rowValues[i][j] = s
		}
	}
	// adjust column sizes depending on caption size
	st.adjustColumnSizes()
}

func (st *STable) adjustColumnSizes() {
	captionLength := len(addExtraPadding(st.caption, st.generalPadding)) + 2
	tot := 0
	for i := range st.columSizeList {
		tot += st.columSizeList[i]
	}
	if tot < captionLength {
		diff := captionLength - tot
		for diff > 0 {
			for i := range st.columSizeList {
				st.columSizeList[i]++
				diff--
			}
		}
	}
}

func (st *STable) createHeader() string {
	sep := st.borderStyle.get(borderStyleIndexHorizontal)
	s := sep
	for i, f := range st.fields {
		val, err := doPadding(f.GetName(), st.columSizeList[i], AlignmentCenter)
		if err != nil {
			fmt.Println(err)
			continue
		}
		s += val
		if i != len(st.fields)-1 {
			s += sep
		}
	}
	s += sep
	return s
}

func createColumn(bs *BorderStyle, fields []*Field, values []string, columnSizeList []int) string {
	sep := bs.get(borderStyleIndexHorizontal)
	s := sep
	for i, f := range fields {
		var val string
		var err error
		if i < len(values) {
			// NOTE f and f.opts nil check maybe after
			val, err = doPadding(values[i], columnSizeList[i], f.opts.Alignment)
			if err != nil {
				fmt.Println(err)
				val, _ = doPadding("ERR", columnSizeList[i], AlignmentCenter)
			}
		} else {
			val, _ = doPadding("", columnSizeList[i], AlignmentCenter)
		}
		s += val
		if i != len(fields)-1 {
			s += sep
		}
	}
	s += sep
	return s
}

func createCaptionBar(bs *BorderStyle, caption string, genericBorder string) string {
	tot := len(genericBorder) - 2
	caption, _ = doPadding(caption, tot, AlignmentCenter)

	captionTopBorder := styleTheBorder(genericBorder,
		bs.get(borderStyleIndexTopLeft),
		bs.get(borderStyleIndexVertical),
		bs.get(borderStyleIndexTopRight),
		bs.get(borderStyleIndexVertical),
	)
	captionMiddleBorder := styleTheBorder(genericBorder,
		bs.get(borderStyleIndexMidLeft),
		bs.get(borderStyleIndexTopCenter),
		bs.get(borderStyleIndexMidRight),
		bs.get(borderStyleIndexVertical),
	)
	hor := bs.get(borderStyleIndexHorizontal)
	s := ""
	s += captionTopBorder + "\n"
	s += hor + caption + hor + "\n"
	s += captionMiddleBorder + "\n"
	return s
}
