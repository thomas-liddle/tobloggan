package stations

import (
	"io/fs"

	"tobloggan/code/contracts"
)

type SourceScanner struct {
	fs fs.FS
}

func NewSourceScanner(fs fs.FS) contracts.Station {
	return &SourceScanner{fs: fs}
}

func (this *SourceScanner) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceDirectory:
	default:
		output(input)
	}
}
