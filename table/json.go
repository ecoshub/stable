package table

import (
	"encoding/json"
	"errors"
	"sort"
)

var (
	ErrNullJSON error = errors.New("null json")
)

func jsonSwitch(j []byte) (*STable, error) {
	if len(j) == 0 {
		return nil, ErrNullJSON
	}
	if j[0] == '[' {
		return encodeJSONArray(j)
	}
	return encodeJSON(j)
}

func encodeJSON(j []byte) (*STable, error) {
	var jj map[string]interface{}
	err := json.Unmarshal(j, &jj)
	if err != nil {
		return nil, err
	}
	return encodeMap(jj)
}

func encodeJSONArray(j []byte) (*STable, error) {
	var jj []map[string]interface{}
	err := json.Unmarshal(j, &jj)
	if err != nil {
		return nil, err
	}
	if len(jj) == 0 {
		return nil, ErrNullJSON
	}
	return encodeMapArray(jj)
}

func encodeMap(m map[string]interface{}) (*STable, error) {
	table := New("json")
	table.AddFields("key", "value")

	fieldNames := make([]string, 0, len(m))
	for k := range m {
		fieldNames = append(fieldNames, k)
	}

	sort.Slice(fieldNames, func(i, j int) bool { return fieldNames[i] < fieldNames[j] })

	for i := 0; i < len(fieldNames); i++ {
		key := fieldNames[i]
		value := m[key]
		table.Row(key, value)
	}
	return table, nil
}

func encodeMapArray(m []map[string]interface{}) (*STable, error) {
	f := m[0]
	fieldNames := make([]string, 0, len(f))
	for k := range f {
		fieldNames = append(fieldNames, k)
	}
	sort.Slice(fieldNames, func(i, j int) bool { return fieldNames[i] < fieldNames[j] })

	table := New("json")
	table.AddFields(fieldNames...)
	for _, j := range m {
		values := make([]interface{}, 0, len(fieldNames))
		for i := 0; i < len(fieldNames); i++ {
			field := fieldNames[i]
			values = append(values, j[field])
		}
		table.Row(values...)
	}
	return table, nil
}
