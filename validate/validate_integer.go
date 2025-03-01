package validate

import (
	"errors"
	"fmt"
	"strings"
)

// IntegerState represents the current parsing state
type IntegerState int

const (
	IntegerStart IntegerState = iota
	IntegerAfterMinus
	IntegerDigits
)

func ValidateBasicInteger(txt string) (string, error) {
	var sanitized strings.Builder
	var invalidChars bool

	// Process byte by byte instead of rune by rune
	bytes := []byte(txt)
	var hasSeenDigit bool
	var hasSeenMinus bool

	for i, b := range bytes {
		if i >= 255 {
			sanitized.WriteString("...")
			break
		}

		if b >= '0' && b <= '9' {
			sanitized.WriteByte(b)
			hasSeenDigit = true
		} else if b == '-' && !hasSeenDigit && !hasSeenMinus {
			sanitized.WriteByte(b)
			hasSeenMinus = true
		} else {
			sanitized.WriteString(fmt.Sprintf("<0x%02X>", b))
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
	} else if !isValidInteger(result) {
		if err != nil {
			err = errors.New("text exceeds 255 characters (truncated) and is not a valid integer format")
		} else {
			err = errors.New("text is not a valid integer format")
		}
	}

	return result, err
}

// isValidInteger uses an FSM to validate the integer format
func isValidInteger(s string) bool {
	if s == "" {
		return false
	}

	state := IntegerStart
	digitSeen := false

	// Process each byte
	for i := 0; i < len(s); i++ {
		b := s[i]

		switch state {
		case IntegerStart:
			if b == '-' {
				state = IntegerAfterMinus
			} else if b >= '0' && b <= '9' {
				state = IntegerDigits
				digitSeen = true
			} else {
				return false
			}

		case IntegerAfterMinus:
			if b >= '0' && b <= '9' {
				state = IntegerDigits
				digitSeen = true
			} else {
				return false
			}

		case IntegerDigits:
			if b >= '0' && b <= '9' {
				// remain in Digits state
			} else {
				return false
			}
		}
	}

	// Must have at least one digit
	return digitSeen
}
