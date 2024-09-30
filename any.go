package any

import (
	"reflect"
)

// Return true if the value given is a basic type. Else, false.
func IsBasic(v any) bool {
	switch reflect.TypeOf(v).Kind() {
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

// Return true if the value given is comparable. Else, false.
func IsComparable(v any) bool {
	return reflect.ValueOf(v).Comparable()
}

// Return true if the value given adheres to cmp.Ordered. Else, false.
func IsOrdered(v any) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.String:
		return true
	default:
		return false
	}
}

// Return true if the value given is a map. Else, false.
func IsMap(v any) bool {
	return isOfReflectKind(v, reflect.Map)
}

// Return true if the value given is a slice. Else, false.
func IsSlice(v any) bool {
	return isOfReflectKind(v, reflect.Slice)
}

// Return true if the value given is an array. Else, false.
func IsArray(v any) bool {
	return isOfReflectKind(v, reflect.Array)
}

// Return true if the value given is a struct. Else, false.
func IsStruct(v any) bool {
	return isOfReflectKind(v, reflect.Struct)
}

// Return true if the value given has a reflect type kind that matches the kind given. Else, false.
func isOfReflectKind(v any, kind reflect.Kind) bool {
	return reflect.TypeOf(v).Kind() == kind
}
