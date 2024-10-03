package stations

import (
	"fmt"
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
		err := fs.WalkDir(this.fs, string(input), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if endsWith(d.Name(), ".md") && !d.IsDir() {
				d.Type()
				output(contracts.SourceFilePath(path))
			}
			return nil
		})
		if err != nil {
			output(fmt.Errorf("error reading directory: %s, %w", input, err))
		}

	default:
		output(input)
	}
}

func endsWith(s, expectedEnding string) bool {
	return len(s) >= len(expectedEnding) && s[len(s)-len(expectedEnding):] == expectedEnding
}
