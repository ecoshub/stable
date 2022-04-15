package stable

import (
	"strings"
)

func (st *STable) String() string {
	if !st.isThereAnyChanges() {
		return st.cache
	}

	if len(st.fields) == 0 {
		return "'stable' error. there is no field for this table"
	}

	st.calculateColumnSizeList()

	if len(st.columSizeList) == 0 {
		return "'stable' error. nothing to show"
	}

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

	s += st.createColumns()

	s += botBorder + "\n"
	st.cache = s
	st.changed = false
	return s
}

func (st *STable) createColumns() string {
	sep := st.borderStyle.get(borderStyleIndexHorizontal)
	s := ""
	for _, r := range st.rowValues {
		c := 0
		values := make([]string, 0, len(st.fields))
		for _, f := range st.fields {
			if f.IsHidden() {
				continue
			}
			val, _ := doPadding(r[c], st.columSizeList[c], f.opts.Alignment)
			values = append(values, val)
			c++
		}
		s += sep + strings.Join(values, sep) + sep + "\n"
	}
	return s
}

func (st *STable) calculateColumnSizeList() {
	st.columSizeList = make([]int, 0, len(st.fields))
	for _, f := range st.fields {
		if f.IsHidden() {
			continue
		}
		columnSize := len(addExtraPadding(f.name, st.generalPadding))
		st.columSizeList = append(st.columSizeList, columnSize)
	}
	if len(st.columSizeList) == 0 {
		return
	}
	st.rowValues = make([][]string, len(st.rows))
	for i := range st.rows {
		row := st.rows[i]
		st.rowValues[i] = make([]string, 0, len(row))
		c := 0
		for j := range row {
			f := st.GetField(j)
			if f.IsHidden() {
				continue
			}
			val := row[j]
			s := f.toString(val, st.generalPadding)
			if len(s) > st.columSizeList[c] {
				st.columSizeList[c] = len(s)
			}
			st.rowValues[i] = append(st.rowValues[i], s)
			c++
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
	c := 0
	values := make([]string, 0, len(st.fields))
	for _, f := range st.fields {
		if f.IsHidden() {
			continue
		}
		val, err := doPadding(f.GetName(), st.columSizeList[c], AlignmentCenter)
		if err != nil {
			continue
		}
		values = append(values, val)
		c++
	}
	s := sep + strings.Join(values, sep) + sep
	return s
}

func createColumn(bs *BorderStyle, fields []*Field, values []string, columnSizeList []int) string {
	sep := bs.get(borderStyleIndexHorizontal)
	s := sep
	c := 0
	for i, f := range fields {
		if f.IsHidden() {
			continue
		}
		var val string
		var err error
		if i < len(values) {
			// NOTE f and f.opts nil check maybe after
			val, err = doPadding(values[c], columnSizeList[c], f.opts.Alignment)
			if err != nil {
				val, _ = doPadding("<err>", columnSizeList[c], AlignmentCenter)
			}
		} else {
			val, _ = doPadding("", columnSizeList[c], AlignmentCenter)
		}
		s += val
		if i != len(fields)-1 {
			s += sep
		}
		c++
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
