package stations

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/mdwhatcott/tobloggan/contracts"
)

type Markdown interface {
	Convert(content string) (string, error)
}

type ArticleParser struct {
	md Markdown
}

func NewArticleParser(md Markdown) *ArticleParser {
	return &ArticleParser{md: md}
}

func (this *ArticleParser) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceFile:
		front, body, divided := bytes.Cut(input, []byte("\n+++\n"))
		if !divided {
			output(StackTraceError(fmt.Errorf("%w (missing divider): %s", contracts.ErrMalformedSource, input)))
			return
		}
		var source contracts.Article
		err := json.Unmarshal(front, &source)
		if err != nil {
			output(StackTraceError(fmt.Errorf("%w (%w): %s", contracts.ErrMalformedSource, err, input)))
			return
		}
		source.Body, err = this.md.Convert(string(bytes.TrimSpace(body)))
		if err != nil {
			output(StackTraceError(fmt.Errorf("%w (%w): %s", contracts.ErrMalformedSource, err, input)))
			return
		}
		output(source)
	default:
		output(input)
	}
}
