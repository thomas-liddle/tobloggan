package stations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"

	"github.com/mdwhatcott/tobloggan/contracts"
)

type SourceReader struct {
	fs fs.FS
	md contracts.Markdown
}

func NewSourceReader(fs fs.FS, md contracts.Markdown) *SourceReader {
	return &SourceReader{fs: fs, md: md}
}

func (this *SourceReader) Do(input any, output func(v any)) {
	switch input := input.(type) {
	case contracts.SourceFilePath:
		raw, err := fs.ReadFile(this.fs, string(input))
		if err != nil {
			output(fmt.Errorf("%w: %s", err, input))
			return
		}
		front, body, divided := bytes.Cut(raw, []byte("\n+++\n"))
		if !divided {
			output(fmt.Errorf("%w (missing divider): %s", contracts.ErrMalformedSource, input))
			return
		}
		var source contracts.Article
		err = json.Unmarshal(front, &source)
		if err != nil {
			output(fmt.Errorf("%w (%w): %s", contracts.ErrMalformedSource, err, input))
			return
		}
		source.Body, err = this.md.Convert(string(bytes.TrimSpace(body)))
		if err != nil {
			output(fmt.Errorf("%w (%w): %s", contracts.ErrMalformedSource, err, input))
			return
		}
		output(source)
	}
}
