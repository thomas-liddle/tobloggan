package stations

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/mdwhatcott/tobloggan/code/contracts"
)

type ArticleParser struct{}

func NewArticleParser() contracts.Station {
	return &ArticleParser{}
}

func (this *ArticleParser) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceFile:
		front, body, divided := bytes.Cut(input, []byte("\n+++\n"))
		if !divided {
			output(contracts.Errorf("%w (missing divider): %s", errMalformedSource, input))
			return
		}
		var source contracts.Article
		err := json.Unmarshal(front, &source)
		if err != nil {
			output(contracts.Errorf("%w (%w): %s", errMalformedSource, err, input))
			return
		}
		source.Body = string(bytes.TrimSpace(body))
		output(source)
	default:
		output(input)
	}
}

var errMalformedSource = errors.New("malformed source")
