package stable

import (
	"fmt"
)

// STable simple table main struct
type STable struct {
	caption       string
	fields        []*Field
	rows          [][]interface{}
	rowValues     [][]string
	columSizeList []int
	borderStyle   *BorderStyle
}

// New new table, first param is table caption if given
func New(caption string) *STable {
	bs, _ := getStyle(BorderStylePrintableLine)
	st := &STable{
		fields:      make([]*Field, 0, 8),
		rows:        make([][]interface{}, 0, 8),
		rowValues:   make([][]string, 0, 8),
		borderStyle: bs,
	}
	st.SetCaption(caption)
	return st
}

// Basic creates basic table with field names
func Basic(fieldNames ...string) *STable {
	st := New("")
	for _, fn := range fieldNames {
		st.AddField(fn)
	}
	return st
}

// AddFields adds fields
func (st *STable) AddFields(fieldNames ...string) *STable {
	for _, fn := range fieldNames {
		st.AddField(fn)
	}
	return st
}

// SetCaption set caption for table
func (st *STable) SetCaption(caption string) {
	if len(caption) > 50 {
		caption = caption[:50]
		caption += "..."
	}
	st.caption = caption
}

// Caption get caption of table
func (st *STable) Caption() string {
	return st.caption
}

// AddField adds a field with name
func (st *STable) AddField(name string) {
	f := NewField(name)
	st.fields = append(st.fields, f)
}

// AddFieldWithOptions adds a field with options
func (st *STable) AddFieldWithOptions(name string, opts *Options) {
	f := NewFieldWithOptions(name, opts)
	st.fields = append(st.fields, f)
}

// GetField GetField
func (st *STable) GetField(index int) *Field {
	if index < len(st.fields) {
		return st.fields[index]
	}
	return nil
}

// GetFieldWithName GetFieldWithName
func (st *STable) GetFieldWithName(name string) *Field {
	for _, f := range st.fields {
		if f.name == name {
			return f
		}
	}
	return nil
}

// Row add row
func (st *STable) Row(values ...interface{}) {
	if len(values) > len(st.fields) {
		err := fmt.Errorf("exrta value(s) at row '%d'. value(s): %v", len(st.rows)+1, values[len(st.fields):])
		fmt.Println(err.Error())
		values = values[:len(st.fields)]
	}
	st.rows = append(st.rows, values)
}

// SetStyle set border style default is "printableBorderStyle"
func (st *STable) SetStyle(styleName borderStyleName) error {
	styl, err := getStyle(styleName)
	if err != nil {
		return err
	}
	st.borderStyle = styl
	return nil
}

func (st *STable) getFieldNames() []string {
	n := make([]string, len(st.fields))
	for i, f := range st.fields {
		n[i] = f.GetName()
	}
	return n
}

func addPaddingToValues(values []string) []string {
	padded := make([]string, len(values))
	for i := range values {
		padded[i] = addExtraPadding(values[i])
	}
	return padded
}

func (st *STable) calculateColumnSizeList() {
	st.columSizeList = make([]int, len(st.fields))
	for i, f := range st.fields {
		st.columSizeList[i] = len(addExtraPadding(f.GetName()))
	}
	st.rowValues = make([][]string, len(st.rows))
	for i := range st.rows {
		row := st.rows[i]
		st.rowValues[i] = make([]string, len(row))
		for j := range row {
			f := st.GetField(j)
			val := row[j]
			s := f.toString(val)
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
	colunmSize := len(st.columSizeList)
	captionLength := len(st.caption)
	for i := range st.columSizeList {
		st.columSizeList[i] += captionLength / colunmSize
	}
}

func (st *STable) createHeader() string {
	sep := st.borderStyle.get(borderStyleIndexHorizontal)
	s := sep
	for i, f := range st.fields {
		val, err := doPadding(f.GetName(), st.columSizeList[i], AlignementCenter)
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
			val, err = doPadding(values[i], columnSizeList[i], f.GetAlignement())
			if err != nil {
				val, _ = doPadding("ERR", columnSizeList[i], AlignementCenter)
			}
		} else {
			val, _ = doPadding("", columnSizeList[i], AlignementCenter)
		}
		s += val
		if i != len(fields)-1 {
			s += sep
		}
	}
	s += sep
	return s
}

func (st *STable) String() string {
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
	return s
}

func createCaptionBar(bs *BorderStyle, caption string, genericBorder string) string {
	tot := len(genericBorder) - 2
	caption, _ = doPadding(caption, tot, AlignementCenter)

	captionTopBorder := styleTheBorder(genericBorder,
		bs.get(borderStyleIndexTopLeft),
		bs.get(borderStyleIndexVertical),
		bs.get(borderStyleIndexTopRight),
		bs.get(borderStyleIndexVertical),
	)
	captionMiddleBordder := styleTheBorder(genericBorder,
		bs.get(borderStyleIndexMidLeft),
		bs.get(borderStyleIndexTopCenter),
		bs.get(borderStyleIndexMidRight),
		bs.get(borderStyleIndexVertical),
	)
	hor := bs.get(borderStyleIndexHorizontal)
	s := ""
	s += captionTopBorder + "\n"
	s += hor + caption + hor + "\n"
	s += captionMiddleBordder + "\n"
	return s
}
