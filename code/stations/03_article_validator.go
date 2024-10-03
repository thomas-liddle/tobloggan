package stations

import (
	"fmt"
	"strings"

	"tobloggan/code/contracts"
	"tobloggan/code/set"
)

type ArticleValidator struct {
	unique set.Set[string]
}

func NewArticleValidator() *ArticleValidator {
	return &ArticleValidator{unique: set.New[string]()}
}

func (this *ArticleValidator) Do(input any, output func(any)) {
	//    input: contracts.Article
	//    output: contracts.Article (or error)
	switch input := input.(type) {
	case contracts.Article:
		if this.isValidSlug(input.Slug) {
			if this.unique.Contains(input.Slug) {
				output(ErrDuplicateSlug)
				return
			}
		} else {
			output(ErrInvalidSlug)
			return
		}
		if input.Title == "" {
			output(ErrInvalidTitle)
			return
		}
		this.unique.Add(input.Slug)
		output(input)

	default:
		output(input)
	}
}

func (this *ArticleValidator) isValidSlug(slug string) bool {
	if slug == "" || strings.Contains(slug, "//") {
		return false
	}

	for _, c := range slug {
		if !validSlugCharacters.Contains(c) {
			return false
		}
	}
	return true
}

var (
	validSlugCharacters = set.New([]rune("abcdefghijklmnopqrstuvwxyz0123456789-/")...)
	ErrDuplicateSlug    = fmt.Errorf("duplicate slug")
	ErrInvalidSlug      = fmt.Errorf("invalid slug")
	ErrInvalidTitle     = fmt.Errorf("invalid title")
)
