package utils

import (
	"errors"
	"reflect"
)

var (
	ErrFieldNotExists = errors.New("field does not exists")
)

// FieldExists check given named field exists in struct
// returns true on if field is exists
func FieldExists[T any](item T, fieldName string) bool {
	elems := reflect.TypeOf(&item).Elem()
	for i := 0; i < elems.NumField(); i++ {
		if elems.Field(i).Name == fieldName {
			return true
		}
	}
	return false
}

// ListByFieldName accepts array of struct, and field name as string
// It will return the array of field name, OR error if supplied field name does not exists
func ListByFieldName[T any](arr []T, fieldName string) (interface{}, error) {
	var x reflect.Value
	if len(arr) > 0 {
		firstItem := arr[0]
		if !FieldExists(firstItem, fieldName) {
			return x, ErrFieldNotExists
		}

		z := reflect.ValueOf(firstItem).FieldByName(fieldName)
		x = reflect.MakeSlice(reflect.SliceOf(z.Type()), len(arr), cap(arr))
	}
	for i, v := range arr {
		val := reflect.ValueOf(v).FieldByName(fieldName)
		x.Index(i).Set(val)
	}
	return x, nil

}
