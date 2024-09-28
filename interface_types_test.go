package interface_types

import (
	"testing"
)

func TestIsPrimitive(t *testing.T) {
	interfaces := []interface{}{
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
	for _, i := range interfaces {
		expected = append(expected, true)
		actual = append(actual, IsPrimitive(i))
	}

	// Test actual values against expected
	for i := 0; i < len(interfaces); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("\nExpected:\t%t, value is a primitive\nActual:\t\t%t, value is NOT a primitive\n", expected[i], actual[i])
		}
	}
}

type MyStruct struct {
	A int
	B string
}

func TestIsNotPrimitive(t *testing.T) {
	myStruct := MyStruct{}
	expected := false
	actual := IsPrimitive(myStruct)

	if actual != expected {
		t.Fatalf("\nExpected:\t%t, value is not a primitive\nActual:\t\t%t, value IS a primitive\n", expected, actual)
	}
}
