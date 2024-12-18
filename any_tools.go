package mttools

import (
	"fmt"
	"reflect"
	"strconv"
)

// Converters from interface{} to particular types

// Convert any value to string. Ok = true if it is a string or Stringer, false if not.
// https://stackoverflow.com/questions/72267243/unioning-an-interface-with-a-type-in-golang
func AnyToStringOk(v any) (s string, b bool) {
	switch v := v.(type) {
	case nil:
		s = ""

	case string:
		s = v

	case fmt.Stringer:
		s = v.String()

	case uint64:
		s = strconv.FormatUint(v, 10)

	case uint32:
		s = strconv.FormatUint(uint64(v), 10)

	case uint16:
		s = strconv.FormatUint(uint64(v), 10)

	case uint8:
		s = strconv.FormatUint(uint64(v), 10)

	case uint:
		s = strconv.FormatUint(uint64(v), 10)

	case int64:
		s = strconv.FormatInt(v, 10)

	case int32:
		s = strconv.FormatInt(int64(v), 10)

	case int16:
		s = strconv.FormatInt(int64(v), 10)

	case int8:
		s = strconv.FormatInt(int64(v), 10)

	case int:
		s = strconv.FormatInt(int64(v), 10)

	case bool:
		s = strconv.FormatBool(v)

	default:
		// handle the remaining type set of ~string
		r := reflect.ValueOf(v)
		if r.Kind() == reflect.String {
			s = r.String()
		} else {
			return "", false
		}
	}

	return s, true
}

// Try best to convert any value to string, even if it is not a string at all.
func AnyToString(v any) string {
	if s, ok := AnyToStringOk(v); ok {
		return s
	}

	// Last chance: lets Go render it as string.
	return fmt.Sprintf("%v", v)
}

// Convert any value to Int64. Ok = true if it was converted
func AnyToInt64Ok(v any) (int64, bool) {
	switch i := v.(type) {
	case nil:
		return 0, false

	case int64:
		return i, true

	case int32:
		return int64(i), true

	case uint32:
		return int64(i), true

	case int16:
		return int64(i), true

	case uint16:
		return int64(i), true

	case int8:
		return int64(i), true

	case uint8:
		return int64(i), true

	case int:
		return int64(i), true

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

// Any value to bool.
func AnyToBool(v any) bool {
	if v == nil {
		return false
	}

	switch v := v.(type) {
	case bool:
		return v

	default:
		return !IsEmpty(v)
	}
}

// Returns true if v value considered as empty (nil, "", 0 for int, 0.0 for float, false for bool).
func IsEmpty(v any) bool {
	if v == nil {
		return true
	}

	if v, ok := v.(bool); ok {
		return !v
	}

	if v, ok := v.(float64); ok {
		return v == 0.0
	}

	if v, ok := v.(float32); ok {
		return v == 0.0
	}

	if v, ok := AnyToInt64Ok(v); ok {
		return v == 0
	}

	if v, ok := AnyToStringOk(v); ok {
		return v == ""
	}

	return false
}
