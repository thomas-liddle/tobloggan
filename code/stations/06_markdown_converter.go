package stations

import (
	"errors"
	"fmt"

	"tobloggan/code/contracts"
)

type Markdown interface {
	Convert(content string) (string, error)
}

type MarkdownConverter struct {
	md Markdown
}

func NewMarkdownConverterStation(md Markdown) contracts.Station {
	return &MarkdownConverter{md: md}
}

func (this *MarkdownConverter) Do(input any, output func(any)) {
	//TODO: given a contracts.Article, use the provided Markdown interface to convert and re-assign the Body field.
	switch input := input.(type) {
	case contracts.Article:
		mdConvert, err := this.md.Convert(input.Body)
		if err != nil {
			output(fmt.Errorf("%w: %s", err, input))
			return
		}
		input.Body = mdConvert
		output(input)

	default:
		output(input)
	}
}

var ErrInvalidMarkdown = errors.New("invalid markdown")
