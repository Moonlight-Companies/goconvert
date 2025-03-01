package validate

import (
	"fmt"
	"testing"
)

func TestSanitize(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		// String tests
		{"Empty string", "", ""},
		{"Normal string", "Hello, world!", "Hello, world!"},
		{"Long string", createLongString(300), createLongString(255) + "..."},
		{"String with non-printable chars", "Hello\x00World", "Hello<0x00>World"},
		{"String with emoji", "Hello 😊 World", "Hello <0xF0><0x9F><0x98><0x8A> World"},

		// Integer tests
		{"Zero", 0, "0"},
		{"Positive integer", 42, "42"},
		{"Negative integer", -123, "-123"},
		{"Large integer", 9223372036854775807, "9223372036854775807"}, // int64 max
		{"Various int types", int8(127), "127"},
		{"Various int types", int16(-32768), "-32768"},
		{"Various int types", int32(2147483647), "2147483647"},
		{"Various int types", int64(-9223372036854775808), "-9223372036854775808"},

		// Unsigned integer tests
		{"Unsigned integer", uint(42), "42"},
		{"Unsigned integer types", uint8(255), "255"},
		{"Unsigned integer types", uint16(65535), "65535"},
		{"Unsigned integer types", uint32(4294967295), "4294967295"},
		{"Unsigned integer types", uint64(18446744073709551615), "18446744073709551615"}, // uint64 max

		// Float tests
		{"Float zero", 0.0, "0"},
		{"Positive float", 3.14159, "3.14159"},
		{"Negative float", -2.718, "-2.718"},
		{"Float with trailing zeros", 1.5000, "1.5"},
		{"Scientific notation", 1e-9, "1<0x65>-09"},
		{"Float32 type", float32(3.14), "3.14"},
		{"Float64 type", float64(-0.00123), "-0.00123"},

		// Boolean tests
		{"True", true, "true"},
		{"False", false, "false"},

		// Nil test
		{"Nil value", nil, "<nil>"},

		// Other types
		{"Struct", struct{ Name string }{"Test"}, "{Test}"},
		{"Slice", []int{1, 2, 3}, "[1 2 3]"},
		{"Map", map[string]int{"a": 1, "b": 2}, "map[a:1 b:2]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sanitize(tt.input)
			if got != tt.expected {
				t.Errorf("Sanitize(%v) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// Additional tests for edge cases and invalid inputs
func TestSanitizeInvalidInputs(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		validateFn  func(string) (string, error)
		expectError bool
	}{
		{"Invalid integer - empty", "", ValidateBasicInteger, true},
		{"Invalid integer - minus only", "-", ValidateBasicInteger, true},
		{"Invalid integer - with letters", "123a456", ValidateBasicInteger, true},
		{"Invalid integer - multiple minus", "--123", ValidateBasicInteger, true},
		{"Invalid integer - minus in middle", "123-456", ValidateBasicInteger, true},

		{"Invalid float - empty", "", ValidateBasicFloat, true},
		{"Invalid float - minus only", "-", ValidateBasicFloat, true},
		{"Invalid float - multiple decimals", "123.456.789", ValidateBasicFloat, true},
		{"Invalid float - no digits", ".", ValidateBasicFloat, true},
		{"Invalid float - with letters", "3.14e2", ValidateBasicFloat, true},
		{"Invalid float - multiple minus", "--1.23", ValidateBasicFloat, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.validateFn(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("Expected error: %v, got: %v", tt.expectError, err)
			}
		})
	}
}

// Test for truncation
func TestSanitizeTruncation(t *testing.T) {
	// Create input strings of different lengths
	tests := []struct {
		length   int
		expected bool
	}{
		{100, false},
		{255, false},
		{256, true},
		{1000, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Length_%d", tt.length), func(t *testing.T) {
			input := createLongString(tt.length)
			sanitized := Sanitize(input)

			// If truncation expected, check for "..."
			if tt.expected {
				if len(sanitized) != 258 { // 255 + "..."
					t.Errorf("Expected length 258, got %d", len(sanitized))
				}
				if sanitized[255:] != "..." {
					t.Errorf("Expected ... at the end, got %s", sanitized[255:])
				}
			} else {
				if len(sanitized) != tt.length {
					t.Errorf("Expected length %d, got %d", tt.length, len(sanitized))
				}
			}
		})
	}
}

// Helper function to create a string of a specific length
func createLongString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = 'a' + byte(i%26)
	}
	return string(result)
}
