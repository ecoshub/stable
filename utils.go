package stable

func addExtraPadding(s string, paddingAmount int) string {
	return repeat(" ", paddingAmount) + s + repeat(" ", paddingAmount)
}

func repeat(char string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += char
	}
	return s
}
