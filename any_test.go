package any

import (
	"reflect"
	"testing"
)

type MyStruct struct {
	A int
	B string
}

func TestIsBasic(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		int(1),                   // int
		int8(2),                  // int8
		int16(3),                 // int16
		int32(4),                 // int32
		int64(5),                 // int64
		uint(6),                  // uint
		uint8(7),                 // uint8
		uint16(8),                // uint16
		uint32(9),                // uint32
		uint64(10),               // uint64
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsBasic(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is a basic\nActual:\t\t%t, %s is NOT a basic\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotBasic(t *testing.T) {
	typedValues := []any{
		MyStruct{},    // struct
		[]any{},       // slice
		[1]any{},      // array
		map[any]any{}, //map
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsBasic(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is a basic\nActual:\t\t%t, %s is NOT a basic\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsComparable(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		int(1),                   // int
		int8(2),                  // int8
		int16(3),                 // int16
		int32(4),                 // int32
		int64(5),                 // int64
		uint(6),                  // uint
		uint8(7),                 // uint8
		uint16(8),                // uint16
		uint32(9),                // uint32
		uint64(10),               // uint64
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
		MyStruct{},               // struct
		[1]any{"array"},          // array
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsComparable(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is comparable\nActual:\t\t%t, %s is NOT comparable\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotComparable(t *testing.T) {
	typedValues := []any{
		[]any{"slice"}, // slice
		map[any]any{},  // map
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsComparable(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not comparable\nActual:\t\t%t, %s IS comparable\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsOrdered(t *testing.T) {
	typedValues := []any{
		int(1),        // int
		int8(2),       // int8
		int16(3),      // int16
		int32(4),      // int32
		int64(5),      // int64
		uint(6),       // uint
		uint8(7),      // uint8
		uint16(8),     // uint16
		uint32(9),     // uint32
		uint64(10),    // uint64
		float32(11.1), // float32
		float64(12.2), // float64
		"hello",       // string
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsOrdered(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is ordered\nActual:\t\t%t, %s is NOT ordered\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotOrdered(t *testing.T) {
	typedValues := []any{
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		MyStruct{},               // struct
		[]any{},                  // slice
		[1]any{"array"},          // array
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsOrdered(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not ordered\nActual:\t\t%t, %s IS ordered\n", expected[i], valueType, actual[i], valueType)
		}
	}
}
