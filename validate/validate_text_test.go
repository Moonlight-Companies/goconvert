package validate

import (
	"strings"
	"testing"
)

func TestValidateBasicText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{"Valid text", "Hello, World!", "Hello, World!", false},
		{"Valid symbols", "a-b_c/d?e", "a-b_c/d?e", false},
		{"Invalid character", "Hello\nWorld", "Hello<0x0A>World", true},
		{"Long valid text", strings.Repeat("a", 255), strings.Repeat("a", 255), false},
		{"Too long text", strings.Repeat("a", 256), strings.Repeat("a", 255) + "...", true},
		{"Long with invalid", strings.Repeat("a", 254) + "\n", strings.Repeat("a", 254) + "<0x0A>", true},
		{"Empty string", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateBasicText(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
			if (err != nil) != tt.hasError {
				t.Errorf("expected hasError %v, got %v", tt.hasError, err != nil)
			}
		})
	}
}
