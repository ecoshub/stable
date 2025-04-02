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
	columnSizeList, rowValues, err := st.calculateColumnSizeList()
	if err != nil {
		return err.Error()
	}

	// if colum size calculation return with zero length
	// it means there is no row to process
	if len(columnSizeList) == 0 {
		return ErrNoRow.Error()
	}

	// calculate and create generic borders for this table
	generic, topBorder, midBorder, botBorder := createBorders(st.borderStyle, columnSizeList)

	// table string
	str := ""

	if st.caption != "" {
		// create caption bar
		str += createCaptionBar(st.borderStyle, st.caption, generic)
	} else {
		str += topBorder + "\n"
	}
	// create header bar (field names)
	headerBar := st.createHeader(columnSizeList)
	str += headerBar

	// create all rows
	columnBars := st.createColumns(columnSizeList, rowValues)
	str += midBorder + "\n"
	str += columnBars

	// Check if any field has a summary configured
	hasSummary := false
	for _, field := range st.fields {
		if field.HasSummary() {
			hasSummary = true
			break
		}
	}

	// add summary row if enabled and there are fields with summary configuration
	if st.showSummary && hasSummary {
		summaryRow := st.getSummaryRow()
		if summaryRow != nil {
			str += midBorder + "\n"
			summaryValues := make([]string, 0, len(summaryRow))
			c := 0
			for i, val := range summaryRow {
				if st.fields[i].IsHidden() {
					continue
				}
				s := st.fields[i].toString(val, st.generalPadding)
				s = doPadding(s, columnSizeList[c], st.fields[i].opts.Alignment)
				summaryValues = append(summaryValues, s)
				c++
			}
			sep := st.borderStyle.get(_BSIHorizontal)
			str += sep + strings.Join(summaryValues, sep) + sep + "\n"
		}
	}

	str += botBorder
	str += "\n"

	st.cache = str
	st.changed = false
	return str
}

func (st *STable) createColumns(columnSizeList []int, rowValues [][]string) string {
	sep := st.borderStyle.get(_BSIHorizontal)
	s := ""
	for _, r := range rowValues {
		values := make([]string, 0, len(st.fields))
		visibleIndex := 0
		for i, f := range st.fields {
			if f.IsHidden() {
				continue
			}
			val := doPadding(r[i], columnSizeList[visibleIndex], f.opts.Alignment)
			values = append(values, val)
			visibleIndex++
		}
		s += sep + strings.Join(values, sep) + sep + "\n"
	}
	return s
}

func (st *STable) calculateColumnSizeList() ([]int, [][]string, error) {
	columnSizeList := make([]int, 0, len(st.fields))
	for _, f := range st.fields {
		if f.IsHidden() {
			continue
		}
		columnSize := len(addExtraPadding(f.name, st.generalPadding))
		columnSizeList = append(columnSizeList, columnSize)
	}
	if len(columnSizeList) == 0 {
		return nil, nil, errors.New("no column defined")
	}
	rowValues := make([][]string, len(st.rows))
	for i := range st.rows {
		row := st.rows[i]
		rowValues[i] = make([]string, len(row)) // Initialize with full length
		visibleIndex := 0
		for j := range row {
			f := st.GetField(j)
			if f.IsHidden() {
				continue
			}
			val := row[j]
			s := f.toString(val, st.generalPadding)
			if len(s) > columnSizeList[visibleIndex] {
				columnSizeList[visibleIndex] = len(s)
			}
			rowValues[i][j] = s // Store at original index
			visibleIndex++
		}
	}

	// Consider summary values in column size calculation
	if st.showSummary {
		summaryRow := st.getSummaryRow()
		if summaryRow != nil {
			visibleIndex := 0
			for i, val := range summaryRow {
				if st.fields[i].IsHidden() {
					continue
				}
				s := st.fields[i].toString(val, st.generalPadding)
				if len(s) > columnSizeList[visibleIndex] {
					columnSizeList[visibleIndex] = len(s)
				}
				visibleIndex++
			}
		}
	}

	// adjust column sizes depending on caption size
	st.adjustColumnSizes(columnSizeList)
	return columnSizeList, rowValues, nil
}

func (st *STable) adjustColumnSizes(columnSizeList []int) {
	captionLength := len(addExtraPadding(st.caption, st.generalPadding)) + 2
	tot := 0
	for i := range columnSizeList {
		tot += columnSizeList[i]
	}
	if tot < captionLength {
		diff := captionLength - tot
		for diff > 0 {
			for i := range columnSizeList {
				columnSizeList[i]++
				diff--
			}
		}
	}
}

func (st *STable) createHeader(columnSizeList []int) string {
	sep := st.borderStyle.get(_BSIHorizontal)
	c := 0
	values := make([]string, 0, len(st.fields))
	for _, f := range st.fields {
		if f.IsHidden() {
			continue
		}
		val := addExtraPadding(f.name, st.generalPadding)
		val = doPadding(val, columnSizeList[c], f.opts.HeaderAlignment)
		values = append(values, val)
		c++
	}
	s := sep + strings.Join(values, sep) + sep + "\n"
	return s
}

func createCaptionBar(bs *BorderStyle, caption string, genericBorder string) string {
	tot := len(genericBorder) - 2
	caption = doPadding(caption, tot, AlignmentCenter)

	captionTopBorder := styleTheBorder(genericBorder,
		bs.get(_BSITopLeft),
		bs.get(_BSIVertical),
		bs.get(_BSITopRight),
		bs.get(_BSIVertical),
	)

	captionMiddleBorder := styleTheBorder(genericBorder,
		bs.get(_BSIMidLeft),
		bs.get(_BSITopCenter),
		bs.get(_BSIMidRight),
		bs.get(_BSIVertical),
	)

	hor := bs.get(_BSIHorizontal)
	s := ""
	s += captionTopBorder + "\n"
	s += hor + caption + hor + "\n"
	s += captionMiddleBorder + "\n"
	return s
}
