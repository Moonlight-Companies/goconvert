package validate

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

	// Convert string to bytes
	bytes := []byte(txt)

	// Process bytes
	currentLen := 0
	i := 0

	for i < len(bytes) {
		// Check what the next chunk would be
		isValid := regex.MatchString(string(rune(bytes[i])))
		nextLen := 1
		if !isValid {
			nextLen = 6 // <0xXX>
			invalidChars = true
		}

		// Check if we can fit this chunk
		if currentLen+nextLen > 255 {
			break
		}

		// Write the chunk
		if isValid {
			sanitized.WriteByte(bytes[i])
		} else {
			sanitized.WriteString(fmt.Sprintf("<0x%02X>", bytes[i]))
		}

		currentLen += nextLen
		i++
	}

	// Only add ... if we have more bytes to process
	var err error
	if i < len(bytes) {
		// Special case: if we have at least 6 chars left and the next byte would
		// be invalid, try to include it
		if i < len(bytes) && currentLen <= 249 {
			sanitized.WriteString(fmt.Sprintf("<0x%02X>", bytes[i]))
			i++
		}
		// Still have more bytes? Add ...
		if i < len(bytes) {
			sanitized.WriteString("...")
		}
		err = errors.New("text exceeds 255 characters (truncated)")
	} else if invalidChars {
		err = errors.New("text contains invalid characters (replaced)")
	}

	return sanitized.String(), err
}
