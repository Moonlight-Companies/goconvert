package convert

import (
	"fmt"
	"reflect"
	"testing"
	"time"
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

func TestConvertIntoTimeTime(t *testing.T) {
	// Reference times for testing
	referenceTime := time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC)
	unixTime := time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC).Unix()
	unixTimeFloat := float64(unixTime) + 0.5 // with fractional seconds

	testCases := []struct {
		name   string
		input  interface{}
		want   time.Time
		wantOk bool
	}{
		// Direct time.Time conversion
		{"time.Time to time.Time", referenceTime, referenceTime, true},

		// String format conversions
		{"ISO8601 to time.Time", "2023-05-15T14:30:45Z", referenceTime, true},
		{"YYYY-MM-DD to time.Time", "2023-05-15", time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), true},
		{"MM/DD/YYYY to time.Time", "05/15/2023", time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), true},
		{"MMDDYYYY to time.Time", "05152023", time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), true},
		{"RFC3339 to time.Time", "2023-05-15T14:30:45Z", referenceTime, true},
		{"RFC1123 to time.Time", "Mon, 15 May 2023 14:30:45 UTC", referenceTime, true},

		// Integer timestamp conversions
		{"int to time.Time", int(unixTime), time.Unix(unixTime, 0), true},
		{"int64 to time.Time", int64(unixTime), time.Unix(unixTime, 0), true},
		{"string Unix timestamp to time.Time", fmt.Sprintf("%d", unixTime), time.Unix(unixTime, 0), true},

		// Float timestamp conversions (with fractional seconds)
		{"float64 to time.Time", unixTimeFloat, time.Unix(unixTime, int64(0.5*1e9)), true},
		{"float32 to time.Time", float32(unixTimeFloat), time.Unix(unixTime, int64(0.5*1e9)), false},
		{"string float timestamp to time.Time", fmt.Sprintf("%f", unixTimeFloat), time.Unix(unixTime, int64(0.5*1e9)), true},

		// Edge cases
		{"invalid string to time.Time", "not a date", time.Time{}, false},
		{"empty string to time.Time", "", time.Time{}, false},
		{"nil to time.Time", nil, time.Time{}, false},
		{"bool to time.Time", true, time.Time{}, false},

		// Additional formats
		{"DD-Mon-YYYY to time.Time", "15-May-2023", time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), true},
		{"Month D, YYYY to time.Time", "May 15, 2023", time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), true},
		{"YYYY/MM/DD to time.Time", "2023/05/15", time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), true},

		// JavaScript-style millisecond timestamp
		{"JavaScript millisecond timestamp string", fmt.Sprintf("%d", unixTime*1000), time.Unix(unixTime, 0), true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ConvertInto[time.Time](tt.input)

			if ok != tt.wantOk {
				t.Errorf("ConvertInto() ok = %v, wantOk %v", ok, tt.wantOk)
				return
			}

			if tt.wantOk {
				// For successful conversions, check if times are equal within a reasonable tolerance
				// This is important especially for floating point timestamp conversions
				if !timesApproximatelyEqual(got, tt.want) {
					t.Errorf("ConvertInto() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// Helper function to compare times with a small tolerance for floating point imprecision
func timesApproximatelyEqual(a, b time.Time) bool {
	diff := a.Sub(b)
	if diff < 0 {
		diff = -diff
	}
	return diff < 10*time.Millisecond // Allow 10ms of tolerance
}

// Test for common date formats used in HTTP headers
func TestConvertIntoTimeHTTPFormats(t *testing.T) {
	httpDateFormats := []struct {
		format string
		date   string
		want   time.Time
	}{
		// HTTP date formats
		{
			"RFC1123",
			"Mon, 15 May 2023 14:30:45 GMT",
			time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC),
		},
		{
			"RFC850",
			"Monday, 15-May-23 14:30:45 GMT",
			time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC),
		},
		{
			"ANSI C",
			"Mon May 15 14:30:45 2023",
			time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC),
		},
		{
			"Custom",
			"May 15 2023 2:30 PM",
			time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC),
		},
	}

	for _, tt := range httpDateFormats {
		t.Run(tt.format, func(t *testing.T) {
			got, ok := ConvertInto[time.Time](tt.date)
			if !ok {
				t.Errorf("ConvertInto() failed to parse %s", tt.date)
				return
			}

			if !timesApproximatelyEqual(got, tt.want) {
				t.Errorf("ConvertInto() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test for handling of edge cases with timezone information
func TestConvertIntoTimeWithTimezones(t *testing.T) {
	timezone := []struct {
		name  string
		input string
		want  time.Time
	}{
		{
			"EST timezone",
			"2023-05-15T09:30:45-05:00",
			time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC), // This is the UTC equivalent
		},
		{
			"JST timezone",
			"2023-05-15T23:30:45+09:00",
			time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC), // This is the UTC equivalent
		},
		{
			"Z timezone explicit",
			"2023-05-15T14:30:45Z",
			time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC),
		},
	}

	for _, tt := range timezone {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ConvertInto[time.Time](tt.input)
			if !ok {
				t.Errorf("ConvertInto() failed to parse %s", tt.input)
				return
			}

			// Convert to UTC for comparison if needed
			gotUTC := got.UTC()

			if !timesApproximatelyEqual(gotUTC, tt.want) {
				t.Errorf("ConvertInto() = %v, want %v", gotUTC, tt.want)
			}
		})
	}
}

// Test for JavaScript-style millisecond timestamps
func TestConvertIntoTimeJavaScriptTimestamps(t *testing.T) {
	unixTime := time.Date(2023, 5, 15, 14, 30, 45, 0, time.UTC).Unix()
	jsTimestamp := unixTime * 1000 // JavaScript uses milliseconds

	testCases := []struct {
		name  string
		input interface{}
		want  time.Time
	}{
		{
			"int64 JavaScript timestamp",
			int64(jsTimestamp),
			time.Unix(unixTime, 0),
		},
		{
			"float64 JavaScript timestamp",
			float64(jsTimestamp),
			time.Unix(unixTime, 0),
		},
		{
			"string JavaScript timestamp",
			fmt.Sprintf("%d", jsTimestamp),
			time.Unix(unixTime, 0),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// We need to implement JavaScript timestamp detection in our converter
			// This test may initially fail until we add that functionality
			got, ok := ConvertInto[time.Time](tt.input)
			if !ok {
				t.Errorf("ConvertInto() failed to parse JavaScript timestamp: %v", tt.input)
				return
			}

			if !timesApproximatelyEqual(got, tt.want) {
				t.Errorf("ConvertInto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolConversion(t *testing.T) {
	testCases := []struct {
		name     string
		input    bool
		expected bool
	}{
		{"true to bool", true, true},    // expect true -> true
		{"false to bool", false, false}, // expect false -> false
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, ok := ConvertInto[bool](tc.input)

			if !ok {
				t.Errorf("ConvertInto[bool](%v) returned ok=false, expected ok=true", tc.input)
			}

			if result != tc.expected {
				t.Errorf("ConvertInto[bool](%v) returned %v, expected %v",
					tc.input, result, tc.expected)
			}
		})
	}

	// Also test direct tracing of the code path to confirm behavior
	t.Run("direct logic trace", func(t *testing.T) {
		// Testing true input
		trueInput := true
		trueResult, trueOk := ConvertInto[bool](trueInput)
		t.Logf("Input: true, Result: %v, Ok: %v", trueResult, trueOk)

		// Testing false input
		falseInput := false
		falseResult, falseOk := ConvertInto[bool](falseInput)
		t.Logf("Input: false, Result: %v, Ok: %v", falseResult, falseOk)
	})
}

// TestConvertIntoTime_JulianDays tests the special time.Time conversion for julian day strings.
// Negative values (from -365 to -1) are interpreted as today minus N days.
// Zero is interpreted as December 31 of the previous year.
// Positive values (1 to 365) are interpreted as the Nth day of the current year.
func TestConvertIntoTime_JulianDays(t *testing.T) {
	// Capture the current year so that expected values remain relative.
	currentYear := time.Now().Year()

	// Define test cases.
	// We use a function for the expected value so that for negative offsets we can capture a reference "now".
	tests := []struct {
		name     string
		input    string
		expected func(ref time.Time) time.Time
	}{
		{
			name:  "Negative 7 days: today minus 7",
			input: "-7",
			expected: func(ref time.Time) time.Time {
				// Note: since ConvertInto calls time.Now(), we compare only year, month, day.
				return ref.AddDate(0, 0, -7)
			},
		},
		{
			name:  "Zero: Dec 31 of previous year",
			input: "0",
			expected: func(_ time.Time) time.Time {
				return time.Date(currentYear-1, time.December, 31, 0, 0, 0, 0, time.UTC)
			},
		},
		{
			name:  "Positive 1: Jan 1 of current year",
			input: "1",
			expected: func(_ time.Time) time.Time {
				return time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC)
			},
		},
		{
			name:  "Positive 365: Dec 31 of current year",
			input: "365",
			expected: func(_ time.Time) time.Time {
				// Day 1 is Jan 1, so day 365 is Jan 1 + 364 days.
				return time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 364)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Capture a reference time that we can use for comparison.
			refNow := time.Now()

			got, ok := ConvertInto[time.Time](tc.input)
			if !ok {
				t.Fatalf("Conversion failed for input %s", tc.input)
			}

			exp := tc.expected(refNow)
			// Since the conversion for negative values depends on time.Now(),
			// compare only the date components (year, month, day).
			if got.Year() != exp.Year() || got.Month() != exp.Month() || got.Day() != exp.Day() {
				t.Errorf("For input %q, expected date %s, got %s",
					tc.input,
					exp.Format("2006-01-02"),
					got.Format("2006-01-02"))
			} else {
				t.Logf("Success: input %q yielded %s", tc.input, got.Format("2006-01-02"))
			}
		})
	}
}

// TestConvertIntoTime_24hFormats tests the 24-hour formats with both non-zero-padded and zero-padded day values.
func TestConvertIntoTime_24hFormats(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  time.Time
	}{
		{
			name:  "24h format (non zero-padded day)",
			input: "Jan 2 2006 13:04",
			want:  time.Date(2006, time.January, 2, 13, 4, 0, 0, time.UTC),
		},
		{
			name:  "24h format (zero padded day)",
			input: "Jan 02 2006 13:04",
			want:  time.Date(2006, time.January, 2, 13, 4, 0, 0, time.UTC),
		},
		{
			name:  "24h format (non zero-padded day)",
			input: "Jan 2 2006 13:04:05",
			want:  time.Date(2006, time.January, 2, 13, 4, 5, 0, time.UTC),
		},
		{
			name:  "24h format (zero padded day)",
			input: "Jan 02 2006 13:04:05",
			want:  time.Date(2006, time.January, 2, 13, 4, 5, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := ConvertInto[time.Time](tc.input)
			if !ok {
				t.Errorf("ConvertInto() failed to parse input %q", tc.input)
				return
			}
			if !got.Equal(tc.want) {
				t.Errorf("ConvertInto() = %v, want %v", got, tc.want)
			}
		})
	}
}
