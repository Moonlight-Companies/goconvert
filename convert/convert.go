package convert

import (
	"fmt"
	"strconv"
)

func ConvertInto[T any](interfaceValue interface{}) (result T, ok bool) {
	if value, ok := interfaceValue.(T); ok {
		return value, true
	}

	switch any(result).(type) {
	case int: // converting into int
		if v, ok := interfaceValue.(int); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(int(1)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(int(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(int(intermediate)).(T)

				return result, true
			}
		}

	case int8: // converting into int8
		if v, ok := interfaceValue.(int); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(int8(1)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(int8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(int8(intermediate)).(T)

				return result, true
			}
		}

	case int16: // converting into int16
		if v, ok := interfaceValue.(int); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(int16(1)).(T)

				return result, true
			}
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(int16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(int16(intermediate)).(T)

				return result, true
			}
		}

	case int32: // converting into int32
		if v, ok := interfaceValue.(int); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(int32(1)).(T)

				return result, true
			}
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(int32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(int32(intermediate)).(T)

				return result, true
			}
		}

	case int64: // converting into int64
		if v, ok := interfaceValue.(int); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(int64(1)).(T)

				return result, true
			}
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(int64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(int64(intermediate)).(T)

				return result, true
			}
		}

	case uint: // converting into uint
		if v, ok := interfaceValue.(int); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(uint(1)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(uint(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(uint(intermediate)).(T)

				return result, true
			}
		}

	case uint8: // converting into uint8
		if v, ok := interfaceValue.(int); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(uint8(1)).(T)

				return result, true
			}
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(uint8(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(uint8(intermediate)).(T)

				return result, true
			}
		}

	case uint16: // converting into uint16
		if v, ok := interfaceValue.(int); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(uint16(1)).(T)

				return result, true
			}
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(uint16(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(uint16(intermediate)).(T)

				return result, true
			}
		}

	case uint32: // converting into uint32
		if v, ok := interfaceValue.(int); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(uint32(1)).(T)

				return result, true
			}
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(uint32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(uint32(intermediate)).(T)

				return result, true
			}
		}

	case uint64: // converting into uint64
		if v, ok := interfaceValue.(int); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(uint64(1)).(T)

				return result, true
			}
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(uint64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(uint64(intermediate)).(T)

				return result, true
			}
		}

	case bool: // converting into bool
		if v, ok := interfaceValue.(int); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if !v {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(float32); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			if v != 0 {
				result = any(bool(true)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if v != "" && v != "0" && v != "false" {
				result = any(bool(true)).(T)
			}

			return result, true
		}

	case float32: // converting into float32
		if v, ok := interfaceValue.(int); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(float32(1)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(float32(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(float32(intermediate)).(T)

				return result, true
			}
		}

	case float64: // converting into float64
		if v, ok := interfaceValue.(int); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any(float64(1)).(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(float64(v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			if intermediate, err := strconv.ParseFloat(v, 64); err == nil {
				result = any(float64(intermediate)).(T)

				return result, true
			}
		}

	case string: // converting into string
		if v, ok := interfaceValue.(int); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int8); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int16); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int32); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint8); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint16); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint32); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(uint64); ok {
			result = any(fmt.Sprintf("%d", v)).(T)
			return result, true
		}

		if v, ok := interfaceValue.(bool); ok {
			if v {
				result = any("1").(T)
			}

			return result, true
		}
		if v, ok := interfaceValue.(float32); ok {
			result = any(fmt.Sprintf("%f", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float64); ok {
			result = any(fmt.Sprintf("%f", v)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(string); ok {
			result = any(string(v)).(T)

			return result, true
		}

	}

	return result, false
}
