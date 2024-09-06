package contracts

import (
	"fmt"
	"runtime"
)

func Error(inner error) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("error at %s:%d - %w", file, line, inner)
}
func Errorf(format string, a ...any) error {
	_, file, line, _ := runtime.Caller(1)
	inner := fmt.Errorf(format, a...)
	return fmt.Errorf("error at %s:%d - %w", file, line, inner)
}
