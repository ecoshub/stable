package table

import (
	"errors"
	"fmt"
	"stable/field"
	"stable/process"
	"stable/style"
)

// STable simple table main struct
type STable struct {
	caption       string
	fields        []*field.Field
	rows          [][]interface{}
	rowValues     [][]string
	columSizeList []int
	borderStyle   *style.BorderStyle
}

// New new table, first param is table caption if given
func New(caption string) *STable {
	bs, _ := style.GetStyle(style.BorderStylePrintableLine)
	st := &STable{
		fields:      make([]*field.Field, 0, 8),
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
	f := field.NewField(name)
	st.fields = append(st.fields, f)
}

// AddFieldWithOptions adds a field with options
func (st *STable) AddFieldWithOptions(name string, opts *field.Options) {
	f := field.NewFieldWithOptions(name, opts)
	st.fields = append(st.fields, f)
}

func (st *STable) SetFields(fields ...*field.Field) {
	st.fields = fields
}

func (st *STable) GetField(index int) *field.Field {
	if index < len(st.fields) {
		return st.fields[index]
	}
	return nil
}

func (st *STable) GetFieldWithName(name string) *field.Field {
	for _, f := range st.fields {
		if f.Name() == name {
			return f
		}
	}
	return nil
}

func (st *STable) SetField(index int, f *field.Field) error {
	if index < len(st.fields) {
		st.fields[index] = f
		return nil
	}
	return errors.New("field not found")
}

func (st *STable) Row(values ...interface{}) {
	if len(values) > len(st.fields) {
		err := fmt.Errorf("exrta value(s) at row '%d'. value(s): %v", len(st.rows)+1, values[len(st.fields):])
		fmt.Println(err.Error())
		values = values[:len(st.fields)]
	}
	st.rows = append(st.rows, values)
}

func (st *STable) SetStyle(styleName style.BorderStyleName) error {
	styl, err := style.GetStyle(styleName)
	if err != nil {
		return err
	}
	st.borderStyle = styl
	return nil
}

func (st *STable) getFieldNames() []string {
	n := make([]string, len(st.fields))
	for i, f := range st.fields {
		n[i] = f.Name()
	}
	return n
}

func addPaddingToValues(values []string) []string {
	padded := make([]string, len(values))
	for i := range values {
		padded[i] = process.AddExtraPadding(values[i])
	}
	return padded
}

func (st *STable) calculateColumnSizeList() {
	st.columSizeList = make([]int, len(st.fields))
	for i, f := range st.fields {
		st.columSizeList[i] = len(process.AddExtraPadding(f.Name()))
	}
	st.rowValues = make([][]string, len(st.rows))
	for i := range st.rows {
		row := st.rows[i]
		st.rowValues[i] = make([]string, len(row))
		for j := range row {
			f := st.GetField(j)
			val := row[j]
			s := f.ToString(val)
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
	sep := st.borderStyle.Get(style.BorderStyleIndexHorizontal)
	s := sep
	for i, f := range st.fields {
		val, err := style.DoPadding(f.Name(), st.columSizeList[i], style.AlignementCenter)
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

func createColumn(bs *style.BorderStyle, fields []*field.Field, values []string, columnSizeList []int) string {
	sep := bs.Get(style.BorderStyleIndexHorizontal)
	s := sep
	for i, f := range fields {
		var val string
		var err error
		if i < len(values) {
			val, err = style.DoPadding(values[i], columnSizeList[i], f.Padding())
			if err != nil {
				val, _ = style.DoPadding("ERR", columnSizeList[i], style.AlignementCenter)
			}
		} else {
			val, _ = style.DoPadding("", columnSizeList[i], style.AlignementCenter)
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
	generic, topBorder, midBorder, botBorder := style.CreateBorders(st.borderStyle, st.columSizeList)

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

func createCaptionBar(bs *style.BorderStyle, caption string, genericBorder string) string {
	tot := len(genericBorder) - 2
	caption, _ = style.DoPadding(caption, tot, style.AlignementCenter)

	captionTopBorder := style.StyleTheBorder(genericBorder,
		bs.Get(style.BorderStyleIndexTopLeft),
		bs.Get(style.BorderStyleIndexVertical),
		bs.Get(style.BorderStyleIndexTopRight),
		bs.Get(style.BorderStyleIndexVertical),
	)
	captionMiddleBordder := style.StyleTheBorder(genericBorder,
		bs.Get(style.BorderStyleIndexMidLeft),
		bs.Get(style.BorderStyleIndexTopCenter),
		bs.Get(style.BorderStyleIndexMidRight),
		bs.Get(style.BorderStyleIndexVertical),
	)
	hor := bs.Get(style.BorderStyleIndexHorizontal)
	s := ""
	s += captionTopBorder + "\n"
	s += hor + caption + hor + "\n"
	s += captionMiddleBordder + "\n"
	return s
}
