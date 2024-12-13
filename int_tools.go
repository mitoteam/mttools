package mttools

import (
	"reflect"
	"strconv"
)

// Convert any value to Int64. Ok = true if it was converted
func AnyToInt64Ok(v any) (int64, bool) {
	switch i := v.(type) {
	case nil:
		return 0, false

	case int64:
		return i, true

	case int32:
	case uint32:
	case int16:
	case uint16:
	case int8:
	case uint8:
	case int:
	case uint:
		return int64(i), true

	default:
		// handle the remaining type set of ~int64
		r := reflect.ValueOf(v)
		if r.Kind() == reflect.Int64 {
			return r.Int(), true
		} else {
			//try to cast it to string and then parse it
			if s, ok := AnyToStringOk(v); ok {
				if i, err := strconv.ParseInt(s, 10, 64); err == nil {
					return i, true
				}
			}
		}
	}

	//not even a string...
	return 0, false
}

// Try best to convert any value to int64. Returns 0 if value can not be converted.
func AnyToInt64OrZero(v any) int64 {
	if i, ok := AnyToInt64Ok(v); ok {
		return i
	}

	return 0
}
