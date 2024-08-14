package convert

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ValidateBasicText(txt string) (string, error) {
	var sanitized strings.Builder
	var invalidChars bool
	regex := regexp.MustCompile(`^[ a-zA-Z!@#$%^&*()-_/\\?]$`)

	for i, r := range txt {
		if i >= 255 {
			sanitized.WriteString("...")
			break
		}

		if regex.MatchString(string(r)) {
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
	} else if invalidChars {
		err = errors.New("text contains invalid characters (replaced)")
	}

	return result, err
}
