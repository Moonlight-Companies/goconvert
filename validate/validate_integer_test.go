package convert

import (
	"strings"
	"testing"
)

func TestValidateBasicInteger(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{"Valid positive integer", "123", "123", false},
		{"Valid negative integer", "-456", "-456", false},
		{"Zero", "0", "0", false},
		{"Invalid decimal point", "123.456", "123<0x2E>456", true},
		{"Invalid character", "12a34", "12<0x61>34", true},
		{"Invalid symbols", "12+34", "12<0x2B>34", true},
		{"Multiple minus signs", "--123", "-<0x2D>123", true},
		{"Minus sign not at start", "123-456", "123<0x2D>456", true},
		{"Long valid integer", strings.Repeat("9", 255), strings.Repeat("9", 255), false},
		{"Too long integer", strings.Repeat("9", 256), strings.Repeat("9", 255) + "...", true},
		{"Long invalid integer", strings.Repeat("9", 254) + "a", strings.Repeat("9", 254) + "<0x61>", true},
		{"Empty string", "", "", true},
		{"Only minus sign", "-", "-", true},
		{"Minus sign after digits", "123-", "123<0x2D>", true},
		{"Multiple minus signs at start", "---123", "-<0x2D><0x2D>123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateBasicInteger(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
			if (err != nil) != tt.hasError {
				t.Errorf("expected hasError %v, got %v", tt.hasError, err != nil)
			}
		})
	}
}
