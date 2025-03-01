package validate

import (
	"fmt"
)

// Sanitize provides a quick and clean way to sanitize various data types for display
func Sanitize(v interface{}) string {
	switch val := v.(type) {
	case string:
		result, _ := ValidateBasicText(val)
		return result
	case int, int8, int16, int32, int64:
		result, _ := ValidateBasicInteger(fmt.Sprintf("%d", val))
		return result
	case uint, uint8, uint16, uint32, uint64:
		result, _ := ValidateBasicInteger(fmt.Sprintf("%d", val))
		return result
	case float32, float64:
		result, _ := ValidateBasicFloat(fmt.Sprintf("%g", val))
		return result
	case bool:
		if val {
			return "true"
		}
		return "false"
	case nil:
		return "<nil>"
	default:
		// For any other type, convert to string and validate as text
		vv := fmt.Sprintf("%v", val)
		result, _ := ValidateBasicText(vv)
		return result
	}
}
