package process

const (
	// ExtraPaddingSpaces sxtra space padding start and end of values
	ExtraPaddingSpaces int = 1
)

func AddExtraPadding(s string) string {
	return repeat(" ", ExtraPaddingSpaces) + s + repeat(" ", ExtraPaddingSpaces)
}
func repeat(char string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += char
	}
	return s
}
