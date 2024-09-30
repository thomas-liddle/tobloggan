package stations

import (
	"errors"
	"strings"

	"tobloggan/code/contracts"
	"tobloggan/code/set"
)

type ArticleValidator struct{}

func NewArticleValidator() contracts.Station {
	return &ArticleValidator{}
}

func (this *ArticleValidator) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		if !this.isValidSlug(input.Slug) {
			output(contracts.Errorf("%w (slug): [%s]", errInvalidContent, input.Slug))
		} else if !this.isValidTitle(input.Title) {
			output(contracts.Errorf("%w (title): [%s]", errInvalidContent, input.Title))
		} else {
			output(input)
		}
	default:
		output(input)
	}
}
func (this *ArticleValidator) isValidSlug(slug string) bool {
	if slug == "" {
		return false
	}
	if len(slug) > 128 {
		return false
	}
	if strings.Contains(slug, "//") {
		return false
	}
	for _, c := range slug {
		if !validSlugCharacters.Contains(c) {
			return false
		}
	}
	return true
}
func (this *ArticleValidator) isValidTitle(title string) bool {
	if title == "" {
		return false
	}
	if len(title) > 256 {
		return false
	}
	return true
}

var errInvalidContent = errors.New("invalid content")

var validSlugCharacters = set.New([]rune("abcdefghijklmnopqrstuvwxyz0123456789-/")...)
