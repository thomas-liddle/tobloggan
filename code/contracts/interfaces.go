package contracts

import (
	"os"
	"time"
)

type Clock func() time.Time

type Logger interface {
	Printf(format string, args ...interface{})
}

type (
	Station interface {
		Do(input any, output func(any))
	}
	Finalizer interface {
		Finalize(output func(any))
	}
)

type (
	FSWriter interface {
		MkdirAll
		WriteFile
	}
	MkdirAll interface {
		MkdirAll(path string, perm os.FileMode) error
	}
	WriteFile interface {
		WriteFile(filename string, data []byte, perm os.FileMode) error
	}
)
