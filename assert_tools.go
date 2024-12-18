package mttools

func AssertNil(v any) {
	if v != nil {
		panic("AssertNil failed")
	}
}

func AssertNotNil(v any) {
	if v == nil {
		panic("AssertNotNil failed")
	}
}

func AssertEmpty(v any) {
	if !IsEmpty(v) {
		panic("AssertEmpty failed")
	}
}

func AssertNotEmpty(v any) {
	if IsEmpty(v) {
		panic("AssertNotEmpty failed")
	}
}

func AssertEqual[T comparable](a T, b T) {
	if a != b {
		panic("AssertEqual failed")
	}
}
