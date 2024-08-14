package convert

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ValidateBasicInteger(txt string) (string, error) {
	var sanitized strings.Builder
	var invalidChars bool
	var hasSeenDigit bool
	var hasSeenMinus bool

	// Regex explanation:
	// ^    : Start of the string
	// -?   : Optional minus sign (zero or one occurrence)
	// \d+  : One or more digits
	// $    : End of the string
	// This regex ensures the entire string is a valid integer format
	regex := regexp.MustCompile(`^-?\d+$`)

	// Always process character by character to ensure proper truncation
	for i, r := range txt {
		if i >= 255 {
			sanitized.WriteString("...")
			break
		}

		if r >= '0' && r <= '9' {
			sanitized.WriteRune(r)
			hasSeenDigit = true
		} else if r == '-' && !hasSeenDigit && !hasSeenMinus {
			sanitized.WriteRune(r)
			hasSeenMinus = true
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
			err = errors.New("text exceeds 255 characters (truncated) and is not a valid integer format")
		} else {
			err = errors.New("text is not a valid integer format")
		}
	}

	return result, err
}
