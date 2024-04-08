package mttools

import (
	"reflect"
)

// Checks if s - variable of struct type and it embeds embedded_type
func IsStructEmbeds(s interface{}, embedded_type reflect.Type) bool {
	if s == nil {
		return false
	}

	st := reflect.TypeOf(s)

	//dereference pointer
	if st.Kind() == reflect.Pointer {
		st = st.Elem()
	}

	//check if they both are structures
	if st.Kind() != reflect.Struct || embedded_type.Kind() != reflect.Struct {
		return false
	}

	//check fields
	for i := 0; i < st.NumField(); i++ {
		struct_field := st.Field(i)

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
