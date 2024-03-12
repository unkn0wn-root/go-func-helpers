package nullchecker

import (
	"reflect"
)

func IsNull[T any](value T) bool {
	// Use reflection to check if the value is a pointer
	if reflect.ValueOf(value).Kind() == reflect.Ptr {
		// Check if the pointer is nil
		return reflect.ValueOf(value).IsNil()
	}
	if !reflect.ValueOf(value).IsValid() {
		return true
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return reflect.ValueOf(value).Len() == 0
	}

	return false
}
