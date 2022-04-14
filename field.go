package stable

import (
	"fmt"
)

const (
	// DefaultFieldName dedfault field name
	DefaultFieldName string = "unamed_field"
)

// Field Field object
type Field struct {
	name string
	opts *Options
}

// Options field options
type Options struct {
	Format     string
	Alignement string
}

// NewField new field with only name option
func NewField(name string) *Field {
	return &Field{name: name, opts: &Options{Alignement: AlignementLeft}}
}

// NewFieldWithOptions NewFieldWithOptions
func NewFieldWithOptions(name string, opts *Options) *Field {
	if opts.Alignement == "" {
		opts.Alignement = DefaultAlignementForValues
	}
	return &Field{
		name: name,
		opts: opts,
	}
}

// GetAlignement GetAlignement
func (f *Field) GetAlignement() string {
	if f == nil {
		return "field is null"
	}
	if f.opts == nil {
		return "unknown_padding"
	}
	return f.opts.Alignement
}

// SetAlignement SetAlignement.
func (f *Field) SetAlignement(alignement string) {
	if f == nil {
		fmt.Println("field is null. SetAlignement()")
		return
	}
	f.opts.Alignement = alignement
}

// AlignCenter easy access alignement choices
func (f *Field) AlignCenter() {
	f.SetAlignement(AlignementCenter)
}

// AlignLeft easy access alignement choices
func (f *Field) AlignLeft() {
	f.SetAlignement(AlignementLeft)
}

// AlignRight easy access alignement choices
func (f *Field) AlignRight() {
	f.SetAlignement(AlignementRight)
}

// GetName GetName
func (f *Field) GetName() string {
	if f == nil {
		return "field is null"
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

func (f *Field) toString(value interface{}) string {
	if f == nil {
		return "field is null"
	}
	v := ""
	if f.opts.Format != "" {
		v = fmt.Sprintf(f.opts.Format, value)
	} else {
		v = fmt.Sprint(value)
	}
	return addExtraPadding(v)
}
