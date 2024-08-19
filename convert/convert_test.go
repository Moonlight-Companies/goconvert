package convert

import (
	"reflect"
	"testing"
)

func TestConvertInto(t *testing.T) {
	intValue := 42
	int8Value := int8(42)
	int16Value := int16(42)
	int32Value := int32(42)
	int64Value := int64(42)
	uintValue := uint(42)
	uint8Value := uint8(42)
	uint16Value := uint16(42)
	uint32Value := uint32(42)
	uint64Value := uint64(42)
	boolValue := true
	float32Value := float32(42.5)
	float64Value := 42.5
	stringValue := "42"

	testCases := []struct {
		name     string
		input    interface{}
		wantType reflect.Type
		want     interface{}
		wantOk   bool
	}{
		// Int conversions
		{"int to int", intValue, reflect.TypeOf(int(0)), int(42), true},
		{"int to int8", intValue, reflect.TypeOf(int8(0)), int8(42), true},
		{"int to int16", intValue, reflect.TypeOf(int16(0)), int16(42), true},
		{"int to int32", intValue, reflect.TypeOf(int32(0)), int32(42), true},
		{"int to int64", intValue, reflect.TypeOf(int64(0)), int64(42), true},
		{"int to uint", intValue, reflect.TypeOf(uint(0)), uint(42), true},
		{"int to uint8", intValue, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"int to uint16", intValue, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"int to uint32", intValue, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"int to uint64", intValue, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"int to bool", intValue, reflect.TypeOf(bool(false)), true, true},
		{"int to float32", intValue, reflect.TypeOf(float32(0)), float32(42), true},
		{"int to float64", intValue, reflect.TypeOf(float64(0)), float64(42), true},
		{"int to string", intValue, reflect.TypeOf(""), "42", true},

		// Int8 conversions
		{"int8 to int", int8Value, reflect.TypeOf(int(0)), int(42), true},
		{"int8 to int8", int8Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"int8 to int16", int8Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"int8 to int32", int8Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"int8 to int64", int8Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"int8 to uint", int8Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"int8 to uint8", int8Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"int8 to uint16", int8Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"int8 to uint32", int8Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"int8 to uint64", int8Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"int8 to bool", int8Value, reflect.TypeOf(bool(false)), true, true},
		{"int8 to float32", int8Value, reflect.TypeOf(float32(0)), float32(42), true},
		{"int8 to float64", int8Value, reflect.TypeOf(float64(0)), float64(42), true},
		{"int8 to string", int8Value, reflect.TypeOf(""), "42", true},

		// Int16 conversions
		{"int16 to int", int16Value, reflect.TypeOf(int(0)), int(42), true},
		{"int16 to int8", int16Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"int16 to int16", int16Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"int16 to int32", int16Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"int16 to int64", int16Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"int16 to uint", int16Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"int16 to uint8", int16Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"int16 to uint16", int16Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"int16 to uint32", int16Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"int16 to uint64", int16Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"int16 to bool", int16Value, reflect.TypeOf(bool(false)), true, true},
		{"int16 to float32", int16Value, reflect.TypeOf(float32(0)), float32(42), true},
		{"int16 to float64", int16Value, reflect.TypeOf(float64(0)), float64(42), true},
		{"int16 to string", int16Value, reflect.TypeOf(""), "42", true},

		// Int32 conversions
		{"int32 to int", int32Value, reflect.TypeOf(int(0)), int(42), true},
		{"int32 to int8", int32Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"int32 to int16", int32Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"int32 to int32", int32Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"int32 to int64", int32Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"int32 to uint", int32Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"int32 to uint8", int32Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"int32 to uint16", int32Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"int32 to uint32", int32Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"int32 to uint64", int32Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"int32 to bool", int32Value, reflect.TypeOf(bool(false)), true, true},
		{"int32 to float32", int32Value, reflect.TypeOf(float32(0)), float32(42), true},
		{"int32 to float64", int32Value, reflect.TypeOf(float64(0)), float64(42), true},
		{"int32 to string", int32Value, reflect.TypeOf(""), "42", true},

		// Int64 conversions
		{"int64 to int", int64Value, reflect.TypeOf(int(0)), int(42), true},
		{"int64 to int8", int64Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"int64 to int16", int64Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"int64 to int32", int64Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"int64 to int64", int64Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"int64 to uint", int64Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"int64 to uint8", int64Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"int64 to uint16", int64Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"int64 to uint32", int64Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"int64 to uint64", int64Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"int64 to bool", int64Value, reflect.TypeOf(bool(false)), true, true},
		{"int64 to float32", int64Value, reflect.TypeOf(float32(0)), float32(42), true},
		{"int64 to float64", int64Value, reflect.TypeOf(float64(0)), float64(42), true},
		{"int64 to string", int64Value, reflect.TypeOf(""), "42", true},

		// Uint conversions
		{"uint to int", uintValue, reflect.TypeOf(int(0)), int(42), true},
		{"uint to int8", uintValue, reflect.TypeOf(int8(0)), int8(42), true},
		{"uint to int16", uintValue, reflect.TypeOf(int16(0)), int16(42), true},
		{"uint to int32", uintValue, reflect.TypeOf(int32(0)), int32(42), true},
		{"uint to int64", uintValue, reflect.TypeOf(int64(0)), int64(42), true},
		{"uint to uint", uintValue, reflect.TypeOf(uint(0)), uint(42), true},
		{"uint to uint8", uintValue, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"uint to uint16", uintValue, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"uint to uint32", uintValue, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"uint to uint64", uintValue, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"uint to bool", uintValue, reflect.TypeOf(bool(false)), true, true},
		{"uint to float32", uintValue, reflect.TypeOf(float32(0)), float32(42), true},
		{"uint to float64", uintValue, reflect.TypeOf(float64(0)), float64(42), true},
		{"uint to string", uintValue, reflect.TypeOf(""), "42", true},

		// Uint8 conversions
		{"uint8 to int", uint8Value, reflect.TypeOf(int(0)), int(42), true},
		{"uint8 to int8", uint8Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"uint8 to int16", uint8Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"uint8 to int32", uint8Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"uint8 to int64", uint8Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"uint8 to uint", uint8Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"uint8 to uint8", uint8Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"uint8 to uint16", uint8Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"uint8 to uint32", uint8Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"uint8 to uint64", uint8Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"uint8 to bool", uint8Value, reflect.TypeOf(bool(false)), true, true},
		{"uint8 to float32", uint8Value, reflect.TypeOf(float32(0)), float32(42), true},
		{"uint8 to float64", uint8Value, reflect.TypeOf(float64(0)), float64(42), true},
		{"uint8 to string", uint8Value, reflect.TypeOf(""), "42", true},

		// Uint16 conversions
		{"uint16 to int", uint16Value, reflect.TypeOf(int(0)), int(42), true},
		{"uint16 to int8", uint16Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"uint16 to int16", uint16Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"uint16 to int32", uint16Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"uint16 to int64", uint16Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"uint16 to uint", uint16Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"uint16 to uint8", uint16Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"uint16 to uint16", uint16Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"uint16 to uint32", uint16Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"uint16 to uint64", uint16Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"uint16 to bool", uint16Value, reflect.TypeOf(bool(false)), true, true},
		{"uint16 to float32", uint16Value, reflect.TypeOf(float32(0)), float32(42), true},
		{"uint16 to float64", uint16Value, reflect.TypeOf(float64(0)), float64(42), true},
		{"uint16 to string", uint16Value, reflect.TypeOf(""), "42", true},

		// Uint32 conversions
		{"uint32 to int", uint32Value, reflect.TypeOf(int(0)), int(42), true},
		{"uint32 to int8", uint32Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"uint32 to int16", uint32Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"uint32 to int32", uint32Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"uint32 to int64", uint32Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"uint32 to uint", uint32Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"uint32 to uint8", uint32Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"uint32 to uint16", uint32Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"uint32 to uint32", uint32Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"uint32 to uint64", uint32Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"uint32 to bool", uint32Value, reflect.TypeOf(bool(false)), true, true},
		{"uint32 to float32", uint32Value, reflect.TypeOf(float32(0)), float32(42), true},
		{"uint32 to float64", uint32Value, reflect.TypeOf(float64(0)), float64(42), true},
		{"uint32 to string", uint32Value, reflect.TypeOf(""), "42", true},

		// Uint64 conversions
		{"uint64 to int", uint64Value, reflect.TypeOf(int(0)), int(42), true},
		{"uint64 to int8", uint64Value, reflect.TypeOf(int8(0)), int8(42), true},
		// Uint64 conversions (continued)
		{"uint64 to int16", uint64Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"uint64 to int32", uint64Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"uint64 to int64", uint64Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"uint64 to uint", uint64Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"uint64 to uint8", uint64Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"uint64 to uint16", uint64Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"uint64 to uint32", uint64Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"uint64 to uint64", uint64Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"uint64 to bool", uint64Value, reflect.TypeOf(bool(false)), true, true},
		{"uint64 to float32", uint64Value, reflect.TypeOf(float32(0)), float32(42), true},
		{"uint64 to float64", uint64Value, reflect.TypeOf(float64(0)), float64(42), true},
		{"uint64 to string", uint64Value, reflect.TypeOf(""), "42", true},

		// Bool conversions
		{"bool to int", boolValue, reflect.TypeOf(int(0)), 1, true},
		{"bool to int8", boolValue, reflect.TypeOf(int8(0)), int8(1), true},
		{"bool to int16", boolValue, reflect.TypeOf(int16(0)), int16(1), true},
		{"bool to int32", boolValue, reflect.TypeOf(int32(0)), int32(1), true},
		{"bool to int64", boolValue, reflect.TypeOf(int64(0)), int64(1), true},
		{"bool to uint", boolValue, reflect.TypeOf(uint(0)), uint(1), true},
		{"bool to uint8", boolValue, reflect.TypeOf(uint8(0)), uint8(1), true},
		{"bool to uint16", boolValue, reflect.TypeOf(uint16(0)), uint16(1), true},
		{"bool to uint32", boolValue, reflect.TypeOf(uint32(0)), uint32(1), true},
		{"bool to uint64", boolValue, reflect.TypeOf(uint64(0)), uint64(1), true},
		{"bool to bool", boolValue, reflect.TypeOf(bool(false)), true, true},
		{"bool to float32", boolValue, reflect.TypeOf(float32(0)), float32(1), true},
		{"bool to float64", boolValue, reflect.TypeOf(float64(0)), float64(1), true},
		{"bool to string", boolValue, reflect.TypeOf(""), "1", true},

		// Float32 conversions
		{"float32 to int", float32Value, reflect.TypeOf(int(0)), 42, true},
		{"float32 to int8", float32Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"float32 to int16", float32Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"float32 to int32", float32Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"float32 to int64", float32Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"float32 to uint", float32Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"float32 to uint8", float32Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"float32 to uint16", float32Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"float32 to uint32", float32Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"float32 to uint64", float32Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"float32 to bool", float32Value, reflect.TypeOf(bool(false)), true, true},
		{"float32 to float32", float32Value, reflect.TypeOf(float32(0)), float32(42.5), true},
		{"float32 to float64", float32Value, reflect.TypeOf(float64(0)), float64(42.5), true},
		{"float32 to string", float32Value, reflect.TypeOf(""), "42.500000", true},

		// Float64 conversions
		{"float64 to int", float64Value, reflect.TypeOf(int(0)), 42, true},
		{"float64 to int8", float64Value, reflect.TypeOf(int8(0)), int8(42), true},
		{"float64 to int16", float64Value, reflect.TypeOf(int16(0)), int16(42), true},
		{"float64 to int32", float64Value, reflect.TypeOf(int32(0)), int32(42), true},
		{"float64 to int64", float64Value, reflect.TypeOf(int64(0)), int64(42), true},
		{"float64 to uint", float64Value, reflect.TypeOf(uint(0)), uint(42), true},
		{"float64 to uint8", float64Value, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"float64 to uint16", float64Value, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"float64 to uint32", float64Value, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"float64 to uint64", float64Value, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"float64 to bool", float64Value, reflect.TypeOf(bool(false)), true, true},
		{"float64 to float32", float64Value, reflect.TypeOf(float32(0)), float32(42.5), true},
		{"float64 to float64", float64Value, reflect.TypeOf(float64(0)), 42.5, true},
		{"float64 to string", float64Value, reflect.TypeOf(""), "42.500000", true},

		// String conversions
		{"string to int", stringValue, reflect.TypeOf(int(0)), 42, true},
		{"string to int8", stringValue, reflect.TypeOf(int8(0)), int8(42), true},
		{"string to int16", stringValue, reflect.TypeOf(int16(0)), int16(42), true},
		{"string to int32", stringValue, reflect.TypeOf(int32(0)), int32(42), true},
		{"string to int64", stringValue, reflect.TypeOf(int64(0)), int64(42), true},
		{"string to uint", stringValue, reflect.TypeOf(uint(0)), uint(42), true},
		{"string to uint8", stringValue, reflect.TypeOf(uint8(0)), uint8(42), true},
		{"string to uint16", stringValue, reflect.TypeOf(uint16(0)), uint16(42), true},
		{"string to uint32", stringValue, reflect.TypeOf(uint32(0)), uint32(42), true},
		{"string to uint64", stringValue, reflect.TypeOf(uint64(0)), uint64(42), true},
		{"string to bool", "1", reflect.TypeOf(bool(false)), true, true},
		{"string to float32", stringValue, reflect.TypeOf(float32(0)), float32(42), true},
		{"string to float64", stringValue, reflect.TypeOf(float64(0)), float64(42), true},
		{"string to string", stringValue, reflect.TypeOf(""), "42", true},

		// Edge cases
		{"string(invalid) to int", "not a number", reflect.TypeOf(int(0)), 0, false},
		{"overflow int8", 1000, reflect.TypeOf(int8(0)), int8(-24), true}, // 1000 % 256 = 232, which is -24 in int8
		{"string to bool (true)", "true", reflect.TypeOf(bool(false)), true, true},
		{"string to bool (false)", "false", reflect.TypeOf(bool(false)), false, true},
		{"string to bool (1)", "1", reflect.TypeOf(bool(false)), true, true},
		{"string to bool (0)", "0", reflect.TypeOf(bool(false)), false, true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.want.(type) {
			case int:
				testConversion[int](t, tt.name, tt.input, tt.want.(int), tt.wantOk)
			case int8:
				testConversion[int8](t, tt.name, tt.input, tt.want.(int8), tt.wantOk)
			case int16:
				testConversion[int16](t, tt.name, tt.input, tt.want.(int16), tt.wantOk)
			case int32:
				testConversion[int32](t, tt.name, tt.input, tt.want.(int32), tt.wantOk)
			case int64:
				testConversion[int64](t, tt.name, tt.input, tt.want.(int64), tt.wantOk)
			case uint:
				testConversion[uint](t, tt.name, tt.input, tt.want.(uint), tt.wantOk)
			case uint8:
				testConversion[uint8](t, tt.name, tt.input, tt.want.(uint8), tt.wantOk)
			case uint16:
				testConversion[uint16](t, tt.name, tt.input, tt.want.(uint16), tt.wantOk)
			case uint32:
				testConversion[uint32](t, tt.name, tt.input, tt.want.(uint32), tt.wantOk)
			case uint64:
				testConversion[uint64](t, tt.name, tt.input, tt.want.(uint64), tt.wantOk)
			case float32:
				testConversion[float32](t, tt.name, tt.input, tt.want.(float32), tt.wantOk)
			case float64:
				testConversion[float64](t, tt.name, tt.input, tt.want.(float64), tt.wantOk)
			case bool:
				testConversion[bool](t, tt.name, tt.input, tt.want.(bool), tt.wantOk)
			case string:
				testConversion[string](t, tt.name, tt.input, tt.want.(string), tt.wantOk)
			default:
				t.Errorf("Unsupported type in test case: %T", tt.want)
			}
		})
	}
}

func testConversion[T any](t *testing.T, name string, input interface{}, want T, wantOk bool) {
	t.Helper()
	got, ok := ConvertInto[T](input)
	if ok != wantOk {
		t.Errorf("%s: ConvertInto() ok = %v, wantOk %v", name, ok, wantOk)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: ConvertInto() = %v, want %v", name, got, want)
	}
}
func TestConvertIntoUnsupportedType(t *testing.T) {
	type unsupportedType struct{}
	_, ok := ConvertInto[unsupportedType](42)
	if ok {
		t.Errorf("Expected conversion to unsupported type to fail")
	}
}
