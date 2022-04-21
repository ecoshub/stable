package stable

import (
	"reflect"
)

func structArrayToTable(s interface{}) (*STable, error) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	te := t.Elem()
	if te.Kind() == reflect.Ptr {
		te = te.Elem()
	}

	caption := t.Name()
	if caption != "" {
		caption = te.Name() + "(s)"
	}
	table := New(caption)

	for i := 0; i < v.Len(); i++ {
		elementValue := v.Index(i)
		if elementValue.Type().Kind() == reflect.Ptr {
			elementValue = reflect.Indirect(elementValue)
		}
		if i == 0 {
			fieldNames := getFieldNames(elementValue, elementValue.Type())
			table.AddFields(fieldNames...)
		}
		fieldValues := getFieldValues(elementValue, elementValue.Type())
		table.Row(fieldValues...)
	}

	return table, nil
}

func structToTable(s interface{}) (*STable, error) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
		t = v.Type()
	}

	caption := t.Name()
	table := New(caption)

	fieldNames := getFieldNames(v, t)
	table.AddFields(fieldNames...)

	fieldValues := getFieldValues(v, t)
	table.Row(fieldValues...)

	return table, nil
}

func getFieldNames(v reflect.Value, t reflect.Type) []string {
	fieldNames := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		tableTag := fieldType.Tag.Get("table")
		if tableTag == "" {
			tableTag = fieldType.Name
		}
		fieldNames[i] = tableTag
	}
	return fieldNames
}

func getFieldValues(v reflect.Value, t reflect.Type) []interface{} {
	fieldValues := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldValues[i] = fieldValue
	}
	return fieldValues
}
