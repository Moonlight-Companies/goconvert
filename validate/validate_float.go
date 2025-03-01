package validate

import (
	"errors"
	"fmt"
	"strings"
)

// FloatState represents the current parsing state
type FloatState int

const (
	FloatStart FloatState = iota
	FloatAfterMinus
	FloatInteger
	FloatAfterDecimal
	FloatFraction
)

func ValidateBasicFloat(txt string) (string, error) {
	var sanitized strings.Builder
	var invalidChars bool

	// Process byte by byte instead of rune by rune
	bytes := []byte(txt)
	for i, b := range bytes {
		if i >= 255 {
			sanitized.WriteString("...")
			break
		}

		// Only allow ASCII digits, period, and minus
		if (b >= '0' && b <= '9') || b == '.' || b == '-' {
			sanitized.WriteByte(b)
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
	} else if !isValidFloat(result) {
		if err != nil {
			err = errors.New("text exceeds 255 characters (truncated) and is not a valid float format")
		} else {
			err = errors.New("text is not a valid float format")
		}
	}

	return result, err
}

// isValidFloat uses an FSM to validate the float format
func isValidFloat(s string) bool {
	if s == "" {
		return false
	}

	state := FloatStart
	decimalSeen := false
	digitSeen := false

	// Process each byte
	for i := 0; i < len(s); i++ {
		b := s[i]

		switch state {
		case FloatStart:
			if b == '-' {
				state = FloatAfterMinus
			} else if b >= '0' && b <= '9' {
				state = FloatInteger
				digitSeen = true
			} else if b == '.' {
				state = FloatAfterDecimal
				decimalSeen = true
			} else {
				return false
			}

		case FloatAfterMinus:
			if b >= '0' && b <= '9' {
				state = FloatInteger
				digitSeen = true
			} else if b == '.' {
				state = FloatAfterDecimal
				decimalSeen = true
			} else {
				return false
			}

		case FloatInteger:
			if b >= '0' && b <= '9' {
				// remain in integer state
			} else if b == '.' && !decimalSeen {
				state = FloatAfterDecimal
				decimalSeen = true
			} else {
				return false
			}

		case FloatAfterDecimal:
			if b >= '0' && b <= '9' {
				state = FloatFraction
				digitSeen = true
			} else {
				return false
			}

		case FloatFraction:
			if b >= '0' && b <= '9' {
				// remain in fraction state
			} else {
				return false
			}
		}
	}

	// Must have at least one digit
	return digitSeen
}
