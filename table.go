package stable

import (
	"fmt"
)

const (
	// DefaultGeneralPadding default general padding of table
	DefaultGeneralPadding int = 2
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
	bs, _ := getStyle(BorderStylePrintableLine)
	st := &STable{
		fields:         make([]*Field, 0, 8),
		rows:           make([][]interface{}, 0, 8),
		rowValues:      make([][]string, 0, 8),
		borderStyle:    bs,
		generalPadding: DefaultGeneralPadding,
	}
	st.SetCaption(caption)
	return st
}

// Basic creates basic table with field names
func Basic(caption string, fieldNames ...string) *STable {
	st := New(caption)
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
	st.changed = true
	return st
}

// SetCaption set caption for table
func (st *STable) SetCaption(caption string) {
	if len(caption) > 50 {
		caption = caption[:50]
		caption += "..."
	}
	st.caption = caption
	st.changed = true

}

// GetGeneralPadding get general table padding
func (st *STable) GetGeneralPadding() int {
	return st.generalPadding
}

// SetGeneralPadding set general table padding
func (st *STable) SetGeneralPadding(padding int) {
	st.generalPadding = padding
	st.changed = true
}

// Caption get caption of table
func (st *STable) Caption() string {
	return st.caption
}

// AddField adds a field with name
func (st *STable) AddField(name string) {
	f := newField(name)
	st.fields = append(st.fields, f)
	st.changed = true
}

// AddFieldWithOptions adds a field with options
func (st *STable) AddFieldWithOptions(name string, opts *Options) {
	f := NewFieldWithOptions(name, opts)
	st.fields = append(st.fields, f)
	st.changed = true
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
		err := fmt.Errorf("extra value(s) at row '%d'. value(s): %v", len(st.rows)+1, values[len(st.fields):])
		fmt.Println(err.Error())
		values = values[:len(st.fields)]
	}
	st.rows = append(st.rows, values)
	st.changed = true
}

// SetStyle set border style default is "printableBorderStyle"
func (st *STable) SetStyle(styleName borderStyleName) error {
	style, err := getStyle(styleName)
	if err != nil {
		return err
	}
	st.borderStyle = style
	st.changed = true
	return nil
}
