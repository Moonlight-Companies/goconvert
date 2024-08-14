package validate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ValidateBasicFloat(txt string) (string, error) {
	var sanitized strings.Builder
	var invalidChars bool

	// Regex for validating floating-point numbers:
	// ^             : Start of the string
	// -?            : Optional minus sign (zero or one occurrence)
	// (             : Start of first capturing group
	//   \d+         : One or more digits
	//   (\.\d*)?    : Optional decimal point followed by zero or more digits
	// |             : OR
	//   \.\d+       : Decimal point followed by one or more digits
	// )             : End of first capturing group
	// $             : End of the string
	//
	// This regex allows for:
	// - Optional leading minus sign for negative numbers
	// - Integers (e.g., "123", "-456")
	// - Floating-point numbers with optional decimal places (e.g., "123.456", "-78.90")
	// - Numbers starting with a decimal point (e.g., ".123")
	// - Numbers ending with a decimal point (e.g., "123.")
	// It does not allow:
	// - Multiple decimal points
	// - Scientific notation (e.g., "1e10")
	// - Empty strings or strings with only a minus sign
	regex := regexp.MustCompile(`^-?(\d+(\.\d*)?|\.\d+)$`)

	// Always process character by character to ensure proper truncation
	for i, r := range txt {
		if i >= 255 {
			sanitized.WriteString("...")
			break
		}

		if (r >= '0' && r <= '9') || r == '.' || r == '-' {
			sanitized.WriteRune(r)
		} else {
			sanitized.WriteString(fmt.Sprintf("<0x%02X>", r))
			invalidChars = true
		}
	}

	result := sanitized.String()

	var err error
	if len(txt) > 255 {
		err = errors.New("text exceeds 255 characters (truncated)")
	}

	if invalidChars {
		if err != nil {
			err = errors.New("text exceeds 255 characters (truncated) and contains invalid characters (replaced)")
		} else {
			err = errors.New("text contains invalid characters (replaced)")
		}
	} else if !regex.MatchString(result) {
		if err != nil {
			err = errors.New("text exceeds 255 characters (truncated) and is not a valid float format")
		} else {
			err = errors.New("text is not a valid float format")
		}
	}

	return result, err
}
