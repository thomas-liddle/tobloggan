package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleParserFixture(t *testing.T) {
	gunit.Run(new(ArticleParserFixture), t)
}

var goodArticle = contracts.Article{
	Date:  time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC),
	Slug:  "/article/1",
	Title: "Article 1",
	Body:  "The contents of article 1.",
}

type ArticleParserFixture struct {
	StationFixture
}

func (this *ArticleParserFixture) Setup() {
	this.station = NewArticleParser()
}

func (this *ArticleParserFixture) TestArticleMetaAndContentReadFromDiskAndEmitted() {
	this.do(contracts.SourceFile(article1Content))
	this.assertOutputs(goodArticle)
}

func (this *ArticleParserFixture) TestMissingDivider() {
	this.do(contracts.SourceFile(article1ContentMissingDivider))
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errMalformedContent)
	}
}

func (this *ArticleParserFixture) TestMalformedMetadata() {
	this.do("la la la")
	this.assertOutputs("la la la")
}

const article1Content = `{
	"date": "2024-09-04T00:00:00Z",
	"slug": "/article/1",
	"title": "Article 1"
}

+++

The contents of article 1.`

const article1ContentMissingDivider = `{
	"date": "2024-09-04T00:00:00Z",
	"slug": "/article/1",
	"title": "Article 1"
}

The contents of article 1.`
