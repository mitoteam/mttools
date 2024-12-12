package mttools

import (
	"reflect"
)

// Checks if s - variable of struct type and it embeds embedded_type
func IsStructEmbeds(s interface{}, embedded_type reflect.Type) bool {
	if s == nil {
		return false
	}

	return IsStructTypeEmbeds(reflect.TypeOf(s), embedded_type)
}

// Checks if s_type - is a struct (or pointer to it) and it embeds embedded_type
func IsStructTypeEmbeds(s_type reflect.Type, embedded_type reflect.Type) bool {
	//dereference pointer
	if s_type.Kind() == reflect.Pointer {
		s_type = s_type.Elem()
	}

	//check if they both are structures
	if s_type.Kind() != reflect.Struct || embedded_type.Kind() != reflect.Struct {
		return false
	}

	//check fields
	for i := 0; i < s_type.NumField(); i++ {
		struct_field := s_type.Field(i)

		if struct_field.Type == embedded_type {
			return true
		}

		//recursion
		if struct_field.Type.Kind() == reflect.Struct {
			if IsStructEmbeds(struct_field.Type, embedded_type) {
				return true
			}
		}
	}

	return false
}
