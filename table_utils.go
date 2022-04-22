package stable

func (st *STable) isThereAnyChanges() bool {
	if st.changed {
		return true
	}
	for _, f := range st.fields {
		if f.changed {
			return true
		}
	}
	return false
}

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

func processCaption(caption string) string {
	if len(caption) > MaxCaptionLength {
		caption = caption[:MaxCaptionLength]
		caption += "..."
	}
	return caption
}

func processPadding(padding int) int {
	if padding > MaxGeneralPadding {
		return MaxGeneralPadding
	}
	return padding
}
