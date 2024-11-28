package mttools

import "fmt"

// Returns number of `value` values found in `values_list`
func CountValues(value interface{}, values_list ...interface{}) (count int) {
	count = 0

	for _, element := range values_list {
		if element == value {
			count++
		}
	}
	return count
}

// Returns slice of same type with just unique elements left in it
func UniqueSlice[S ~[]E, E any](slice S) S {
	unique_map := make(map[string]bool, len(slice))

	result := make(S, 0, len(slice))
	for _, value := range slice {
		var unique_key string
		var ok bool

		if unique_key, ok = any(value).(string); !ok {
			unique_key = fmt.Sprint(value)
		}

		if len(unique_key) != 0 {
			if !unique_map[unique_key] {
				result = append(result, value)
				unique_map[unique_key] = true
			}
		}
	}

	return result
}
