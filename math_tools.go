package mttools

// Returns [+1;0;-1] depending on value's sign.
func IntSign(value int) int {
	if value > 0 {
		return +1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}

// Returns [+1;0;-1] depending on value's sign.
func Int64Sign(value int64) int64 {
	if value > 0 {
		return +1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}
