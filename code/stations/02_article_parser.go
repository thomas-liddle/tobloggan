package stations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"tobloggan/code/contracts"
)

type ArticleParser struct {
}

//func (this *ArticleParser) Do(input any, output func(any)) {
//    TODO: given a contracts.SourceFile, parse the JSON metadata and save the body on a contracts.Article.
//    input: contracts.SourceFile
//    input: contracts.Article
//}

func NewArticleParser() *ArticleParser {
	return &ArticleParser{}
}

var separator = []byte{'+', '+', '+'}

func (this *ArticleParser) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceFile:
		var article contracts.Article
		jsonHalf, bodyPart, found := bytes.Cut(input, separator)
		if !found {
			output(fmt.Errorf("no separator: %w", errMalformedContent))
			return
		}
		err := json.Unmarshal(jsonHalf, &article)
		if err != nil {
			output(fmt.Errorf("%w: %w", errMalformedContent, err))
			return
		}
		article.Body = string(bytes.TrimSpace(bodyPart))
		output(article)
	default:
		output(input)
	}
}

var errMalformedContent = errors.New("malformed content")
