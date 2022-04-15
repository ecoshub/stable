package stable

import (
	"fmt"
	"strings"
)

const (
	// NilValueString if a value is nil prints this
	NilValueString string = "-"
	// FieldIsNil if field pointer is nil return this line
	FieldIsNil string = "field is nil."
)

// Field object
// every field object is a column
type Field struct {
	name string
	opts *Options
}

// Options field options
type Options struct {
	Format    string
	Alignment alignment
}

// creates a new field with name
func newField(name string) *Field {
	return &Field{name: name, opts: &Options{Alignment: DefaultAlignment}}
}

// NewFieldWithOptions NewFieldWithOptions
func NewFieldWithOptions(name string, opts *Options) *Field {
	if opts == nil {
		return newField(name)
	}
	if opts.Alignment == "" {
		opts.Alignment = DefaultAlignment
	}
	return &Field{
		name: name,
		opts: opts,
	}
}

// GetName GetName
func (f *Field) GetName() string {
	if f == nil {
		return FieldIsNil
	}
	return f.name
}

// SetName SetName
func (f *Field) SetName(name string) {
	if f == nil {
		return
	}
	f.name = name
}

// GetAlignment GetAlignment
func (f *Field) GetAlignment() string {
	if f == nil {
		return FieldIsNil
	}
	if f.opts == nil {
		return string(DefaultAlignment)
	}
	return string(f.opts.Alignment)
}

// SetAlignment SetAlignment.
func (f *Field) SetAlignment(alignment alignment) {
	if f == nil {
		fmt.Println(FieldIsNil)
		return
	}
	f.opts.Alignment = alignment
}

// AlignCenter easy access alignment choices
func (f *Field) AlignCenter() {
	f.SetAlignment(AlignmentCenter)
}

// AlignLeft easy access alignment choices
func (f *Field) AlignLeft() {
	f.SetAlignment(AlignmentLeft)
}

// AlignRight easy access alignment choices
func (f *Field) AlignRight() {
	f.SetAlignment(AlignmentRight)
}

// SetOption set field option.
func (f *Field) SetOption(opts *Options) {
	if f == nil {
		fmt.Println(FieldIsNil)
		return
	}
	if opts.Alignment == "" {
		opts.Alignment = DefaultAlignment
	}
	f.opts = opts
}

func (f *Field) toString(value interface{}, paddingAmount int) string {
	if value == nil {
		value = NilValueString
	}
	v := ""
	if f.opts.Format != "" {
		v = fmt.Sprintf(f.opts.Format, value)
	} else {
		v = fmt.Sprint(value)
	}
	v = strings.Replace(v, "\t", "    ", -1)
	return addExtraPadding(v, paddingAmount)
}
