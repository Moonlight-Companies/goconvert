package validate

import (
	"strings"
	"testing"
)

func TestValidateBasicFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{"Valid integer", "123", "123", false},
		{"Valid negative integer", "-456", "-456", false},
		{"Valid float", "123.456", "123.456", false},
		{"Valid negative float", "-78.90", "-78.90", false},
		{"Valid float starting with dot", ".123", ".123", false},
		{"Valid float ending with dot", "123.", "123.", false},
		{"Invalid character", "12.3e4", "12.3<0x65>4", true},
		{"Multiple dots", "1.2.3", "1.2.3", true},
		{"Invalid symbols", "12+34", "12<0x2B>34", true},
		{"Long valid float", strings.Repeat("9", 255), strings.Repeat("9", 255), false},
		{"Too long float", strings.Repeat("9", 256), strings.Repeat("9", 255) + "...", true},
		{"Long invalid float", strings.Repeat("9", 254) + "a", strings.Repeat("9", 254) + "<0x61>", true},
		{"Empty string", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateBasicFloat(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
			if (err != nil) != tt.hasError {
				t.Errorf("expected hasError %v, got %v", tt.hasError, err != nil)
			}
		})
	}
}
