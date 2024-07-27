package message

import "github.com/Moonlight-Companies/goconvert/convert"

func Read[T any](m map[string]interface{}, key string) (value T, ok bool) {
	temp, ok := m[key]
	if !ok {
		return value, false
	}

	return convert.ConvertInto[T](temp)
}

func ReadList[T any](m map[string]interface{}, key string) (values []T, ok bool) {
	temp, ok := m[key]
	if !ok {
		return values, false
	}

	// we still intend to return the values that were converted successfully and
	// "ok" false if there was any failure
	success := true
	if list, ok := temp.([]interface{}); ok {
		values = make([]T, len(list))
		for i, item := range list {
			if value, ok := convert.ConvertInto[T](item); ok {
				values[i] = value
			} else {
				success = false
			}
		}
	} else {
		return values, false
	}

	return values, success
}
