package stable

type alignment string

const (
	// AlignmentLeft alignment type left
	AlignmentLeft alignment = "left"
	// AlignmentCenter alignment type center
	AlignmentCenter alignment = "center"
	// AlignmentRight alignment type right
	AlignmentRight alignment = "right"
	// DefaultValueAlignment default value alignment type
	DefaultValueAlignment alignment = AlignmentLeft
	// DefaultHeaderAlignment default header alignment type
	DefaultHeaderAlignment alignment = AlignmentCenter
)
