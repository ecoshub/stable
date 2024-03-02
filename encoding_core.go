package stable

import (
	"encoding/json"
	"errors"
	"reflect"
)

var (
	// ErrNotSupported error not supported type for convert to table
	ErrNotSupported error = errors.New("stable error. Type not supported")
)

// ToTable coverts other data types to *STable type
func ToTable(i interface{}) (*STable, error) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	// Indirect/redirect pointer to resolve its underlying type
	if t.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
		t = v.Type()
	}

	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		t = t.Elem()
		switch t.Kind() {
		case reflect.Ptr:
			if isElementKindStruct(t) {
				return structArrayToTable(i)
			}
		case reflect.Struct:
			return structArrayToTable(i)
		case reflect.Uint8:
			data := i.([]byte)
			if json.Valid(data) {
				return jsonSwitch(data)
			}
		case reflect.Map:
			if isAStringInterfaceMap(t) {
				return encodeMapArray(i.([]map[string]interface{}))
			}
		}
	case reflect.String:
		data := []byte(i.(string))
		if json.Valid(data) {
			return jsonSwitch(data)
		}
	case reflect.Map:
		if isAStringInterfaceMap(t) {
			return encodeMap(i.(map[string]interface{}))
		}
	case reflect.Struct:
		return structToTable(i)
	}
	return nil, ErrNotSupported
}
