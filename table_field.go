// Package stable provides functionality for creating ASCII tables from various data sources
// including structs, arrays, JSON, maps, and CSV data.
package stable

import (
	"fmt"
	"strings"
)

// Field represents a column in the table with its properties and formatting options.
// Each field has a name, width, and various formatting options that control how
// its values are displayed.
type Field struct {
	name    string
	opts    *Options
	changed bool
}

// Options defines the formatting and display options for a field.
type Options struct {
	// Format specifies the printf-style format string for value formatting.
	// Example: For value 12.34 with Format "%0.2f (ms)", output will be "12.34 (ms)"
	Format string

	// Alignment controls the horizontal alignment of field values
	// (AlignmentLeft | AlignmentCenter | AlignmentRight)
	Alignment alignment

	// HeaderAlignment controls the horizontal alignment of the field header
	// (AlignmentLeft | AlignmentCenter | AlignmentRight)
	HeaderAlignment alignment

	// Hide determines whether the field should be hidden from display
	Hide bool

	// CharLimit specifies the maximum number of characters to display.
	// If set to 0, no limit is applied.
	CharLimit int

	// LimitFromStart determines whether to truncate from the start or end
	// when CharLimit is exceeded. If true, truncates from start (shows end),
	// if false, truncates from end (shows start).
	LimitFromStart bool

	// Summary configures summary calculations for the field
	Summary *SummaryOptions
}

// SummaryOptions defines the configuration for field summary calculations.
type SummaryOptions struct {
	Type   string  // Type of summary (sum, avg, min, max)
	Format string  // Format string for displaying the summary value
	Value  float64 // Calculated summary value
}

const (
	// NilValueString is displayed when a field value is nil
	NilValueString = "-"

	// FieldIsNil is returned when a field pointer is nil
	FieldIsNil = "field is nil."

	// Summary types
	SummarySum = "sum"
	SummaryAvg = "avg"
	SummaryMin = "min"
	SummaryMax = "max"
)

// newField creates a new field with the given name and default options.
func newField(name string) *Field {
	return &Field{
		name: name,
		opts: &Options{
			Alignment:       DefaultValueAlignment,
			HeaderAlignment: DefaultHeaderAlignment,
		},
	}
}

// NewFieldWithOptions creates a new field with the given name and options.
// If opts is nil, default options will be used.
func NewFieldWithOptions(name string, opts *Options) *Field {
	if opts == nil {
		return newField(name)
	}

	// Ensure alignment options are set
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

// GetName returns the field's name. Returns FieldIsNil if the field is nil.
func (f *Field) GetName() string {
	if f == nil {
		return FieldIsNil
	}
	return f.name
}

// SetName sets a new name for the field.
// No-op if the field is nil.
func (f *Field) SetName(name string) {
	if f == nil {
		return
	}
	f.name = name
	f.changed = true
}

// SetHeaderAlignment sets the alignment for the field's header.
// Prints FieldIsNil and returns if the field is nil.
func (f *Field) SetHeaderAlignment(alignment alignment) {
	if f == nil {
		fmt.Println(FieldIsNil)
		return
	}
	f.opts.HeaderAlignment = alignment
	f.changed = true
}

// GetAlignment returns the field's value alignment as a string.
// Returns FieldIsNil if the field is nil.
func (f *Field) GetAlignment() string {
	if f == nil {
		return FieldIsNil
	}
	if f.opts == nil {
		return string(DefaultValueAlignment)
	}
	return string(f.opts.Alignment)
}

// SetAlignment sets the alignment for the field's values.
// Prints FieldIsNil and returns if the field is nil.
func (f *Field) SetAlignment(alignment alignment) {
	if f == nil {
		fmt.Println(FieldIsNil)
		return
	}
	f.opts.Alignment = alignment
	f.changed = true
}

// AlignCenter sets the field's value alignment to center.
func (f *Field) AlignCenter() {
	f.SetAlignment(AlignmentCenter)
}

// AlignLeft sets the field's value alignment to left.
func (f *Field) AlignLeft() {
	f.SetAlignment(AlignmentLeft)
}

// AlignRight sets the field's value alignment to right.
func (f *Field) AlignRight() {
	f.SetAlignment(AlignmentRight)
}

// SetOptions sets all options for the field.
// Prints FieldIsNil and returns if the field is nil.
func (f *Field) SetOptions(opts *Options) {
	if f == nil {
		fmt.Println(FieldIsNil)
		return
	}

	// Ensure alignment options are set
	if opts.Alignment == "" {
		opts.Alignment = DefaultValueAlignment
	}
	if opts.HeaderAlignment == "" {
		opts.HeaderAlignment = DefaultHeaderAlignment
	}

	f.opts = opts
	f.changed = true
}

// IsHidden returns whether the field is hidden.
func (f *Field) IsHidden() bool {
	return f != nil && f.opts != nil && f.opts.Hide
}

// ChangeVisibility changes the visibility of the field.
// No-op if the field is nil.
func (f *Field) ChangeVisibility(hide bool) {
	if f == nil || f.opts == nil {
		return
	}
	f.opts.Hide = hide
	f.changed = true
}

// Show makes the field visible.
func (f *Field) Show() {
	f.ChangeVisibility(false)
}

// Hide makes the field hidden.
func (f *Field) Hide() {
	f.ChangeVisibility(true)
}

// toString formats a value according to the field's options and adds padding.
func (f *Field) toString(value interface{}, paddingAmount int) string {
	if value == nil {
		value = NilValueString
	}

	var v string
	if f.opts != nil && f.opts.Format != "" {
		v = fmt.Sprintf(f.opts.Format, value)
	} else {
		v = fmt.Sprint(value)
	}

	// Replace tabs with spaces
	v = strings.ReplaceAll(v, "\t", "    ")

	// Apply character limit if set
	if f.opts != nil && f.opts.CharLimit > 0 && f.opts.CharLimit < len(v) {
		if f.opts.LimitFromStart {
			v = "..." + v[len(v)-f.opts.CharLimit:]
		} else {
			v = v[:f.opts.CharLimit] + "..."
		}
	}

	return addExtraPadding(v, paddingAmount)
}

// SetSummary sets the summary configuration for the field.
// No-op if the field is nil.
func (f *Field) SetSummary(summaryType string, format string) {
	if f == nil {
		return
	}
	if f.opts == nil {
		f.opts = &Options{}
	}
	f.opts.Summary = &SummaryOptions{
		Type:   summaryType,
		Format: format,
	}
	f.changed = true
}

// HasSummary returns whether the field has a summary configured.
func (f *Field) HasSummary() bool {
	return f != nil && f.opts != nil && f.opts.Summary != nil
}
