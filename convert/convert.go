package convert

import (
	"fmt"
	"strconv"
	"time"
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
			result = any(v).(T)
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

	case time.Time: // converting into time.Time
		if v, ok := interfaceValue.(time.Time); ok {
			result = any(time.Time(v)).(T)
			return result, true
		}
		// Try to convert from string using various date formats
		if v, ok := interfaceValue.(string); ok {
			// Common date formats to try
			dateFormats := []string{
				"2006-01-02",                    // YYYY-MM-DD
				"Jan 2 2006 3:04 PM",            // Mon DD YYYY HH:MM PM
				"Jan 02 2006 3:04 PM",           // Mon DD YYYY HH:MM PM
				"2006-01-02T15:04:05Z07:00",     // RFC3339
				"2006-01-02T15:04:05.999Z07:00", // RFC3339Nano
				"1/2/2006",                      // M/D/YYYY
				"01/02/2006",                    // MM/DD/YYYY
				"01022006",                      // MMDDYYYY
				"2006/01/02",                    // YYYY/MM/DD
				"2-Jan-2006",                    // D-Mon-YYYY
				"2 Jan 2006",                    // D Mon YYYY
				"02-Jan-06",                     // DD-Mon-YY
				"Jan 2, 2006",                   // Mon D, YYYY
				"January 2, 2006",               // Month D, YYYY
				"2006-01-02 15:04:05",           // YYYY-MM-DD HH:MM:SS
				"02/01/2006 15:04:05",           // DD/MM/YYYY HH:MM:SS
				time.RFC822,
				time.RFC822Z,
				time.RFC850,
				time.RFC1123,
				time.RFC1123Z,
				time.RFC3339,
				time.RFC3339Nano,
				time.ANSIC,
				time.UnixDate,
				time.RubyDate,
			}

			for _, format := range dateFormats {
				fmt.Println("TRYING", format)
				if parsedTime, err := time.Parse(format, v); err == nil {
					result = any(parsedTime).(T)
					return result, true
				} else {
					fmt.Println("FAILED", format, err)
				}
			}

			// Try to parse integer timestamp
			if timestamp, err := strconv.ParseInt(v, 10, 64); err == nil {
				// Check if this might be a JavaScript millisecond timestamp
				if timestamp > 1000000000000 { // Timestamps after 2001 in milliseconds are > 10^12
					result = any(time.Unix(timestamp/1000, (timestamp%1000)*1000000)).(T)
					return result, true
				}
				result = any(time.Unix(timestamp, 0)).(T)
				return result, true
			}

			// Try to parse float timestamp (unix seconds with fractional part)
			if timestamp, err := strconv.ParseFloat(v, 64); err == nil {
				// Check if this might be a JavaScript millisecond timestamp
				if timestamp > 1000000000000 { // Timestamps after 2001 in milliseconds are > 10^12
					seconds := timestamp / 1000
					secondsInt := int64(seconds)
					nanoseconds := int64((seconds - float64(secondsInt)) * 1e9)
					result = any(time.Unix(secondsInt, nanoseconds)).(T)
					return result, true
				}

				seconds := int64(timestamp)
				nanoseconds := int64((timestamp - float64(seconds)) * 1e9)
				result = any(time.Unix(seconds, nanoseconds)).(T)
				return result, true
			}
		}

		// Try to convert from int/int64 as unix timestamp
		if v, ok := interfaceValue.(int); ok {
			// Check if this might be a JavaScript millisecond timestamp
			if v > 1000000000000 { // Timestamps after 2001 in milliseconds are > 10^12
				seconds := v / 1000
				millis := v % 1000
				result = any(time.Unix(int64(seconds), int64(millis)*1000000)).(T)
				return result, true
			}
			result = any(time.Unix(int64(v), 0)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(int64); ok {
			// Check if this might be a JavaScript millisecond timestamp
			if v > 1000000000000 { // Timestamps after 2001 in milliseconds are > 10^12
				seconds := v / 1000
				millis := v % 1000
				result = any(time.Unix(seconds, millis*1000000)).(T)
				return result, true
			}
			result = any(time.Unix(v, 0)).(T)
			return result, true
		}

		// Try to convert from float as unix timestamp with fractions
		if v, ok := interfaceValue.(float64); ok {
			// Check if this might be a JavaScript millisecond timestamp
			if v > 1000000000000 { // Timestamps after 2001 in milliseconds are > 10^12
				v = v / 1000 // Convert to seconds
			}
			seconds := int64(v)
			nanoseconds := int64((v - float64(seconds)) * 1e9)
			result = any(time.Unix(seconds, nanoseconds)).(T)
			return result, true
		}
		if v, ok := interfaceValue.(float32); ok {
			v64 := float64(v)
			// Check if this might be a JavaScript millisecond timestamp
			if v64 > 1000000000000 { // Timestamps after 2001 in milliseconds are > 10^12
				v64 = v64 / 1000 // Convert to seconds
			}
			seconds := int64(v64)
			// Use exact calculation to avoid float32 precision issues
			fractional := v64 - float64(seconds)
			nanoseconds := int64(fractional * 1e9)
			result = any(time.Unix(seconds, nanoseconds)).(T)

			// And finally result and false because we should NOT even be doing this conversion
			return result, false
		}
	}

	return result, false
}
