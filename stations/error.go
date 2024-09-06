package stations

import (
	"fmt"
	"runtime"
)

func SourcedError(inner error) error {
	if inner == nil {
		return nil
	}
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("error at %s:%d - %w", file, line, inner)
}
