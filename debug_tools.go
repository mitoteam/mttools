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

func PanicWithSignature(message string) string {
	s := CallerSignature(2)

	if len(message) > 0 {
		s += ": " + message
	}

	panic(s)
}

func PanicWithSignatureF(message string, args ...any) string {
	panic(CallerSignature(2) + ": " + fmt.Sprintf(message, args...))
}