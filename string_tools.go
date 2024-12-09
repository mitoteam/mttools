package mttools

import (
	"fmt"
	"reflect"
	"regexp"
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
