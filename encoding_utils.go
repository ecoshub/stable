package stable

import (
	"encoding/json"
	"reflect"
)

func isJSONEncoded(i interface{}) bool {
	_, ok := i.([]byte)
	if !ok {
		return false
	}
	if !json.Valid(i.([]byte)) {
		return false
	}
	return true
}

func isAStringInterfaceMap(t reflect.Type) bool {
	return t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.Interface
}

func isKindStruct(v reflect.Value) bool {
	if v.Type().Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}
	return v.Type().Kind() == reflect.Struct
}

func isElementKindStruct(t reflect.Type) bool {
	te := t.Elem()
	if te.Kind() == reflect.Ptr {
		te = te.Elem()
	}
	return te.Kind() == reflect.Struct
}
