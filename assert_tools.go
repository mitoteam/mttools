package mttools

func AssertNil(v any, message ...string) {
	if v != nil {
		panic(messageFromArgs("AssertNil failed", message))
	}
}

func AssertNotNil(v any, message ...string) {
	if v == nil {
		panic(messageFromArgs("AssertNotNil failed", message))
	}
}

func AssertEmpty(v any, message ...string) {
	if !IsEmpty(v) {
		panic(messageFromArgs("AssertEmpty failed", message))
	}
}

func AssertNotEmpty(v any, message ...string) {
	if IsEmpty(v) {
		panic(messageFromArgs("AssertNotEmpty failed", message))
	}
}

func AssertEqual[T comparable](a T, b T, message ...string) {
	if a != b {
		panic(messageFromArgs("AssertEqual failed", message))
	}
}

func messageFromArgs(defaultMessage string, messageArgs []string) string {
	if len(messageArgs) > 0 {
		return messageArgs[0]
	} else {
		return defaultMessage
	}
}
