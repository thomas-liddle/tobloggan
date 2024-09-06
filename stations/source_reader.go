package stations

import (
	"fmt"
	"io/fs"

	"github.com/mdwhatcott/tobloggan/contracts"
)

type SourceReader struct {
	fs fs.FS
}

func NewSourceReader(fs fs.FS) *SourceReader {
	return &SourceReader{fs: fs}
}

func (this *SourceReader) Do(input any, output func(v any)) {
	switch input := input.(type) {
	case contracts.SourceFilePath:
		raw, err := fs.ReadFile(this.fs, string(input))
		if err != nil {
			output(SourcedError(fmt.Errorf("%w: %s", err, input)))
		} else {
			output(contracts.SourceFile(raw))
		}
	default:
		output(input)
	}
}
