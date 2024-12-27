package mttools

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func CallerSignature(skip int) string {
	if pc, file, line, ok := runtime.Caller(skip); ok {
		funcName := runtime.FuncForPC(pc).Name()
		return fmt.Sprintf("%s() at %s:%d", funcName, filepath.Base(file), line)
	}

	return "[unknown caller]"
}

func PanicWithSignature(message string, args ...any) string {
	s := CallerSignature(2)

	if len(message) > 0 {
		s += ": " + fmt.Sprintf(message, args...)
	}

	panic(s)
}
