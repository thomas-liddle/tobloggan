package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleValidatorFixture(t *testing.T) {
	gunit.Run(new(ArticleValidatorFixture), t)
}

type ArticleValidatorFixture struct {
	StationFixture
}

func (this *ArticleValidatorFixture) Setup() {
	this.station = NewArticleValidator()
}

func (this *ArticleValidatorFixture) TestValidArticle() {
	this.do(goodArticle)
	this.assertOutputs(goodArticle)
}

func (this *ArticleValidatorFixture) TestInvalidSlugs() {
	this.do(invalidSlugArticle)
	this.So(this.outputs, should.HaveLength, 1)
	this.So(this.outputs[0], should.Wrap, ErrInvalidSlug)
}

func (this *ArticleValidatorFixture) TestInvalidTitles() {
	this.do(invalidTitleArticle)
	this.So(this.outputs, should.HaveLength, 1)
	this.So(this.outputs[0], should.Wrap, ErrInvalidTitle)
}

func (this *ArticleValidatorFixture) TestSlugsMustBeUnique() {
	this.do(goodArticle)
	this.do(goodArticle)
	this.So(this.outputs, should.HaveLength, 2)
	_, isArticle := this.outputs[0].(contracts.Article)
	this.So(isArticle, should.BeTrue)
	this.So(this.outputs[1], should.Wrap, ErrDuplicateSlug)
}

var invalidSlugArticle = contracts.Article{
	Date:  time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC),
	Slug:  "//*",
	Title: "Article 1",
	Body:  "The contents of article 1.",
}

var invalidTitleArticle = contracts.Article{
	Date:  time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC),
	Slug:  "/article/1",
	Title: "",
	Body:  "The contents of article 1.",
}
