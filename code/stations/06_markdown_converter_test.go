package stations

import (
	"testing"

	"tobloggan/code/contracts"
	"tobloggan/code/markdown"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

type badMd struct {
}

func (this *badMd) Convert(content string) (string, error) {
	return content, ErrInvalidMarkdown
}

func TestMarkdownConverterFixture(t *testing.T) {
	gunit.Run(new(MarkdownConverterFixture), t)
}

type MarkdownConverterFixture struct {
	*gunit.Fixture
}

func (this *MarkdownConverterFixture) TestBodyConverted() {
	var outputs []any
	station := NewMarkdownConverterStation(markdown.NewConverter())
	station.Do(goodArticle, func(v any) {
		outputs = append(outputs, v)
	})

	this.So(outputs, should.HaveLength, 1)
	_, isArticle := outputs[0].(contracts.Article)
	this.So(isArticle, should.BeTrue)
}

func (this *MarkdownConverterFixture) TestInvalidMarkdown() {
	var outputs []any
	station := NewMarkdownConverterStation(&badMd{})
	station.Do(goodArticle, func(v any) {
		outputs = append(outputs, v)
	})

	this.So(outputs, should.HaveLength, 1)
	this.So(outputs[0], should.Wrap, ErrInvalidMarkdown)
}
