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
	name    string
	opts    *Options
	changed bool
}

// Options field options
type Options struct {
	Format          string
	Alignment       alignment
	HeaderAlignment alignment
	Hide            bool
	CharLimit       int
	LimitFromStart  bool
}

// creates a new field with name
func newField(name string) *Field {
	return &Field{name: name, opts: &Options{
		Alignment:       DefaultValueAlignment,
		HeaderAlignment: DefaultHeaderAlignment,
	},
	}
}

// NewFieldWithOptions NewFieldWithOptions
func NewFieldWithOptions(name string, opts *Options) *Field {
	if opts == nil {
		return newField(name)
	}
	if opts.Alignment == "" {
		opts.Alignment = DefaultValueAlignment
		opts.HeaderAlignment = DefaultHeaderAlignment
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
	f.changed = true
}

// SetHeaderAlignment SetHeaderAlignment.
func (f *Field) SetHeaderAlignment(alignment alignment) {
	if f == nil {
		fmt.Println(FieldIsNil)
		return
	}
	f.opts.HeaderAlignment = alignment
	f.changed = true
}

// GetAlignment GetAlignment
func (f *Field) GetAlignment() string {
	if f == nil {
		return FieldIsNil
	}
	if f.opts == nil {
		return string(DefaultValueAlignment)
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
	f.changed = true
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

// SetOptions set field option.
func (f *Field) SetOptions(opts *Options) {
	if f == nil {
		fmt.Println(FieldIsNil)
		return
	}
	if opts.Alignment == "" {
		opts.Alignment = DefaultValueAlignment
	}
	if opts.HeaderAlignment == "" {
		opts.HeaderAlignment = DefaultHeaderAlignment
	}
	f.opts = opts
	f.changed = true

}

// IsHidden get visibility of field
func (f *Field) IsHidden() bool {
	return f.opts.Hide == true
}

// ChangeVisibility change visibility of field
func (f *Field) ChangeVisibility(hide bool) {
	f.opts.Hide = hide
	f.changed = true
}

// Show change visibility of field to 'show'
func (f *Field) Show() {
	f.ChangeVisibility(false)
}

// Hide change visibility of field to 'hide'
func (f *Field) Hide() {
	f.ChangeVisibility(true)
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
	if f.opts.CharLimit > 0 {
		if f.opts.CharLimit < len(v) {
			if f.opts.LimitFromStart {
				v = "..." + v[len(v)-f.opts.CharLimit:]
			} else {
				v = v[:f.opts.CharLimit] + "..."
			}
		}
	}
	v = addExtraPadding(v, paddingAmount)
	return v
}
