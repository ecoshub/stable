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
	// print format of field
	// standart print format abbreviations are valid
	// example: (value 12.34)
	//     Fortmat: "%0.2f (ms)"
	// output: 12.34 (ms)
	Format string

	// alignment of value ( AlignmentCLeft | AlignmentCenter | AlignmentRight )
	// example:
	//     Alignment: AlignmentCenter
	Alignment alignment

	// alignment of field header ( AlignmentCLeft | AlignmentCenter | AlignmentRight )
	// example:
	//     Alignment: AlignmentRight
	HeaderAlignment alignment

	// hide | show the field
	Hide bool

	// limit the char output of value
	// example: (value: "/var/log/sys/crontab.log")
	//     CharLimit: 12
	// output: " /var/log/sys..."
	CharLimit int

	// char limit orientation
	// example: (value: "/var/log/sys/crontab.log")
	//     CharLimit: 12
	//     LimitFromStart: true
	// output: ".../crontab.log"
	LimitFromStart bool
}

// creates a new field with name
func newField(name string) *Field {
	return &Field{name: name, opts: &Options{
		Alignment:       DefaultValueAlignment,
		HeaderAlignment: DefaultHeaderAlignment,
	},
	}
}

// NewFieldWithOptions create new field with options (alignment, char limit, format etc.)
func NewFieldWithOptions(name string, opts *Options) *Field {
	if opts == nil {
		return newField(name)
	}
	if opts.Alignment == "" {
		opts.Alignment = DefaultValueAlignment
	}
	if opts.HeaderAlignment == "" {
		opts.HeaderAlignment = DefaultHeaderAlignment
	}
	return &Field{
		name: name,
		opts: opts,
	}
}

// GetName get fields name
func (f *Field) GetName() string {
	if f == nil {
		return FieldIsNil
	}
	return f.name
}

// SetName set field a new name
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

// GetAlignment get value alignment ( left | center | right )
func (f *Field) GetAlignment() string {
	if f == nil {
		return FieldIsNil
	}
	if f.opts == nil {
		return string(DefaultValueAlignment)
	}
	return string(f.opts.Alignment)
}

// SetAlignment set a new alignment for field values.
func (f *Field) SetAlignment(alignment alignment) {
	if f == nil {
		fmt.Println(FieldIsNil)
		return
	}
	f.opts.Alignment = alignment
	f.changed = true
}

// AlignCenter align field values to center
func (f *Field) AlignCenter() {
	f.SetAlignment(AlignmentCenter)
}

// AlignLeft align field values to left
func (f *Field) AlignLeft() {
	f.SetAlignment(AlignmentLeft)
}

// AlignRight align field values to right
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

// ChangeVisibility change visibility of field ( hide | show )
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
