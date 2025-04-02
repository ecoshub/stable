package stable

import (
	"reflect"
)

func isAStringInterfaceMap(t reflect.Type) bool {
	return t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.Interface
}

func isElementKindStruct(t reflect.Type) bool {
	te := t.Elem()
	if te.Kind() == reflect.Ptr {
		te = te.Elem()
	}
	return te.Kind() == reflect.Struct
}
