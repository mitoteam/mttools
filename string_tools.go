package mttools

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

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

// Splits string into array of strings.
// Each argument in string is single word or number, or string in double quotes.
func SplitArgumentsString(str string) []string {
	str = strings.TrimSpace(str)

	// \p{L} = any unicode letter
	re := regexp.MustCompile(`([\p{L}\d_\-\+\@\|\!\/]+)|(".*?[^\\]")|("")`)

	result := re.FindAllString(str, -1) //-1 = all

	// remove double quotes
	for index, item := range result {
		if len(item) > 0 && item[0] == '"' {
			item = item[1 : len(item)-1]

			result[index] = item
		}
	}

	//DBG
	//fmt.Printf("%+v\n", result)

	return result
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
