package stable

import (
	"fmt"
)

const (
	// DefaultGeneralPadding default general padding of table
	DefaultGeneralPadding int = 2

	// MaxCaptionLength max caption length
	// if its greater than this constant it will trim the end with "..."
	MaxCaptionLength int = 50

	// MaxGeneralPadding maximum general padding
	MaxGeneralPadding int = 8
)

// STable simple table main struct
type STable struct {
	caption        string
	fields         []*Field
	rows           [][]interface{}
	rowValues      [][]string
	columSizeList  []int
	borderStyle    *BorderStyle
	generalPadding int
	changed        bool
	cache          string
}

// New new table, first param is table caption if given
func New(caption string) *STable {
	return &STable{
		caption:        processCaption(caption),
		fields:         make([]*Field, 0, 8),
		rows:           make([][]interface{}, 0, 8),
		rowValues:      make([][]string, 0, 8),
		borderStyle:    DefaultLineStyle,
		generalPadding: DefaultGeneralPadding,
	}
}

// Basic creates basic table with field names
func Basic(caption string, fieldNames ...string) *STable {
	fields := make([]*Field, len(fieldNames))
	for i := range fields {
		fields[i] = newField(fieldNames[i])
	}
	st := New(caption)
	st.fields = fields
	return st
}

// AddField adds a field with name
func (st *STable) AddField(name string) {
	f := newField(name)
	st.fields = append(st.fields, f)
	st.changed = true
}

// AddFields add new fields with name
func (st *STable) AddFields(fieldNames ...string) *STable {
	for _, fn := range fieldNames {
		st.AddField(fn)
	}
	st.changed = true
	return st
}

// AddFieldWithOptions adds a field with options
func (st *STable) AddFieldWithOptions(name string, opts *Options) {
	f := NewFieldWithOptions(name, opts)
	st.fields = append(st.fields, f)
	st.changed = true
}

// Caption get caption of table
func (st *STable) Caption() string {
	return st.caption
}

// SetCaption set caption for table
func (st *STable) SetCaption(caption string) {
	st.caption = processCaption(caption)
	st.changed = true
}

// GetGeneralPadding get general table padding
func (st *STable) GetGeneralPadding() int {
	return st.generalPadding
}

// SetGeneralPadding set general table padding
func (st *STable) SetGeneralPadding(padding int) {
	st.generalPadding = processPadding(padding)
	st.changed = true
}

// GetField get field with field index
func (st *STable) GetField(index int) *Field {
	if index < len(st.fields) {
		return st.fields[index]
	}
	return nil
}

// GetFieldByName get field by field name
func (st *STable) GetFieldByName(name string) *Field {
	for _, f := range st.fields {
		if f.name == name {
			return f
		}
	}
	return nil
}

// Row add a row
func (st *STable) Row(values ...interface{}) {
	if len(values) > len(st.fields) {
		err := fmt.Errorf("extra value(s) at row '%d'. value(s): %v", len(st.rows)+1, values[len(st.fields):])
		fmt.Println(err.Error())
		values = values[:len(st.fields)]
	}
	st.rows = append(st.rows, values)
	st.changed = true
}

// SetStyle set border style (BorderStyleDoubleLine | BorderStyleSingleLine | BorderStylePrintableLine)
func (st *STable) SetStyle(styleName borderStyleName) error {
	style, err := getStyle(styleName)
	if err != nil {
		return err
	}
	st.borderStyle = style
	st.changed = true
	return nil
}

// FieldCount get field count of table
func (st *STable) FieldCount() int {
	return len(st.fields)
}
