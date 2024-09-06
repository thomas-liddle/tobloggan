package stations

import (
	"errors"
	"testing"
	"time"

	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleParserFixture(t *testing.T) {
	gunit.Run(new(ArticleParserFixture), t)
}

type ArticleParserFixture struct {
	*gunit.Fixture
	StationFixture
	parser      *ArticleParser
	markdownErr error
}

func (this *ArticleParserFixture) Convert(content string) (string, error) {
	return content + " CONVERTED", this.markdownErr
}

func (this *ArticleParserFixture) Setup() {
	this.parser = NewArticleParser(this)
}

func (this *ArticleParserFixture) TestUnhandledTypeEmitted() {
	this.parser.Do("wrong-type", this.Output)
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *ArticleParserFixture) TestArticleMetaAndContentReadFromDiskAndEmitted() {
	this.parser.Do(contracts.SourceFile(article1Content), this.Output)
	this.So(this.outputs, should.Equal, []any{
		contracts.Article{
			Slug:  "/article/1",
			Title: "Article 1",
			Date:  time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC),
			Body:  "The contents of article 1. CONVERTED",
		},
	})
}
func (this *ArticleParserFixture) TestMissingDivider() {
	this.parser.Do(contracts.SourceFile("{} Content without separator"), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errMalformedSource)
	}
}
func (this *ArticleParserFixture) TestMalformedMetadata() {
	this.parser.Do(contracts.SourceFile("{bad-json}\n+++\nContent"), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errMalformedSource)
	}
}
func (this *ArticleParserFixture) TestInvalidMarkdown() {
	this.markdownErr = errors.New("boink")
	this.parser.Do(contracts.SourceFile(article1Content), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errMalformedSource)
	}
}

const article1Content = `{
	"date": "2024-09-04T00:00:00Z",
	"slug": "/article/1",
	"title": "Article 1"
}

+++

The contents of article 1.`
