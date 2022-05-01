package stable

import "fmt"

// Scheme main table scheme struct
type Scheme struct {
	Caption         string
	BorderStyleName borderStyleName
	GeneralPadding  int
	FieldOptions    map[string]*Options
}

// InjectScheme inject a scheme to input to create a table with scheme
func InjectScheme(sc *Scheme, i interface{}) (*STable, error) {
	t, err := ToTable(i)
	if err != nil {
		return nil, err
	}

	t.caption = processCaption(sc.Caption)
	style, err := processBorderStyle(sc.BorderStyleName)
	if err != nil {
		return nil, err
	}
	t.borderStyle = style
	t.generalPadding = processPadding(sc.GeneralPadding)
	for name, opts := range sc.FieldOptions {
		t.GetFieldByName(name).SetOptions(opts)
	}
	return t, nil
}

// PrintWithScheme print table with given table scheme and input
func PrintWithScheme(sc *Scheme, i interface{}) error {
	t, err := InjectScheme(sc, i)
	if err != nil {
		return err
	}
	fmt.Println(t)
	return nil
}
