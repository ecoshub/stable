package stable

import (
	"encoding/json"
	"errors"
	"reflect"
)

var (
	// ErrNotSupported error not supported type for convert to table
	ErrNotSupported error = errors.New("not supported type")
)

// ToTable coverts other data types to *STable type
func ToTable(i interface{}) (*STable, error) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		// check is struct element kind is struct
		if isElementKindStruct(t) {
			return structArrayToTable(i)
		}
		switch t.Elem().Kind() {
		//  byte array types
		case reflect.Uint8:
			b, ok := i.([]byte)
			if !ok {
				return nil, ErrNotSupported
			}
			if !json.Valid(i.([]byte)) {
				return nil, ErrNotSupported
			}
			return jsonSwitch(b)
		case reflect.Map:
			tt := t.Elem()
			if tt.Key().Kind() == reflect.String && tt.Elem().Kind() == reflect.Interface {
				return encodeMapArray(i.([]map[string]interface{}))
			}
		}
	case reflect.Map:
		if t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.Interface {
			return encodeMap(i.(map[string]interface{}))
		}
	case reflect.Ptr:
		if isKindStruct(v) {
			return structToTable(i)
		}
	case reflect.Struct:
		return structToTable(i)
	}
	return nil, ErrNotSupported
}
