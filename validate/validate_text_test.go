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
		// Original test cases
		{"Valid text", "Hello, World!", "Hello, World!", false},
		{"Valid symbols", "a-b_c/d?e", "a-b_c/d?e", false},
		{"Invalid character", "Hello\nWorld", "Hello<0x0A>World", true},
		{"Long valid text", strings.Repeat("a", 255), strings.Repeat("a", 255), false},
		{"Too long text", strings.Repeat("a", 256), strings.Repeat("a", 255) + "...", true},
		{"Long with invalid", strings.Repeat("a", 254) + "\n", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa...", true},
		{"Empty string", "", "", false},

		// Unicode handling
		{"ASCII extended", "café", "caf<0xC3><0xA9>", true},
		{"CJK character", "漢", "<0xE6><0xBC><0xA2>", true},
		{"Emoji", "👍", "<0xF0><0x9F><0x91><0x8D>", true},
		{"Mixed valid and Unicode", "abc漢def", "abc<0xE6><0xBC><0xA2>def", true},

		// Control characters
		{"Tab character", "a\tb", "a<0x09>b", true},
		{"Carriage return", "a\rb", "a<0x0D>b", true},
		{"Null byte", "a\x00b", "a<0x00>b", true},
		{"Bell character", "a\ab", "a<0x07>b", true},
		{"Escape character", "a\x1bb", "a<0x1B>b", true},

		// Special cases
		{"All allowed symbols", "!@#$%^&*()-_/\\?", "!@#$%^&*()-_/\\?", false},
		{"Mixed spaces", "a b\tc\rd\ne", "a b<0x09>c<0x0D>d<0x0A>e", true},
		{"Binary data", "\x00\x01\x02\x03", "<0x00><0x01><0x02><0x03>", true},

		// Multiple replacements
		{"Multiple Unicode", "漢字漢字", "<0xE6><0xBC><0xA2><0xE5><0xAD><0x97><0xE6><0xBC><0xA2><0xE5><0xAD><0x97>", true},
		{"Mixed invalid chars", "a\nb\tc\rd", "a<0x0A>b<0x09>c<0x0D>d", true},

		// Edge cases with spaces and symbols
		{"Space at start", " abc", " abc", false},
		{"Space at end", "abc ", "abc ", false},
		{"Multiple spaces", "a  b   c", "a  b   c", false},
		{"Only spaces", "   ", "   ", false},
		{"Only symbols", "!@#$%", "!@#$%", false},

		// Complex combinations
		{"Mixed everything", "Hello,\nWorld!漢字@#$", "Hello,<0x0A>World!<0xE6><0xBC><0xA2><0xE5><0xAD><0x97>@#$", true},
		{"Alternating valid-invalid", "a\nb\tc漢d", "a<0x0A>b<0x09>c<0xE6><0xBC><0xA2>d", true},
		{"Multiple truncation", strings.Repeat("漢", 100), "<0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2><0xE6><0xBC><0xA2>...", true},

		// Zero-width and special Unicode
		{"Zero-width space", "a\u200bb", "a<0xE2><0x80><0x8B>b", true},
		{"RTL mark", "a\u200fb", "a<0xE2><0x80><0x8F>b", true},
		{"Combining diacritic", "e\u0301", "e<0xCC><0x81>", true},
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

func TestValidateBasicTextErrors(t *testing.T) {
	errorTests := []struct {
		name        string
		input       string
		expectedErr string
	}{
		{"Invalid chars only", "\n\t", "text contains invalid characters (replaced)"},
		{"Too long", strings.Repeat("a", 256), "text exceeds 255 characters (truncated)"},
		{"Both invalid and too long", strings.Repeat("a", 255) + "\n", "text exceeds 255 characters (truncated)"},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateBasicText(tt.input)
			if err == nil {
				t.Error("expected error, got nil")
				return
			}
			if err.Error() != tt.expectedErr {
				t.Errorf("expected error %q, got %q", tt.expectedErr, err.Error())
			}
		})
	}
}
