package any

import (
	"reflect"
)

func IsBasic(i any) bool {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		return true
	default:
		return false
	}
}

func IsComparable(i any) bool {
	return reflect.ValueOf(i).Comparable()
}
