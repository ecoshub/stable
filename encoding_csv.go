package stable

import (
	"encoding/csv"
	"errors"
	"io"
	"strings"
)

var (
	// ErrMalformedCSV malformed csv
	ErrMalformedCSV error = errors.New("'stable' error. malformed csv")
)

// CSVToTable convert csv encoded string to *STable
func CSVToTable(data string) (*STable, error) {
	if len(data) == 0 {
		return nil, ErrMalformedCSV
	}
	r := csv.NewReader(strings.NewReader(data))
	lines := make([][]string, 0, 16)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lines = append(lines, record)
	}
	if len(lines) < 1 {
		return nil, ErrMalformedCSV
	}
	t := New("csv")
	t.AddFields(lines[0]...)
	for _, r := range lines[1:] {
		inter := make([]interface{}, len(r))
		for i, v := range r {
			inter[i] = v
		}
		t.Row(inter...)
	}
	return t, nil
}
