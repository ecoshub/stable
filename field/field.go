package field

import (
	"fmt"
	"stable/process"
	"stable/style"
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

// NewField new field with only name option
func NewField(name string) *Field {
	return &Field{name: name, opts: &Options{Alignement: style.AlignementLeft}}
}

// Options field options
type Options struct {
	Format     string
	Alignement string
}

func (f *Field) Padding() string {
	if f == nil {
		return "field is null"
	}
	if f.opts == nil {
		return "unknown_padding"
	}
	return f.opts.Alignement
}

// NewFieldWithOptions NewFieldWithOptions
func NewFieldWithOptions(name string, opts *Options) *Field {
	if opts.Alignement == "" {
		opts.Alignement = style.DefaultAlignementForValues
	}
	return &Field{
		name: name,
		opts: opts,
	}
}

func (f *Field) Alignement() string {
	if f == nil {
		return "field is null"
	}
	return f.opts.Alignement
}

func (f *Field) SetAlignement(alignement string) {
	if f == nil {
		fmt.Println("field is null. SetAlignement()")
		return
	}
	f.opts.Alignement = alignement
}

func (f *Field) AlignCenter() {
	f.SetAlignement(style.AlignementCenter)
}

func (f *Field) AlignLeft() {
	f.SetAlignement(style.AlignementLeft)
}

func (f *Field) AlignRight() {
	f.SetAlignement(style.AlignementRight)
}

func (f *Field) Name() string {
	if f == nil {
		return "field is null"
	}
	return f.name
}

func (f *Field) SetName(name string) {
	if f == nil {
		return
	}
	f.name = name
}

func (f *Field) ToString(value interface{}) string {
	if f == nil {
		return "field is null"
	}
	v := ""
	if f.opts.Format != "" {
		v = fmt.Sprintf(f.opts.Format, value)
	} else {
		v = fmt.Sprint(value)
	}
	return process.AddExtraPadding(v)
}
