package stable

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrNoField no field errors
	ErrNoField error = errors.New("'stable' error. there is no field on this table")

	// ErrNoRow no row error
	ErrNoRow error = errors.New("'stable' error. no row to show")
)

// Print print the given type as table
func Print(i interface{}) {
	st, err := ToTable(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(st)
}

// PrintWithCaption print the given type as table with caption
func PrintWithCaption(caption string, i interface{}) {
	st, err := ToTable(i)
	if err != nil {
		fmt.Println(err)
		return
	}

	st.SetCaption(caption)

	fmt.Println(st)
}

// String table to string
func (st *STable) String() string {
	// if no changes occurred return the cache
	if !st.isThereAnyChanges() {
		return st.cache
	}

	if len(st.fields) == 0 {
		return ErrNoField.Error()
	}

	// iterate all columns and get the logest ones to determin min column size
	st.calculateColumnSizeList()

	// if colum size calculation return with zero length
	// it means there is no row to process
	if len(st.columSizeList) == 0 {
		return ErrNoRow.Error()
	}

	// calculate and create generic borders for this table
	generic, topBorder, midBorder, botBorder := createBorders(st.borderStyle, st.columSizeList)

	// table string
	str := ""

	if st.caption != "" {
		// create caption bar
		str += createCaptionBar(st.borderStyle, st.caption, generic)
	} else {
		str += topBorder + "\n"
	}
	// create header bar (field names)
	headerBar := st.createHeader()
	str += headerBar

	// create all rows
	columnBars := st.createColumns()
	str += midBorder + "\n"
	str += columnBars
	str += botBorder
	str += "\n"

	st.cache = str
	st.changed = false
	return str
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
			val := doPadding(r[c], st.columSizeList[c], f.opts.Alignment)
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
		val := addExtraPadding(f.name, st.generalPadding)
		val = doPadding(val, st.columSizeList[c], f.opts.HeaderAlignment)
		values = append(values, val)
		c++
	}
	s := sep + strings.Join(values, sep) + sep + "\n"
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
		if i < len(values) {
			// NOTE f and f.opts nil check maybe after
			val = doPadding(values[c], columnSizeList[c], f.opts.Alignment)
		} else {
			val = doPadding("", columnSizeList[c], AlignmentCenter)
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
	caption = doPadding(caption, tot, AlignmentCenter)

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
